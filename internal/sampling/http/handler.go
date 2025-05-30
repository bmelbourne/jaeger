// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	t2p "github.com/jaegertracing/jaeger/internal/converter/thrift/jaeger"
	"github.com/jaegertracing/jaeger/internal/metrics"
	p2json "github.com/jaegertracing/jaeger/internal/uimodel/converter/v1/json"
)

const mimeTypeApplicationJSON = "application/json"

var errBadRequest = errors.New("bad request")

type ClientConfigManager interface {
	GetSamplingStrategy(ctx context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error)
}

// HandlerParams contains parameters that must be passed to NewHTTPHandler.
type HandlerParams struct {
	ConfigManager  ClientConfigManager // required
	MetricsFactory metrics.Factory     // required

	// BasePath will be used as a prefix for the endpoints, e.g. "/api"
	BasePath string

	// LegacySamplingEndpoint enables returning sampling strategy from "/" endpoint
	// using Thrift 0.9.2 enum codes.
	LegacySamplingEndpoint bool
}

// Handler implements endpoints for used by Jaeger clients to retrieve client configuration,
// such as sampling strategies.
type Handler struct {
	params  HandlerParams
	metrics struct {
		// Number of good sampling requests
		SamplingRequestSuccess metrics.Counter `metric:"http-server.requests" tags:"type=sampling"`

		// Number of good sampling requests against the old endpoint / using Thrift 0.9.2 enum codes
		LegacySamplingRequestSuccess metrics.Counter `metric:"http-server.requests" tags:"type=sampling-legacy"`

		// Number of bad requests (400s)
		BadRequest metrics.Counter `metric:"http-server.errors" tags:"status=4xx,source=all"`

		// Number of collector proxy failures
		CollectorProxyFailures metrics.Counter `metric:"http-server.errors" tags:"status=5xx,source=collector-proxy"`

		// Number of bad responses due to malformed thrift
		BadThriftFailures metrics.Counter `metric:"http-server.errors" tags:"status=5xx,source=thrift"`

		// Number of bad responses due to proto conversion
		BadProtoFailures metrics.Counter `metric:"http-server.errors" tags:"status=5xx,source=proto"`

		// Number of failed response writes from http server
		WriteFailures metrics.Counter `metric:"http-server.errors" tags:"status=5xx,source=write"`
	}
}

// NewHandler creates new HTTPHandler.
func NewHandler(params HandlerParams) *Handler {
	handler := &Handler{params: params}
	metrics.MustInit(&handler.metrics, params.MetricsFactory, nil)
	return handler
}

// RegisterRoutes registers configuration handlers with Gorilla Router.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	prefix := h.params.BasePath
	if h.params.LegacySamplingEndpoint {
		router.HandleFunc(
			prefix+"/",
			func(w http.ResponseWriter, r *http.Request) {
				h.serveSamplingHTTP(w, r, h.encodeThriftLegacy)
			},
		).Methods(http.MethodGet)
	}
	router.HandleFunc(
		prefix+"/sampling",
		func(w http.ResponseWriter, r *http.Request) {
			h.serveSamplingHTTP(w, r, h.encodeProto)
		},
	).Methods(http.MethodGet)
}

// RegisterRoutes registers configuration handlers with HTTP Router.
func (h *Handler) RegisterRoutesWithHTTP(router *http.ServeMux) {
	prefix := h.params.BasePath
	router.HandleFunc(
		prefix+"/",
		func(w http.ResponseWriter, r *http.Request) {
			h.serveSamplingHTTP(w, r, h.encodeThriftLegacy)
		},
	)
}

func (h *Handler) serviceFromRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	services := r.URL.Query()["service"]
	if len(services) != 1 {
		h.metrics.BadRequest.Inc(1)
		http.Error(w, "'service' parameter must be provided once", http.StatusBadRequest)
		return "", errBadRequest
	}
	return services[0], nil
}

func (h *Handler) writeJSON(w http.ResponseWriter, jsonData []byte) error {
	w.Header().Add("Content-Type", mimeTypeApplicationJSON)
	if _, err := w.Write(jsonData); err != nil {
		h.metrics.WriteFailures.Inc(1)
		return err
	}
	return nil
}

func (h *Handler) serveSamplingHTTP(
	w http.ResponseWriter,
	r *http.Request,
	encoder func(strategy *api_v2.SamplingStrategyResponse) ([]byte, error),
) {
	service, err := h.serviceFromRequest(w, r)
	if err != nil {
		return
	}
	resp, err := h.params.ConfigManager.GetSamplingStrategy(r.Context(), service)
	if err != nil {
		h.metrics.CollectorProxyFailures.Inc(1)
		http.Error(w, fmt.Sprintf("collector error: %+v", err), http.StatusInternalServerError)
		return
	}
	jsonBytes, err := encoder(resp)
	if err != nil {
		http.Error(w, "cannot marshall to JSON", http.StatusInternalServerError)
		return
	}
	if err = h.writeJSON(w, jsonBytes); err != nil {
		return
	}
}

func (h *Handler) encodeThriftLegacy(strategy *api_v2.SamplingStrategyResponse) ([]byte, error) {
	tStrategy, err := t2p.ConvertSamplingResponseFromDomain(strategy)
	if err != nil {
		h.metrics.BadThriftFailures.Inc(1)
		return nil, fmt.Errorf("ConvertSamplingResponseFromDomain failed: %w", err)
	}
	jsonBytes, err := json.Marshal(tStrategy)
	if err != nil {
		h.metrics.BadThriftFailures.Inc(1)
		return nil, err
	}
	jsonBytes = h.encodeThriftEnums092(jsonBytes)
	h.metrics.LegacySamplingRequestSuccess.Inc(1)
	return jsonBytes, nil
}

func (h *Handler) encodeProto(strategy *api_v2.SamplingStrategyResponse) ([]byte, error) {
	str, err := p2json.SamplingStrategyResponseToJSON(strategy)
	if err != nil {
		h.metrics.BadProtoFailures.Inc(1)
		return nil, fmt.Errorf("SamplingStrategyResponseToJSON failed: %w", err)
	}
	h.metrics.SamplingRequestSuccess.Inc(1)
	return []byte(str), nil
}

var samplingStrategyTypes = []api_v2.SamplingStrategyType{
	api_v2.SamplingStrategyType_PROBABILISTIC,
	api_v2.SamplingStrategyType_RATE_LIMITING,
}

// Replace string enum values produced from Thrift 0.9.3 generated classes
// with integer codes produced from Thrift 0.9.2 generated classes.
//
// For example:
//
// Thrift 0.9.2 classes generate this JSON:
// {"strategyType":0,"probabilisticSampling":{"samplingRate":0.5},"rateLimitingSampling":null,"operationSampling":null}
//
// Thrift 0.9.3 classes generate this JSON:
// {"strategyType":"PROBABILISTIC","probabilisticSampling":{"samplingRate":0.5}}
func (*Handler) encodeThriftEnums092(jsonData []byte) []byte {
	str := string(jsonData)
	for _, strategyType := range samplingStrategyTypes {
		str = strings.Replace(
			str,
			fmt.Sprintf(`"strategyType":"%s"`, strategyType.String()),
			fmt.Sprintf(`"strategyType":%d`, strategyType),
			1,
		)
	}
	return []byte(str)
}
