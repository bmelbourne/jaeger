### 🇷🇺 A message to people of Russia

If you currently live in Russia, please read [this message](./_To_People_of_Russia.md).

Changes by Version
==================

<details>
<summary>next release template</summary>

next release v1.x.x / v2.x.x (yyyy-mm-dd)
-------------------------------

### Backend Changes

run `make changelog` to generate content

### 📊 UI Changes

copy from UI changelog

</details>

v1.71.0 / v2.8.0 (2025-07-03)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* [es] materialize span.kind and span.status tags ([@pipiland2612](https://github.com/pipiland2612) in [#7272](https://github.com/jaegertracing/jaeger/pull/7272))
* Make jaeger.es.disablelegacyid feature stable ([@yurishkuro](https://github.com/yurishkuro) in [#7267](https://github.com/jaegertracing/jaeger/pull/7267))

#### ✨ New Features

* [v2] switch memory backend to storage api v2 implementation ([@Manik2708](https://github.com/Manik2708) in [#7157](https://github.com/jaegertracing/jaeger/pull/7157))

#### 🐞 Bug fixes, Minor Improvements

* Fix panic when reading malformed warning attribute ([@yurishkuro](https://github.com/yurishkuro) in [#7293](https://github.com/jaegertracing/jaeger/pull/7293))
* [fix] prevent panic when sanitizing read-only traces with multiple exporters ([@victornguen](https://github.com/victornguen) in [#7245](https://github.com/jaegertracing/jaeger/pull/7245))
* Update elasticsearch to use olivere/elastic/v7 ([@pipiland2612](https://github.com/pipiland2612) in [#7244](https://github.com/jaegertracing/jaeger/pull/7244))
* Repoint docker compose files to use cr.jaegertracing.io ([@jkowall](https://github.com/jkowall) in [#7240](https://github.com/jaegertracing/jaeger/pull/7240))
* [es/v2] add metrics decorator for trace reader ([@Manik2708](https://github.com/Manik2708) in [#7201](https://github.com/jaegertracing/jaeger/pull/7201))

#### 🚧 Experimental Features

* [spm] getcallrate implementation ([@pipiland2612](https://github.com/pipiland2612) in [#7229](https://github.com/jaegertracing/jaeger/pull/7229))
* [cassandra] give responsibility of creating v2 factory to storage backend ([@Manik2708](https://github.com/Manik2708) in [#7228](https://github.com/jaegertracing/jaeger/pull/7228))
* Jaeger demo on kubernetes ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#7262](https://github.com/jaegertracing/jaeger/pull/7262))
* [clickhouse] implement gettraces for clickhouse storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7207](https://github.com/jaegertracing/jaeger/pull/7207))
* [spm] and esclient to es/reader.go and refactor test ([@pipiland2612](https://github.com/pipiland2612) in [#7216](https://github.com/jaegertracing/jaeger/pull/7216))
* Add skeleton implementation for es as metricstorage for spm ([@pipiland2612](https://github.com/pipiland2612) in [#7209](https://github.com/jaegertracing/jaeger/pull/7209))

### 📊 UI Changes

#### ⚙️ Refactoring

* Convert `opsgraph.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2914](https://github.com/jaegertracing/jaeger-ui/pull/2914))
* Convert `regiondemo.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2910](https://github.com/jaegertracing/jaeger-ui/pull/2910))
* Convert `dividerdemo.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2909](https://github.com/jaegertracing/jaeger-ui/pull/2909))
* Convert `draggablemanagerdemo.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2908](https://github.com/jaegertracing/jaeger-ui/pull/2908))
* Convert `nameselector.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2889](https://github.com/jaegertracing/jaeger-ui/pull/2889))
* Convert `copyicon.tsx` from class component to functional component ([@Parship999](https://github.com/Parship999) in [#2887](https://github.com/jaegertracing/jaeger-ui/pull/2887))

</details>

v1.70.0 / v2.7.0 (2025-06-10)
-------------------------------

### Backend Changes

#### ✨ New Features

* [feat] use v2 es/os storage in jaeger-v2 ([@Manik2708](https://github.com/Manik2708) in [#7151](https://github.com/jaegertracing/jaeger/pull/7151))

#### 🐞 Bug fixes, Minor Improvements

* Feat: add option to disable elasticsearch health check ([@timonegk](https://github.com/timonegk) in [#7212](https://github.com/jaegertracing/jaeger/pull/7212))
* Fix(elasticsearch): respect explicitly configured replicas=0 in index… ([@masihkhatibzadeh99](https://github.com/masihkhatibzadeh99) in [#7160](https://github.com/jaegertracing/jaeger/pull/7160))
* Add sanitizers for negative span duration ([@iypetrov](https://github.com/iypetrov) in [#7122](https://github.com/jaegertracing/jaeger/pull/7122))
* [fix] fix prometheus label value is not valid utf 8 cause api timeout ([@iypetrov](https://github.com/iypetrov) in [#7128](https://github.com/jaegertracing/jaeger/pull/7128))
* Add retries to ilm client ([@iypetrov](https://github.com/iypetrov) in [#7120](https://github.com/jaegertracing/jaeger/pull/7120))
* Add retry configuration to storage exporter ([@kumarlokesh](https://github.com/kumarlokesh) in [#7132](https://github.com/jaegertracing/jaeger/pull/7132))
* [fix] restore es metrics ([@AnmolxSingh](https://github.com/AnmolxSingh) in [#7006](https://github.com/jaegertracing/jaeger/pull/7006))
* [fix] propagate environment variables to binary from integration tests ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7112](https://github.com/jaegertracing/jaeger/pull/7112))

#### 🚧 Experimental Features

* [refactor] rework clickhouse schema structure ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7181](https://github.com/jaegertracing/jaeger/pull/7181))
* [clickhouse] implement getoperations for trace reader in clickhouse storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7180](https://github.com/jaegertracing/jaeger/pull/7180))
* [clickhouse] implement getservices for trace reader in clickhouse storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7159](https://github.com/jaegertracing/jaeger/pull/7159))
* [v2] implement `getdependencies` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7154](https://github.com/jaegertracing/jaeger/pull/7154))
* [v2] implement `gettraces` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7152](https://github.com/jaegertracing/jaeger/pull/7152))
* [v2] implement `findtraceids` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7143](https://github.com/jaegertracing/jaeger/pull/7143))
* Add description for docker-compose-elasticsearch.yml ([@pipiland2612](https://github.com/pipiland2612) in [#7146](https://github.com/jaegertracing/jaeger/pull/7146))
* Add e2e test for docker-compose-elasticsearch.yml file ([@pipiland2612](https://github.com/pipiland2612) in [#7145](https://github.com/jaegertracing/jaeger/pull/7145))
* Add docker-compose-elasticsearch.yml and its sample configuration ([@pipiland2612](https://github.com/pipiland2612) in [#7144](https://github.com/jaegertracing/jaeger/pull/7144))
* [v2] implement `findtraces` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7062](https://github.com/jaegertracing/jaeger/pull/7062))

#### ⚙️ Refactoring

* Relocate the docker directory from the root directory to under scripts/build ([@ris-tlp](https://github.com/ris-tlp) in [#7189](https://github.com/jaegertracing/jaeger/pull/7189))
* [refactor] move sanitizer ([@yurishkuro](https://github.com/yurishkuro) in [#7158](https://github.com/jaegertracing/jaeger/pull/7158))
* [es][v2] refactor the factory of v1 to make it reusable for v2 ([@Manik2708](https://github.com/Manik2708) in [#7086](https://github.com/jaegertracing/jaeger/pull/7086))
* [refactor] allow storage cleaner to be overridden via environment variable ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7114](https://github.com/jaegertracing/jaeger/pull/7114))
* [refactor] remove archive storage from grpc config ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7113](https://github.com/jaegertracing/jaeger/pull/7113))
* [grpc] allow remote storage endpoint to be set via environment variable ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7111](https://github.com/jaegertracing/jaeger/pull/7111))
* [refactor] use proto files from `jaeger-idl` for remote storage api ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7104](https://github.com/jaegertracing/jaeger/pull/7104))


### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Fix react fragment key issues in multiple components ([@Parship999](https://github.com/Parship999) in [#2823](https://github.com/jaegertracing/jaeger-ui/pull/2823))
* Move tracediff header chevron icon ([@Parship999](https://github.com/Parship999) in [#2845](https://github.com/jaegertracing/jaeger-ui/pull/2845))
* Feat: filter logs based on the selected time range ([@tejas-raskar](https://github.com/tejas-raskar) in [#2844](https://github.com/jaegertracing/jaeger-ui/pull/2844))
* Enhance tracediff ui components ([@Parship999](https://github.com/Parship999) in [#2806](https://github.com/jaegertracing/jaeger-ui/pull/2806))
* Rewrite computeselftime to improve performance on trace statistics page ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2767](https://github.com/jaegertracing/jaeger-ui/pull/2767))
* Fix array return pattern in `labeledlist` component ([@Parship999](https://github.com/Parship999) in [#2812](https://github.com/jaegertracing/jaeger-ui/pull/2812))
* Allow json logs to occupy entire available width ([@tejas-raskar](https://github.com/tejas-raskar) in [#2814](https://github.com/jaegertracing/jaeger-ui/pull/2814))
* Feat: convert monitoratmemptystate to a functional component ([@vishvamsinh28](https://github.com/vishvamsinh28) in [#2790](https://github.com/jaegertracing/jaeger-ui/pull/2790))
* Replace deprecated `overlayclassname` with `classnames.root` ([@abhayporwals](https://github.com/abhayporwals) in [#2772](https://github.com/jaegertracing/jaeger-ui/pull/2772))
* [fix]: reduce default minimum allowed zoom ([@hari45678](https://github.com/hari45678) in [#2775](https://github.com/jaegertracing/jaeger-ui/pull/2775))
* Fix dependencygraph dag extra render ([@mdwyer6](https://github.com/mdwyer6) in [#2749](https://github.com/jaegertracing/jaeger-ui/pull/2749))

#### ⚙️ Refactoring

* Convert qualitymetrics components to functional components ([@Parship999](https://github.com/Parship999) in [#2856](https://github.com/jaegertracing/jaeger-ui/pull/2856))
* Refactor: spandetailrow to functional component ([@tejas-raskar](https://github.com/tejas-raskar) in [#2827](https://github.com/jaegertracing/jaeger-ui/pull/2827))
* Migrate tracetimelineviewerimpl to a functional component ([@tejas-raskar](https://github.com/tejas-raskar) in [#2816](https://github.com/jaegertracing/jaeger-ui/pull/2816))
* Refactor canvasspangraph to functional component and improve test coverage ([@vishvamsinh28](https://github.com/vishvamsinh28) in [#2824](https://github.com/jaegertracing/jaeger-ui/pull/2824))

v1.69.0 / v2.6.0 (2025-05-08)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* Feat(elasticsearch): Add flag to enable gzip compression by default ([@timonegk](https://github.com/timonegk) in [#7072](https://github.com/jaegertracing/jaeger/pull/7072))
* Only Remote Storage API v2 is supported in Jaeger v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6969](https://github.com/jaegertracing/jaeger/pull/6969))


#### ✨ New Features

* Add filterprocessor ([@yurishkuro](https://github.com/yurishkuro) in [#7094](https://github.com/jaegertracing/jaeger/pull/7094))

#### 🐞 Bug fixes, Minor Improvements

* Upgrade reverse-proxy example to jaeger-v2 ([@yurishkuro](https://github.com/yurishkuro) in [#7076](https://github.com/jaegertracing/jaeger/pull/7076))
* Add pprof extension ([@denysvitali](https://github.com/denysvitali) in [#7073](https://github.com/jaegertracing/jaeger/pull/7073))
* [es][v1] change the db tag value from `string` to `any` type ([@Manik2708](https://github.com/Manik2708) in [#6998](https://github.com/jaegertracing/jaeger/pull/6998))
* Remove outdated info related to jaeger exporter ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#6987](https://github.com/jaegertracing/jaeger/pull/6987))
* [bug] fix the version module path in ldflags ([@developer-guy](https://github.com/developer-guy) in [#6990](https://github.com/jaegertracing/jaeger/pull/6990))

#### 🚧 Experimental Features

* [es][v2] implement `getdependenies` and `writedependencies` ([@Manik2708](https://github.com/Manik2708) in [#7085](https://github.com/jaegertracing/jaeger/pull/7085))
* [clickhouse] convert otel traces model to  native format ([@zhengkezhou1](https://github.com/zhengkezhou1) in [#6935](https://github.com/jaegertracing/jaeger/pull/6935))
* [es] make `nestedtags` and `elevatedtags` distinction at `corespanreader` level ([@Manik2708](https://github.com/Manik2708) in [#7067](https://github.com/jaegertracing/jaeger/pull/7067))
* [es][v2] move `coredependencystore` and `dbmodel` from v1 to v2 ([@Manik2708](https://github.com/Manik2708) in [#7079](https://github.com/jaegertracing/jaeger/pull/7079))
* [grpc][v2] use standard otlp receiver for grpc storage write path ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7065](https://github.com/jaegertracing/jaeger/pull/7065))
* [es][v2] implement `findtraces` for es/os for v2 ([@Manik2708](https://github.com/Manik2708) in [#7021](https://github.com/jaegertracing/jaeger/pull/7021))
* [refactor] remove `jaeger_query` extension from remote storage backend config ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7059](https://github.com/jaegertracing/jaeger/pull/7059))
* [v2][remote-storage] implement remote storage extension ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7043](https://github.com/jaegertracing/jaeger/pull/7043))
* [es][v2] implement `gettraces` for es/os ([@Manik2708](https://github.com/Manik2708) in [#7054](https://github.com/jaegertracing/jaeger/pull/7054))
* [v2] implement `getoperations` and `getservices` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7053](https://github.com/jaegertracing/jaeger/pull/7053))
* [v2] implement `findtraceids` for es/os ([@Manik2708](https://github.com/Manik2708) in [#7035](https://github.com/jaegertracing/jaeger/pull/7035))
* [v2] implement `writetraces` for memory backend ([@Manik2708](https://github.com/Manik2708) in [#7027](https://github.com/jaegertracing/jaeger/pull/7027))
* [es] refactor `dependencystore` to make it reusable for v2 apis ([@Manik2708](https://github.com/Manik2708) in [#7044](https://github.com/jaegertracing/jaeger/pull/7044))
* [grpc][v2] use v2 grpc factory in storage extension ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6969](https://github.com/jaegertracing/jaeger/pull/6969))
* [grpc][v2] register grpc v2 handler in remote-storage server ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7037](https://github.com/jaegertracing/jaeger/pull/7037))
* [es][v2] implement `getoperations` and `getservices` for v2 ([@Manik2708](https://github.com/Manik2708) in [#7025](https://github.com/jaegertracing/jaeger/pull/7025))
* [es][v2] implement `writetraces` for v2 ([@Manik2708](https://github.com/Manik2708) in [#7020](https://github.com/jaegertracing/jaeger/pull/7020))
* [grpc][v2] implement `getdependencies` in grpc v2 server ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7016](https://github.com/jaegertracing/jaeger/pull/7016))
* [grpc][v2] implement otlp exporter api in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7012](https://github.com/jaegertracing/jaeger/pull/7012))
* [es][v2] change the db tag value from `string` to `any` type ([@Manik2708](https://github.com/Manik2708) in [#6994](https://github.com/jaegertracing/jaeger/pull/6994))
* [grpc][v2] implement findtraceids in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7003](https://github.com/jaegertracing/jaeger/pull/7003))
* [grpc][v2] implement `findtraces` in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6992](https://github.com/jaegertracing/jaeger/pull/6992))
* [grpc][v2] implement `gettraces` in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6985](https://github.com/jaegertracing/jaeger/pull/6985))
* [grpc][v2] implement getservices in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6984](https://github.com/jaegertracing/jaeger/pull/6984))
* [grpc][v2] implement `getservices` in grpc v2 handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6980](https://github.com/jaegertracing/jaeger/pull/6980))

#### 👷 CI Improvements

* Do not run binary size check on push to main ([@yurishkuro](https://github.com/yurishkuro) in [#7096](https://github.com/jaegertracing/jaeger/pull/7096))
* Check that version number is corectly embedded in the binary ([@yurishkuro](https://github.com/yurishkuro) in [#7092](https://github.com/jaegertracing/jaeger/pull/7092))
* Update module github.com/vektra/mockery/v2 to v3 ([@AnmolxSingh](https://github.com/AnmolxSingh) in [#7051](https://github.com/jaegertracing/jaeger/pull/7051))
* [fix] add query integration test to workflows file ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7056](https://github.com/jaegertracing/jaeger/pull/7056))
* Enable mockery/with-expecter ([@yurishkuro](https://github.com/yurishkuro) in [#7046](https://github.com/jaegertracing/jaeger/pull/7046))
* Fix paths in mockery config ([@yurishkuro](https://github.com/yurishkuro) in [#7045](https://github.com/jaegertracing/jaeger/pull/7045))
* Fix flakiness in runindexcleanertest by filtering jaeger indices ([@0xShubhamSolanki](https://github.com/0xShubhamSolanki) in [#7004](https://github.com/jaegertracing/jaeger/pull/7004))
* Add e2e integration test for query service ([@pipiland2612](https://github.com/pipiland2612) in [#6966](https://github.com/jaegertracing/jaeger/pull/6966))
* #5608 improve spm e2e test with test for error rate ([@pipiland2612](https://github.com/pipiland2612) in [#6991](https://github.com/jaegertracing/jaeger/pull/6991))

#### ⚙️ Refactoring

* [es] make `nestedtags` and `fieldtags` distinction at `corespanwriter` level ([@Manik2708](https://github.com/Manik2708) in [#6946](https://github.com/jaegertracing/jaeger/pull/6946))
* [refactor] change remote storage server to accept v2 factories ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7024](https://github.com/jaegertracing/jaeger/pull/7024))
* [refactor] consolidate v1/v2 writer factory adapter functionality ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7022](https://github.com/jaegertracing/jaeger/pull/7022))
* [refactor] consolidate v1/v2 reader factory adapter functionality ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7019](https://github.com/jaegertracing/jaeger/pull/7019))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Reduce load time of trace page by deferring critical path tooltip ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2718](https://github.com/jaegertracing/jaeger-ui/pull/2718))
* Migrate copyicon tests ([@nojaf](https://github.com/nojaf) in [#2727](https://github.com/jaegertracing/jaeger-ui/pull/2727))
* [fix]: make reset icon in sdg more intuitive ([@hari45678](https://github.com/hari45678) in [#2723](https://github.com/jaegertracing/jaeger-ui/pull/2723))
* Migrate from enzyme to @testing-library/react in keyboardshortshelp ([@nojaf](https://github.com/nojaf) in [#2725](https://github.com/jaegertracing/jaeger-ui/pull/2725))
* Improve performance of trace statistics page when grouping by tag ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2724](https://github.com/jaegertracing/jaeger-ui/pull/2724))
* Improve performance of expanding and collapsing spans ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2722](https://github.com/jaegertracing/jaeger-ui/pull/2722))
* Improve performance of trace statistics ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2721](https://github.com/jaegertracing/jaeger-ui/pull/2721))
* [feat]: add context menu on node to dag ([@hari45678](https://github.com/hari45678) in [#2719](https://github.com/jaegertracing/jaeger-ui/pull/2719))
* Fix grouping on trace statistics page for tags ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2717](https://github.com/jaegertracing/jaeger-ui/pull/2717))
* Improve performance when expanding/collapsing span details ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2716](https://github.com/jaegertracing/jaeger-ui/pull/2716))

#### 👷 CI Improvements

* Add ability to use typescript in tests ([@DamianMaslanka5](https://github.com/DamianMaslanka5) in [#2731](https://github.com/jaegertracing/jaeger-ui/pull/2731))

#### ⚙️ Refactoring

* [es] make `nestedtags` and `fieldtags` distinction at `corespanwriter` level ([@Manik2708](https://github.com/Manik2708) in [#6946](https://github.com/jaegertracing/jaeger/pull/6946))
* [refactor] change remote storage server to accept v2 factories ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7024](https://github.com/jaegertracing/jaeger/pull/7024))
* [refactor] consolidate v1/v2 writer factory adapter functionality ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7022](https://github.com/jaegertracing/jaeger/pull/7022))
* [refactor] consolidate v1/v2 reader factory adapter functionality ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#7019](https://github.com/jaegertracing/jaeger/pull/7019))

v1.68.0 / v2.5.0 (2025-04-05)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* Remove sampling.strategies.bugfix-5270 flag and mark feature stable ([@yurishkuro](https://github.com/yurishkuro) in [#6872](https://github.com/jaegertracing/jaeger/pull/6872))

#### 🐞 Bug fixes, Minor Improvements

* Minor fixes to release checklist generator ([@albertteoh](https://github.com/albertteoh) in [#6976](https://github.com/jaegertracing/jaeger/pull/6976))
* Support configuring prometheus.extra_query_parameters via cli ([@andreasgerstmayr](https://github.com/andreasgerstmayr) in [#6931](https://github.com/jaegertracing/jaeger/pull/6931))
* Cleanup legacy models ([@yurishkuro](https://github.com/yurishkuro) in [#6875](https://github.com/jaegertracing/jaeger/pull/6875))
* 🪦 remove agent code ([@yurishkuro](https://github.com/yurishkuro) in [#6868](https://github.com/jaegertracing/jaeger/pull/6868))
* [es] refactor `findtraces` and `gettrace` of spanreader to make them reusable for v2 apis ([@Manik2708](https://github.com/Manik2708) in [#6845](https://github.com/jaegertracing/jaeger/pull/6845))
* [es] add feature to stop legacy trace ids handling in spanreader ([@Manik2708](https://github.com/Manik2708) in [#6848](https://github.com/jaegertracing/jaeger/pull/6848))
* Feat: move pkg/testutils to internal/testutils ([@jinjiaKarl](https://github.com/jinjiaKarl) in [#6840](https://github.com/jaegertracing/jaeger/pull/6840))
* [fix] allow es-index-cleaner to delete indices based on current time ([@Asatyam](https://github.com/Asatyam) in [#6790](https://github.com/jaegertracing/jaeger/pull/6790))
* [chore] remove gogoproto annotations from `trace_storage.proto` and `dependency_storage.proto` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6819](https://github.com/jaegertracing/jaeger/pull/6819))

#### 🚧 Experimental Features

* [es][v2] add snapshot tests for spans conversion ([@Manik2708](https://github.com/Manik2708) in [#6970](https://github.com/jaegertracing/jaeger/pull/6970))
* [grpc][v2] implement grpc v2 factory ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6968](https://github.com/jaegertracing/jaeger/pull/6968))
* [grpc][v2] implement `findtraces` call in grpc reader for remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6962](https://github.com/jaegertracing/jaeger/pull/6962))
* [grpc][v2] implement `gettraces` call in grpc reader for remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6857](https://github.com/jaegertracing/jaeger/pull/6857))
* [es][v2] refactor `from_dbmodel` and `to_dbmodel` to accept and return db spans ([@Manik2708](https://github.com/Manik2708) in [#6934](https://github.com/jaegertracing/jaeger/pull/6934))
* [grpc][v2] implement v2 grpc dependency reader ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6933](https://github.com/jaegertracing/jaeger/pull/6933))
* [es][v2] copy jaeger<->otlp translator from otel contrib ([@Manik2708](https://github.com/Manik2708) in [#6923](https://github.com/jaegertracing/jaeger/pull/6923))
* [grpc][v2] implement v2 grpc trace writer ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6919](https://github.com/jaegertracing/jaeger/pull/6919))
* [grpc][v2] implement `findtraceids` call in grpc reader for remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6858](https://github.com/jaegertracing/jaeger/pull/6858))
* [grpc][v2] implement `getoperations` call in grpc reader for remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6843](https://github.com/jaegertracing/jaeger/pull/6843))
* [grpc][v2] implement `getservices` call in grpc reader for remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6829](https://github.com/jaegertracing/jaeger/pull/6829))
* [refactor] return chunk of traces from remote storage api v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6830](https://github.com/jaegertracing/jaeger/pull/6830))

#### 👷 CI Improvements

* [all-in-one] avoid multi-arch builds in merge queue (#6880) ([@sAchin-680](https://github.com/sAchin-680) in [#6882](https://github.com/jaegertracing/jaeger/pull/6882))
* Instruct renovate to pin github action hashes ([@yurishkuro](https://github.com/yurishkuro) in [#6860](https://github.com/jaegertracing/jaeger/pull/6860))

#### ⚙️ Refactoring

* Move model/json/model.go to internal/uimodel/converter/v1 ([@pipiland2612](https://github.com/pipiland2612) in [#6973](https://github.com/jaegertracing/jaeger/pull/6973))
* Add usetesting linter and fix lint issues (#6892) ([@anurag-rajawat](https://github.com/anurag-rajawat) in [#6972](https://github.com/jaegertracing/jaeger/pull/6972))
* Delete empty pkg package ([@pipiland2612](https://github.com/pipiland2612) in [#6967](https://github.com/jaegertracing/jaeger/pull/6967))
* Move pkg/otelsemconv to internal/telemetry/otelsemconv ([@pipiland2612](https://github.com/pipiland2612) in [#6961](https://github.com/jaegertracing/jaeger/pull/6961))
* Move pkg/cassandra to internal/storage/cassandra ([@pipiland2612](https://github.com/pipiland2612) in [#6960](https://github.com/jaegertracing/jaeger/pull/6960))
* Move pkg/adjuster to cmd/query/app/querysvc/internal/adjuster ([@pipiland2612](https://github.com/pipiland2612) in [#6956](https://github.com/jaegertracing/jaeger/pull/6956))
* Remove package pkg/netutils ([@pipiland2612](https://github.com/pipiland2612) in [#6955](https://github.com/jaegertracing/jaeger/pull/6955))
* [refactor] remove `traceschunk` type and stream otlp traces directly ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6954](https://github.com/jaegertracing/jaeger/pull/6954))
* [es] remove pointer signatures from `fromdbmodel` and `todbmodel` ([@Manik2708](https://github.com/Manik2708) in [#6942](https://github.com/jaegertracing/jaeger/pull/6942))
* Move proto-gen to internal ([@yurishkuro](https://github.com/yurishkuro) in [#6941](https://github.com/jaegertracing/jaeger/pull/6941))
* Move pkg/es to internal/storage/elasticsearch ([@danish9039](https://github.com/danish9039) in [#6937](https://github.com/jaegertracing/jaeger/pull/6937))
* Move pkg/distributedlock to internal/distributedlock ([@danish9039](https://github.com/danish9039) in [#6903](https://github.com/jaegertracing/jaeger/pull/6903))
* Move pkg/httpmetrics to internal/httpmetrics ([@danish9039](https://github.com/danish9039) in [#6905](https://github.com/jaegertracing/jaeger/pull/6905))
* Move pkg/{gogocodec,httpfs,bearertoken,boundqueue} to internal ([@sAchin-680](https://github.com/sAchin-680) in [#6896](https://github.com/jaegertracing/jaeger/pull/6896))
* Move pkg/metrics to internal/metrics ([@danish9039](https://github.com/danish9039) in [#6901](https://github.com/jaegertracing/jaeger/pull/6901))
* Move pkg/kafka to internal/kafka ([@danish9039](https://github.com/danish9039) in [#6908](https://github.com/jaegertracing/jaeger/pull/6908))
* Move pkg/prometheus to internal/config/promcfg ([@danish9039](https://github.com/danish9039) in [#6911](https://github.com/jaegertracing/jaeger/pull/6911))
* Move model/proto to internal/proto ([@danish9039](https://github.com/danish9039) in [#6918](https://github.com/jaegertracing/jaeger/pull/6918))
* [es] move db model out of `v1/elasticsearch/internal/spanstore/internal` ([@Manik2708](https://github.com/Manik2708) in [#6894](https://github.com/jaegertracing/jaeger/pull/6894))
* Move pkg/version to internal/version ([@danish9039](https://github.com/danish9039) in [#6913](https://github.com/jaegertracing/jaeger/pull/6913))
* Move model/converter to internal/converter ([@danish9039](https://github.com/danish9039) in [#6917](https://github.com/jaegertracing/jaeger/pull/6917))
* Move pkg/gzipfs to internal/gzipfs ([@sAchin-680](https://github.com/sAchin-680) in [#6897](https://github.com/jaegertracing/jaeger/pull/6897))
* Move pkg/jtracer to internal/jtracer ([@danish9039](https://github.com/danish9039) in [#6907](https://github.com/jaegertracing/jaeger/pull/6907))
* Move pkg/telemetry to internal/telemetry ([@danish9039](https://github.com/danish9039) in [#6912](https://github.com/jaegertracing/jaeger/pull/6912))
* Move pkg/fswatcher to internal/fswatcher ([@sAchin-680](https://github.com/sAchin-680) in [#6895](https://github.com/jaegertracing/jaeger/pull/6895))
* [es] separate the `corespanwriter` from `spanwriter` ([@Manik2708](https://github.com/Manik2708) in [#6883](https://github.com/jaegertracing/jaeger/pull/6883))
* Move pkg/config to internal/config ([@gentcod](https://github.com/gentcod) in [#6884](https://github.com/jaegertracing/jaeger/pull/6884))
* Move pkg/healthcheck to internal/healthcheck ([@danish9039](https://github.com/danish9039) in [#6888](https://github.com/jaegertracing/jaeger/pull/6888))
* Moved pkg/hostname to internal/hostname ([@danish9039](https://github.com/danish9039) in [#6886](https://github.com/jaegertracing/jaeger/pull/6886))
* Move pkg/recoveryhandler to internal/recoveryhandler ([@danish9039](https://github.com/danish9039) in [#6887](https://github.com/jaegertracing/jaeger/pull/6887))
* Move pkg/tenancy to internal/tenancy ([@danish9039](https://github.com/danish9039) in [#6889](https://github.com/jaegertracing/jaeger/pull/6889))
* Move pkg/normalizer to collector ([@danish9039](https://github.com/danish9039) in [#6877](https://github.com/jaegertracing/jaeger/pull/6877))
* Replace the use of model/converter/thrift/zipkin ([@shuraih775](https://github.com/shuraih775) in [#6879](https://github.com/jaegertracing/jaeger/pull/6879))
* [es] remove pointer signatures from `corespanreader` ([@Manik2708](https://github.com/Manik2708) in [#6874](https://github.com/jaegertracing/jaeger/pull/6874))
* [refactor] move interface to remove cmd/agent dependency ([@yurishkuro](https://github.com/yurishkuro) in [#6863](https://github.com/jaegertracing/jaeger/pull/6863))
* [agent] refactor udp server ([@yurishkuro](https://github.com/yurishkuro) in [#6852](https://github.com/jaegertracing/jaeger/pull/6852))
* [agent] remove unnecessary server interface ([@yurishkuro](https://github.com/yurishkuro) in [#6851](https://github.com/jaegertracing/jaeger/pull/6851))
* [es] refactor the `findtraceids` of spanreader to make them reusable for v2 apis ([@Manik2708](https://github.com/Manik2708) in [#6831](https://github.com/jaegertracing/jaeger/pull/6831))
* [es] refactor the `getoperations` and `getservices` of spanreader to make them reusable for v2 apis ([@Manik2708](https://github.com/Manik2708) in [#6828](https://github.com/jaegertracing/jaeger/pull/6828))


### 📊 UI Changes

#### ✨ New Features

* Feat: change dag styling and add search functionality ([@hari45678](https://github.com/hari45678) in [#2710](https://github.com/jaegertracing/jaeger-ui/pull/2710))

#### 🐞 Bug fixes, Minor Improvements

* Add sample graph data when in dev mode ([@hari45678](https://github.com/hari45678) in [#2698](https://github.com/jaegertracing/jaeger-ui/pull/2698))
* Add depth and layout controls and sfdp layout to dag view ([@hari45678](https://github.com/hari45678) in [#2691](https://github.com/jaegertracing/jaeger-ui/pull/2691))
* Add sfdp engine in @jaegertracing/plexus ([@hari45678](https://github.com/hari45678) in [#2690](https://github.com/jaegertracing/jaeger-ui/pull/2690))
* Add handling or error for invalid json formats and tests ([@rohitlohar45](https://github.com/rohitlohar45) in [#2689](https://github.com/jaegertracing/jaeger-ui/pull/2689))


v1.67.0 / v2.4.0 (2025-03-07)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* [query] drop support for shared grpc/http query server ports ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6695](https://github.com/jaegertracing/jaeger/pull/6695))

#### 🐞 Bug fixes, Minor Improvements

* [es] refactor the es spanwriter to make it reusable for v2 apis ([@Manik2708](https://github.com/Manik2708) in [#6796](https://github.com/jaegertracing/jaeger/pull/6796))
* [refactor] move internal `tracesdata` type to package `jptrace` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6809](https://github.com/jaegertracing/jaeger/pull/6809))
* Use empty slices instead of nil ([@zhengkezhou1](https://github.com/zhengkezhou1) in [#6799](https://github.com/jaegertracing/jaeger/pull/6799))
* [refactor] refactor `jptrace/attributes_tests.go` for readability ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6786](https://github.com/jaegertracing/jaeger/pull/6786))
* [refactor] converge v2 api with v2 remote storage api ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6784](https://github.com/jaegertracing/jaeger/pull/6784))
* Feat: enable configuration of hostnames for hotrod services ([@w-h-a](https://github.com/w-h-a) in [#6782](https://github.com/jaegertracing/jaeger/pull/6782))
* [refactor] change `tracequeryparams` to accept typed attributes ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6780](https://github.com/jaegertracing/jaeger/pull/6780))
* [refactor] decouple `tracequeryparams` from `query` in integration tests ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6779](https://github.com/jaegertracing/jaeger/pull/6779))
* [refactor] inline proto definiton of `keyvalue` from otel ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6775](https://github.com/jaegertracing/jaeger/pull/6775))
* [refactor] return start and end timestamps from findtraceids in v2 remote storage api ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6772](https://github.com/jaegertracing/jaeger/pull/6772))
* [refactor] return start and end timestamps from `findtraceids` in v2 api ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6770](https://github.com/jaegertracing/jaeger/pull/6770))
* Revert "add 'features' command to print available feature gates" ([@yurishkuro](https://github.com/yurishkuro) in [#6771](https://github.com/jaegertracing/jaeger/pull/6771))
* [remote-storage][v2] add complete idl for trace storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6737](https://github.com/jaegertracing/jaeger/pull/6737))
* [remote-storage][v2] add idl for dependency storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6738](https://github.com/jaegertracing/jaeger/pull/6738))
* [remote-storage][v2] add proto definition for `getservices` and `getoperations` rpc ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6736](https://github.com/jaegertracing/jaeger/pull/6736))
* Fix /qualitymetrics to return data in expected format ([@yurishkuro](https://github.com/yurishkuro) in [#6733](https://github.com/jaegertracing/jaeger/pull/6733))
* [remote-storage][v2] add proto definition for `gettraces` rpc ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6730](https://github.com/jaegertracing/jaeger/pull/6730))
* [bug][storage] make es-rollover idempotent by checking if the index or alias already exist ([@Manik2708](https://github.com/Manik2708) in [#6638](https://github.com/jaegertracing/jaeger/pull/6638))
* [refactor] use plain loops with iterators ([@yurishkuro](https://github.com/yurishkuro) in [#6722](https://github.com/jaegertracing/jaeger/pull/6722))
* Use stdlib iterators ([@yurishkuro](https://github.com/yurishkuro) in [#6714](https://github.com/jaegertracing/jaeger/pull/6714))
* Create a /quality-metrics endpoint ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#6608](https://github.com/jaegertracing/jaeger/pull/6608))
* Move pkg/cache to internal ([@won-js](https://github.com/won-js) in [#6720](https://github.com/jaegertracing/jaeger/pull/6720))
* [storage] change storage extension to hold v2 factories ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6699](https://github.com/jaegertracing/jaeger/pull/6699))
* Fix go alpine version to 1.24.0 ([@yurishkuro](https://github.com/yurishkuro) in [#6713](https://github.com/jaegertracing/jaeger/pull/6713))
* [refactor] conditionally implement interfaces in v1adapter factory ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6710](https://github.com/jaegertracing/jaeger/pull/6710))
* [fix] revert changes to tracereader adapter ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6705](https://github.com/jaegertracing/jaeger/pull/6705))
* [refactor] conditionally implement interfaces in `v1adapter` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6701](https://github.com/jaegertracing/jaeger/pull/6701))
* [refactor] use `gettracestorefactory` instead of `getstoragefactory` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6696](https://github.com/jaegertracing/jaeger/pull/6696))
* [storage] add helper to storage extension for retrieving purger ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6694](https://github.com/jaegertracing/jaeger/pull/6694))
* Import nop receiver/exporter and add a sample query service config ([@danish9039](https://github.com/danish9039) in [#6687](https://github.com/jaegertracing/jaeger/pull/6687))
* [storage] add helper to storage extension for retrieving sampling store factory ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6689](https://github.com/jaegertracing/jaeger/pull/6689))

#### 👷 CI Improvements

* [idl check] fetch tags ([@yurishkuro](https://github.com/yurishkuro) in [#6758](https://github.com/jaegertracing/jaeger/pull/6758))
* [test]: check for jaeger-idl version mismatch ([@ary82](https://github.com/ary82) in [#6753](https://github.com/jaegertracing/jaeger/pull/6753))
* Allow dependency-review workflow to run from merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#6729](https://github.com/jaegertracing/jaeger/pull/6729))
* Do not run dco-check from merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#6727](https://github.com/jaegertracing/jaeger/pull/6727))
* Do not run label check from merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#6726](https://github.com/jaegertracing/jaeger/pull/6726))
* Allow ci workflows to run from merge queue ([@danish9039](https://github.com/danish9039) in [#6719](https://github.com/jaegertracing/jaeger/pull/6719))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Replace react-vis with recharts ([@hari45678](https://github.com/hari45678) in [#2679](https://github.com/jaegertracing/jaeger-ui/pull/2679))
* Add config option to allow displaying full traceid ([@avinpy-255](https://github.com/avinpy-255) in [#2536](https://github.com/jaegertracing/jaeger-ui/pull/2536))



v1.66.0 / v2.3.0 (2025-02-03)
-------------------------------
#### ⛔ Breaking Changes

* [refactor] remove archive reader and writer from remote storage grpc handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6611](https://github.com/jaegertracing/jaeger/pull/6611))
* Delete grpc metricsqueryservice, metricsquery.proto and related code ([@yurishkuro](https://github.com/yurishkuro) in [#6616](https://github.com/jaegertracing/jaeger/pull/6616))
* [storage] remove distinction between primary and `archive` storage interfaces ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6567](https://github.com/jaegertracing/jaeger/pull/6567))
* [v2][query] create archive reader/writer using regular factory methods ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6519](https://github.com/jaegertracing/jaeger/pull/6519))

#### 🐞 Bug fixes, Minor Improvements

* [fix] replace deprecated address field in service::telemetry ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6679](https://github.com/jaegertracing/jaeger/pull/6679))
* [fix] change metrics port in kafka ingester config to avoid conflict with collector ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6678](https://github.com/jaegertracing/jaeger/pull/6678))
* Update elasticsearch article link ([@timyip3](https://github.com/timyip3) in [#6662](https://github.com/jaegertracing/jaeger/pull/6662))
* [chore] move scylladb implementation to `docker-compose` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6652](https://github.com/jaegertracing/jaeger/pull/6652))
* [fix] refactor archive storage initialization and remove error log ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6636](https://github.com/jaegertracing/jaeger/pull/6636))
* Update import paths for jaeger thrift files to use jaeger-idl ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6635](https://github.com/jaegertracing/jaeger/pull/6635))
* [v2][query] apply "max clock skew adjust" setting ([@dnaka91](https://github.com/dnaka91) in [#6566](https://github.com/jaegertracing/jaeger/pull/6566))
* Alias samping.thrift and clean thrift files ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6630](https://github.com/jaegertracing/jaeger/pull/6630))
* Fix(hotrod): include ca certificates for hotrod dockerfile ([@prashant-shahi](https://github.com/prashant-shahi) in [#6627](https://github.com/jaegertracing/jaeger/pull/6627))
* Replace all imports of jaeger/thrift-gen/* with jaeger-idl/thrift-gen/* ([@danish9039](https://github.com/danish9039) in [#6621](https://github.com/jaegertracing/jaeger/pull/6621))
* Redefine thrift-gen types as aliases to jaeger-idl ([@danish9039](https://github.com/danish9039) in [#6619](https://github.com/jaegertracing/jaeger/pull/6619))
* Add 'features' command to print available feature gates ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#6542](https://github.com/jaegertracing/jaeger/pull/6542))
* Replace jaeger_image_tag with jaeger_version ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#6614](https://github.com/jaegertracing/jaeger/pull/6614))
* Use jeager-idl/proto-gen/api_v2 ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6609](https://github.com/jaegertracing/jaeger/pull/6609))
* Additional model/ cleanup ([@yurishkuro](https://github.com/yurishkuro) in [#6610](https://github.com/jaegertracing/jaeger/pull/6610))
* Return 400 instead of 500 on incorrect otlp payload ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#6599](https://github.com/jaegertracing/jaeger/pull/6599))
* Replace model imports with jaeger-idl ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6595](https://github.com/jaegertracing/jaeger/pull/6595))
* Redefine model/ and api_v2/ types as aliases to jaeger-idl/ types ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6602](https://github.com/jaegertracing/jaeger/pull/6602))
* Add example of es/os server_urls to configs ([@yurishkuro](https://github.com/yurishkuro) in [#6601](https://github.com/jaegertracing/jaeger/pull/6601))
* Sanitize cassandra version before use it ([@rubenvp8510](https://github.com/rubenvp8510) in [#6596](https://github.com/jaegertracing/jaeger/pull/6596))
* Feat: add esmapping-generator into jaeger binary ([@Rohanraj123](https://github.com/Rohanraj123) in [#6327](https://github.com/jaegertracing/jaeger/pull/6327))
* Add replication parameter to cassandra schema script ([@asimchoudhary](https://github.com/asimchoudhary) in [#6582](https://github.com/jaegertracing/jaeger/pull/6582))
* Exclude idl/ as a source of go code ([@yurishkuro](https://github.com/yurishkuro) in [#6591](https://github.com/jaegertracing/jaeger/pull/6591))
* Change model.tootelxxxid() to accept id argument ([@yurishkuro](https://github.com/yurishkuro) in [#6589](https://github.com/jaegertracing/jaeger/pull/6589))
* [refactor][storage][badger]refactored the prefilling of cache to reader ([@Manik2708](https://github.com/Manik2708) in [#6575](https://github.com/jaegertracing/jaeger/pull/6575))
* Move span.getsamplerparams out of model/ into sampling/aggregator ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6583](https://github.com/jaegertracing/jaeger/pull/6583))
* Remove logger parameter in adaptive/aggregator.go ([@Nabil-Salah](https://github.com/Nabil-Salah) in [#6586](https://github.com/jaegertracing/jaeger/pull/6586))
* Separate model parts into more independent pieces ([@yurishkuro](https://github.com/yurishkuro) in [#6581](https://github.com/jaegertracing/jaeger/pull/6581))
* [storage]generate mocks for dependency writer of v2 ([@Manik2708](https://github.com/Manik2708) in [#6576](https://github.com/jaegertracing/jaeger/pull/6576))
* [chore] remove unused method from grpc handler ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6580](https://github.com/jaegertracing/jaeger/pull/6580))
* Document usage of feature gates for breaking changes ([@yurishkuro](https://github.com/yurishkuro) in [#6568](https://github.com/jaegertracing/jaeger/pull/6568))
* [refactor] move sampling strategy providers to internal/sampling/samplingstrategy ([@ary82](https://github.com/ary82) in [#6561](https://github.com/jaegertracing/jaeger/pull/6561))
* [v2][storage] implement reverse adapter to translate v2 writer to v1 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6555](https://github.com/jaegertracing/jaeger/pull/6555))
* [refactor] move sampling strategy interfaces to internal/sampling/strategy ([@ary82](https://github.com/ary82) in [#6547](https://github.com/jaegertracing/jaeger/pull/6547))
* Switch v1 receivers to use v2 write path ([@yurishkuro](https://github.com/yurishkuro) in [#6532](https://github.com/jaegertracing/jaeger/pull/6532))
* [refactor] move plugin/sampling/leaderelection to internal/leaderelection ([@ary82](https://github.com/ary82) in [#6546](https://github.com/jaegertracing/jaeger/pull/6546))
* [refactor] move sampling http handler to internal/sampling/http ([@ary82](https://github.com/ary82) in [#6545](https://github.com/jaegertracing/jaeger/pull/6545))
* [storage] remove dependency on archive flag in es reader ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6490](https://github.com/jaegertracing/jaeger/pull/6490))
* [refactor] move sampling grpc handler to internal/sampling/grpc ([@ary82](https://github.com/ary82) in [#6540](https://github.com/jaegertracing/jaeger/pull/6540))
* Correct references in cmd readme.md ([@jyoungs](https://github.com/jyoungs) in [#6539](https://github.com/jaegertracing/jaeger/pull/6539))
* Use jaeger-v2 by default in hotrod and monitor examples ([@zzzk1](https://github.com/zzzk1) in [#6523](https://github.com/jaegertracing/jaeger/pull/6523))
* Pass context through span processors ([@yurishkuro](https://github.com/yurishkuro) in [#6534](https://github.com/jaegertracing/jaeger/pull/6534))

#### 👷 CI Improvements

* Upgrade storage integration test: use v2 archive readerwriter ([@ekefan](https://github.com/ekefan) in [#6489](https://github.com/jaegertracing/jaeger/pull/6489))
* [chore][tests] clean up integration tests to remove archive reader / writer ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6625](https://github.com/jaegertracing/jaeger/pull/6625))
* Bump jaeger-idl ([@yurishkuro](https://github.com/yurishkuro) in [#6569](https://github.com/jaegertracing/jaeger/pull/6569))
* [storage]upgraded integration tests to use dependency writer of storage_v2 ([@Manik2708](https://github.com/Manik2708) in [#6559](https://github.com/jaegertracing/jaeger/pull/6559))
* [ci] fix binary-size-check workflow ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6552](https://github.com/jaegertracing/jaeger/pull/6552))
* [ci] scrape and verify metrics at the end of e2e tests ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6330](https://github.com/jaegertracing/jaeger/pull/6330))
* [ci] add workflow to guard against increases in the binary size ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6529](https://github.com/jaegertracing/jaeger/pull/6529))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Remove defaultprops from minimap.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2615](https://github.com/jaegertracing/jaeger-ui/pull/2615))
* Remove defaultprops from scatterplot.jsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2618](https://github.com/jaegertracing/jaeger-ui/pull/2618))
* Migrate empasizednode from class based to function based component ([@AdiIsHappy](https://github.com/AdiIsHappy) in [#2638](https://github.com/jaegertracing/jaeger-ui/pull/2638))
* Remove defaultprops from accordiantext.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2612](https://github.com/jaegertracing/jaeger-ui/pull/2612))
* Remove defaultprops from ticks.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2617](https://github.com/jaegertracing/jaeger-ui/pull/2617))
* Remove defaultprops from timelinerow.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2616](https://github.com/jaegertracing/jaeger-ui/pull/2616))
* Remove defaultprops from traceheader.jsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2620](https://github.com/jaegertracing/jaeger-ui/pull/2620))
* Remove defaultprops from accordianlogs.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2613](https://github.com/jaegertracing/jaeger-ui/pull/2613))
* Remove defaultprops fromt breakabletext.tsx ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2611](https://github.com/jaegertracing/jaeger-ui/pull/2611))
* Remove defaultprops from errormessage & newwindowicon ([@ADI-ROXX](https://github.com/ADI-ROXX) in [#2609](https://github.com/jaegertracing/jaeger-ui/pull/2609))
* [loadingindicator]: replace defaultprops with destructuring ([@its-me-abhishek](https://github.com/its-me-abhishek) in [#2601](https://github.com/jaegertracing/jaeger-ui/pull/2601))
* [fix]: disable submit button on invalid minduration or maxduration ([@hari45678](https://github.com/hari45678) in [#2600](https://github.com/jaegertracing/jaeger-ui/pull/2600))
* [deps]: remove dependency on redux-form ([@hari45678](https://github.com/hari45678) in [#2593](https://github.com/jaegertracing/jaeger-ui/pull/2593))
* [fix]: remove redux-form dependency from sort selector ([@hari45678](https://github.com/hari45678) in [#2569](https://github.com/jaegertracing/jaeger-ui/pull/2569))
* [revert]: revert redux and react-redux dependency upgrades ([@yurishkuro](https://github.com/yurishkuro) in [#2577](https://github.com/jaegertracing/jaeger-ui/pull/2577))
* Fix: deep clone trace data for consistency ([@Zen-cronic](https://github.com/Zen-cronic) in [#2571](https://github.com/jaegertracing/jaeger-ui/pull/2571))
* [fix]: remove redux-form dependency from monitor page ([@hari45678](https://github.com/hari45678) in [#2562](https://github.com/jaegertracing/jaeger-ui/pull/2562))
* Fix tracediff graph pan and zoom issue ([@anshgoyalevil](https://github.com/anshgoyalevil) in [#2566](https://github.com/jaegertracing/jaeger-ui/pull/2566))

#### 👷 CI Improvements

* Remove unused matrix from codeql workflow ([@yurishkuro](https://github.com/yurishkuro) in [#2635](https://github.com/jaegertracing/jaeger-ui/pull/2635))
* Rename dco->dco check ([@yurishkuro](https://github.com/yurishkuro) in [#2633](https://github.com/jaegertracing/jaeger-ui/pull/2633))
* Add fake dco check for merge queue events ([@yurishkuro](https://github.com/yurishkuro) in [#2632](https://github.com/jaegertracing/jaeger-ui/pull/2632))
* Don't run label check in merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#2631](https://github.com/jaegertracing/jaeger-ui/pull/2631))
* Don't run codeql from merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#2630](https://github.com/jaegertracing/jaeger-ui/pull/2630))
* Enable workflows to run in merge queue ([@yurishkuro](https://github.com/yurishkuro) in [#2629](https://github.com/jaegertracing/jaeger-ui/pull/2629))
* [ci] fix cache resolution and syntax in check binary size workflow ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#2591](https://github.com/jaegertracing/jaeger-ui/pull/2591))
* [ci]: add workflow to guard against growing bundle size ([@hari45678](https://github.com/hari45678) in [#2586](https://github.com/jaegertracing/jaeger-ui/pull/2586))

v1.65.0 / v2.2.0 (2025-01-08)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* [sampling] inherit default per-operation strategies ([@yurishkuro](https://github.com/yurishkuro) in [#6441](https://github.com/jaegertracing/jaeger/pull/6441))
* [query] enable trace adjusters in api_v2 and api_v3 handlers ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6423](https://github.com/jaegertracing/jaeger/pull/6423))

#### ✨ New Features

* [feat] add time window for gettrace in span store interface ([@rim99](https://github.com/rim99) in [#6242](https://github.com/jaegertracing/jaeger/pull/6242))

#### 🐞 Bug fixes, Minor Improvements

* Return errors from span processor creation ([@yurishkuro](https://github.com/yurishkuro) in [#6488](https://github.com/jaegertracing/jaeger/pull/6488))
* Change collector's queue to use generics ([@yurishkuro](https://github.com/yurishkuro) in [#6486](https://github.com/jaegertracing/jaeger/pull/6486))
* Refactor collector pipeline to allow v1/v2 data model ([@yurishkuro](https://github.com/yurishkuro) in [#6484](https://github.com/jaegertracing/jaeger/pull/6484))
* [v2][storage] implement reverse adapter to translate v2 storage api to v1 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6485](https://github.com/jaegertracing/jaeger/pull/6485))
* [refractor] remove dependency on tlscfg.options ([@Saumya40-codes](https://github.com/Saumya40-codes) in [#6478](https://github.com/jaegertracing/jaeger/pull/6478))
* [query] update v1 query service to check for adapter at construction ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6482](https://github.com/jaegertracing/jaeger/pull/6482))
* [api_v3][query] change api_v3 http handler to use v2 query service ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6459](https://github.com/jaegertracing/jaeger/pull/6459))
* [api_v3][query] change api_v3 grpc handler to use query service v2 ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6452](https://github.com/jaegertracing/jaeger/pull/6452))
* [v2][storage] create v2 query service to operate on otlp data model ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6343](https://github.com/jaegertracing/jaeger/pull/6343))
* Support sampling file reload interval ([@yurishkuro](https://github.com/yurishkuro) in [#6440](https://github.com/jaegertracing/jaeger/pull/6440))
* [jptrace] add spaniter helper function ([@yurishkuro](https://github.com/yurishkuro) in [#6407](https://github.com/jaegertracing/jaeger/pull/6407))
* [refactor][query] propagate `rawtraces` flag to  query service ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6438](https://github.com/jaegertracing/jaeger/pull/6438))
* [v1][adjuster] change v1 adjuster interface to not return error and modify trace in place ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6426](https://github.com/jaegertracing/jaeger/pull/6426))
* [chore] move es/spanstore/dbmodel to internal directory ([@zzzk1](https://github.com/zzzk1) in [#6428](https://github.com/jaegertracing/jaeger/pull/6428))
* [refactor] move model<->otlp translation from `jptrace` to `v1adapter` ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6414](https://github.com/jaegertracing/jaeger/pull/6414))
* Enable udp ports in all-in-one ([@yurishkuro](https://github.com/yurishkuro) in [#6413](https://github.com/jaegertracing/jaeger/pull/6413))
* [refactor] introduce helper for creating map of spans ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6406](https://github.com/jaegertracing/jaeger/pull/6406))
* [fix] fix incorrect usage of iter package in aggregator ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6403](https://github.com/jaegertracing/jaeger/pull/6403))
* [v2][query] implement helper to buffer sequence of traces ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6401](https://github.com/jaegertracing/jaeger/pull/6401))
* [v2][adjuster] implement model to otlp translator with post processing ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6397](https://github.com/jaegertracing/jaeger/pull/6397))
* [v2][adjuster] implement function to get standard adjusters to operate on otlp format ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6396](https://github.com/jaegertracing/jaeger/pull/6396))
* [v2][adjuster] implement otlp to model translator with post processing ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6394](https://github.com/jaegertracing/jaeger/pull/6394))
* [v2][adjuster] implement adjuster for correct timestamps for clockskew ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6392](https://github.com/jaegertracing/jaeger/pull/6392))
* [v2][adjuster] implement adjuster for deduplicating spans ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6391](https://github.com/jaegertracing/jaeger/pull/6391))
* Add optional time window for gettrace & searchtrace of http_handler ([@rim99](https://github.com/rim99) in [#6159](https://github.com/jaegertracing/jaeger/pull/6159))
* [v2][adjuster] implement adjuster for sorting attributes and events ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6389](https://github.com/jaegertracing/jaeger/pull/6389))
* Support extra custom query parameters in requests to prometheus backend ([@akstron](https://github.com/akstron) in [#6360](https://github.com/jaegertracing/jaeger/pull/6360))
* [v2][adjuster] remove error return from adjuster interface ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6383](https://github.com/jaegertracing/jaeger/pull/6383))
* [fix][query] filter out tracing for access to static ui assets ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6374](https://github.com/jaegertracing/jaeger/pull/6374))
* [v2][adjuster] implement span id uniquifier adjuster to operate on otlp data model ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6367](https://github.com/jaegertracing/jaeger/pull/6367))
* [api_v3] add time window for gettrace http_gateway ([@rim99](https://github.com/rim99) in [#6372](https://github.com/jaegertracing/jaeger/pull/6372))
* [v2][adjuster] add warning to span links adjuster ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6381](https://github.com/jaegertracing/jaeger/pull/6381))
* Feat: add time window for gettrace of anonymizer ([@rim99](https://github.com/rim99) in [#6368](https://github.com/jaegertracing/jaeger/pull/6368))
* [v2][adjuster] rework adjuster interface and refactor adjusters to return implemented struct ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6362](https://github.com/jaegertracing/jaeger/pull/6362))
* [v2][adjuster] implement otel attribute adjuster to operate on otlp data model ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6358](https://github.com/jaegertracing/jaeger/pull/6358))
* Respond correctly to stream send error ([@yurishkuro](https://github.com/yurishkuro) in [#6357](https://github.com/jaegertracing/jaeger/pull/6357))
* [v2][adjuster] implement ip attribute adjuster to operate on otlp data model ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6355](https://github.com/jaegertracing/jaeger/pull/6355))
* Remove tls loading and replace with otel configtls ([@yurishkuro](https://github.com/yurishkuro) in [#6345](https://github.com/jaegertracing/jaeger/pull/6345))
* [jaeger][v2] implement span links adjuster to operate on otlp data model ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6354](https://github.com/jaegertracing/jaeger/pull/6354))
* [remote-storage] use otel helper instead of tlscfg ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6351](https://github.com/jaegertracing/jaeger/pull/6351))
* Add go leak check for badgerstore, grpc and memstore e2e tests ([@Manik2708](https://github.com/Manik2708) in [#6347](https://github.com/jaegertracing/jaeger/pull/6347))
* [v2][query] add interface for adjuster to operate on otlp data format ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6346](https://github.com/jaegertracing/jaeger/pull/6346))

#### 🚧 Experimental Features

* Change storage_v2 gettrace to gettraces plural ([@yurishkuro](https://github.com/yurishkuro) in [#6361](https://github.com/jaegertracing/jaeger/pull/6361))
* Change storage v2 api to use streaming ([@yurishkuro](https://github.com/yurishkuro) in [#6359](https://github.com/jaegertracing/jaeger/pull/6359))

#### 👷 CI Improvements

* Upgrade storage integration tests: `dependencyreader` to v2 ([@zzzk1](https://github.com/zzzk1) in [#6477](https://github.com/jaegertracing/jaeger/pull/6477))
* Move remaining util scripts ([@danish9039](https://github.com/danish9039) in [#6472](https://github.com/jaegertracing/jaeger/pull/6472))
* Move lint scripts to scripts/lint ([@danish9039](https://github.com/danish9039) in [#6449](https://github.com/jaegertracing/jaeger/pull/6449))
* Move util scripts to scripts/util ([@danish9039](https://github.com/danish9039) in [#6463](https://github.com/jaegertracing/jaeger/pull/6463))
* Upgrade storage integration test: use `tracewriter` ([@ekefan](https://github.com/ekefan) in [#6437](https://github.com/jaegertracing/jaeger/pull/6437))
* Move e2e scripts to scripts/e2e ([@danish9039](https://github.com/danish9039) in [#6448](https://github.com/jaegertracing/jaeger/pull/6448))
* Move build scripts under scripts/build/ ([@danish9039](https://github.com/danish9039) in [#6446](https://github.com/jaegertracing/jaeger/pull/6446))
* Replace apiv2 with apiv3 client in e2e tests ([@yurishkuro](https://github.com/yurishkuro) in [#6424](https://github.com/jaegertracing/jaeger/pull/6424))
* Do not test with kafka 2.x ([@yurishkuro](https://github.com/yurishkuro) in [#6427](https://github.com/jaegertracing/jaeger/pull/6427))
* Upgrade storage integration test to v2 trace reader ([@ekefan](https://github.com/ekefan) in [#6388](https://github.com/jaegertracing/jaeger/pull/6388))
* Enhance kafka integration tests to support multiple kafka versions ([@zzzk1](https://github.com/zzzk1) in [#6400](https://github.com/jaegertracing/jaeger/pull/6400))
* [fix] fix test expectations for translator to avoid depending on order ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6404](https://github.com/jaegertracing/jaeger/pull/6404))


### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* [clean-up]: remove deprecated plexus/directedgraph ([@hari45678](https://github.com/hari45678) in [#2548](https://github.com/jaegertracing/jaeger-ui/pull/2548))
* [fix]: make plexus demo work again ([@hari45678](https://github.com/hari45678) in [#2538](https://github.com/jaegertracing/jaeger-ui/pull/2538))
* Upgrade from raven-js to sentry/browser ([@avinpy-255](https://github.com/avinpy-255) in [#2509](https://github.com/jaegertracing/jaeger-ui/pull/2509))

v1.64.0 / v2.1.0 (2024-12-06)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* [metrics][storage] move metrics reader decorator to metrics storage factory ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6287](https://github.com/jaegertracing/jaeger/pull/6287))
* [v2][storage] move span reader decorator to storage factories ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6280](https://github.com/jaegertracing/jaeger/pull/6280))

#### ✨ New Features

* [v2][storage] implement read path for v2 storage interface ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6170](https://github.com/jaegertracing/jaeger/pull/6170))
* Create cassandra db schema on session initialization ([@akstron](https://github.com/akstron) in [#5922](https://github.com/jaegertracing/jaeger/pull/5922))

#### 🐞 Bug fixes, Minor Improvements

* Fix password in integration test ([@akstron](https://github.com/akstron) in [#6284](https://github.com/jaegertracing/jaeger/pull/6284))
* [cassandra] change compaction window default to 2hrs ([@yurishkuro](https://github.com/yurishkuro) in [#6282](https://github.com/jaegertracing/jaeger/pull/6282))
* Improve telemetry.settings ([@yurishkuro](https://github.com/yurishkuro) in [#6275](https://github.com/jaegertracing/jaeger/pull/6275))
* [kafka] otel helper instead of tlscfg package ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6270](https://github.com/jaegertracing/jaeger/pull/6270))
* [refactor] fix package misspelling: telemetery->telemetry ([@yurishkuro](https://github.com/yurishkuro) in [#6269](https://github.com/jaegertracing/jaeger/pull/6269))
* [prometheus] use otel helper instead of tlscfg package ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6266](https://github.com/jaegertracing/jaeger/pull/6266))
* [fix] use metrics decorator around metricstorage ([@yurishkuro](https://github.com/yurishkuro) in [#6262](https://github.com/jaegertracing/jaeger/pull/6262))
* Use real metrics factory instead of nullfactory ([@yurishkuro](https://github.com/yurishkuro) in [#6261](https://github.com/jaegertracing/jaeger/pull/6261))
* [v2] use only version number for buildinfo ([@yurishkuro](https://github.com/yurishkuro) in [#6260](https://github.com/jaegertracing/jaeger/pull/6260))
* [refactor] move spm v2 config to cmd/jaeger/ with all other configs ([@yurishkuro](https://github.com/yurishkuro) in [#6256](https://github.com/jaegertracing/jaeger/pull/6256))
* [es-index-cleaner] use otel helper instead of tlscfg ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6259](https://github.com/jaegertracing/jaeger/pull/6259))
* [api_v2] change time fields in archivetracerequest to non-nullable ([@rim99](https://github.com/rim99) in [#6251](https://github.com/jaegertracing/jaeger/pull/6251))
* [es-rollover] use otel helpers for tls config instead of tlscfg ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6238](https://github.com/jaegertracing/jaeger/pull/6238))
* Enable usestdlibvars linter ([@mmorel-35](https://github.com/mmorel-35) in [#6249](https://github.com/jaegertracing/jaeger/pull/6249))
* [storage_v1] add time window to gettracerequest ([@rim99](https://github.com/rim99) in [#6244](https://github.com/jaegertracing/jaeger/pull/6244))
* [fix][query] fix misconfiguration in tls settings from using otel http helper ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6239](https://github.com/jaegertracing/jaeger/pull/6239))
* Auto-generate gogo annotations for api_v3 ([@yurishkuro](https://github.com/yurishkuro) in [#6233](https://github.com/jaegertracing/jaeger/pull/6233))
* Use confighttp in expvar extension ([@yurishkuro](https://github.com/yurishkuro) in [#6227](https://github.com/jaegertracing/jaeger/pull/6227))
* Parameterize listen host and override when in container ([@yurishkuro](https://github.com/yurishkuro) in [#6231](https://github.com/jaegertracing/jaeger/pull/6231))
* Remove 0.0.0.0 overrides in hotrod ci ([@yurishkuro](https://github.com/yurishkuro) in [#6226](https://github.com/jaegertracing/jaeger/pull/6226))
* [storage][v2] add reader adapter that just exposes the underlying v1 reader ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6221](https://github.com/jaegertracing/jaeger/pull/6221))
* Change start/end time in gettrace request to not be pointers ([@yurishkuro](https://github.com/yurishkuro) in [#6218](https://github.com/jaegertracing/jaeger/pull/6218))
* Pass real meterprovider to components ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6173](https://github.com/jaegertracing/jaeger/pull/6173))
* [v2] update versions in readme ([@yurishkuro](https://github.com/yurishkuro) in [#6206](https://github.com/jaegertracing/jaeger/pull/6206))
* Fix: testcreatecollectorproxy unit test failing on go-tip ([@Saumya40-codes](https://github.com/Saumya40-codes) in [#6204](https://github.com/jaegertracing/jaeger/pull/6204))
* Respect environment variables when creating internal tracer ([@akstron](https://github.com/akstron) in [#6179](https://github.com/jaegertracing/jaeger/pull/6179))

#### 🚧 Experimental Features

* [v2]add script for metrics markdown table ([@vvs-personalstash](https://github.com/vvs-personalstash) in [#5941](https://github.com/jaegertracing/jaeger/pull/5941))

#### 👷 CI Improvements

* Allow using different container runtime ([@rim99](https://github.com/rim99) in [#6247](https://github.com/jaegertracing/jaeger/pull/6247))
* K8s integration test for hotrod ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6155](https://github.com/jaegertracing/jaeger/pull/6155))
* Pass username/password to cassandra docker-compose health check ([@akstron](https://github.com/akstron) in [#6214](https://github.com/jaegertracing/jaeger/pull/6214))
* [fix][ci] change the prometheus healthcheck endpoint ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6217](https://github.com/jaegertracing/jaeger/pull/6217))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Add new formatting function "add" ([@drewcorlin1](https://github.com/drewcorlin1) in [#2507](https://github.com/jaegertracing/jaeger-ui/pull/2507))
* Add pad_start link formatting function #2505 ([@drewcorlin1](https://github.com/drewcorlin1) in [#2504](https://github.com/jaegertracing/jaeger-ui/pull/2504))
* Allow formatting link parameter values as iso date #2487 ([@drewcorlin1](https://github.com/drewcorlin1) in [#2501](https://github.com/jaegertracing/jaeger-ui/pull/2501))


v1.63.0 / v2.0.0 (2024-11-10)
-------------------------------

Jaeger v2 is here! 🎉 🎉 🎉

### Backend Changes

#### ⛔ Breaking Changes

* Remove jaeger-agent from distributions ([@yurishkuro](https://github.com/yurishkuro) in [#6081](https://github.com/jaegertracing/jaeger/pull/6081))

#### 🐞 Bug fixes, Minor Improvements

* Fix possible null pointer deference ([@vaidikcode](https://github.com/vaidikcode) in [#6184](https://github.com/jaegertracing/jaeger/pull/6184))
* Chore: enable all rules of perfsprint linter ([@mmorel-35](https://github.com/mmorel-35) in [#6164](https://github.com/jaegertracing/jaeger/pull/6164))
* Chore: enable err-error and errorf rules from perfsprint linter ([@mmorel-35](https://github.com/mmorel-35) in [#6160](https://github.com/jaegertracing/jaeger/pull/6160))
* [query] move trace handler to server level ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6147](https://github.com/jaegertracing/jaeger/pull/6147))
* [fix][query] remove bifurcation for grpc query server ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6145](https://github.com/jaegertracing/jaeger/pull/6145))
* [jaeger-v2] add hotrod integration test for jaeger-v2 ([@Saumya40-codes](https://github.com/Saumya40-codes) in [#6138](https://github.com/jaegertracing/jaeger/pull/6138))
* [query] use otel's helpers for http server ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6121](https://github.com/jaegertracing/jaeger/pull/6121))
* Use grpc interceptors instead of explicit context wrappers ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6133](https://github.com/jaegertracing/jaeger/pull/6133))
* Fix command in v2 example ([@haoqixu](https://github.com/haoqixu) in [#6134](https://github.com/jaegertracing/jaeger/pull/6134))
* Fix span deduplication via correct ordering of adjusters ([@cdanis](https://github.com/cdanis) in [#6116](https://github.com/jaegertracing/jaeger/pull/6116))
* Move all query service http handlers into one function ([@yurishkuro](https://github.com/yurishkuro) in [#6128](https://github.com/jaegertracing/jaeger/pull/6128))
* [fix][grpc] disable tracing in grpc storage writer clients ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6125](https://github.com/jaegertracing/jaeger/pull/6125))
* Feat: automatically publish readme to docker hub ([@inosmeet](https://github.com/inosmeet) in [#6118](https://github.com/jaegertracing/jaeger/pull/6118))
* Use grpc interceptors for bearer token ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6063](https://github.com/jaegertracing/jaeger/pull/6063))
* [fix][query] correct query server legacy condition ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6120](https://github.com/jaegertracing/jaeger/pull/6120))
* [query] use otel's helpers for grpc server ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6055](https://github.com/jaegertracing/jaeger/pull/6055))
* Enable lint rule: import-shadowing ([@inosmeet](https://github.com/inosmeet) in [#6102](https://github.com/jaegertracing/jaeger/pull/6102))
* [refractor] switch to enums for es mappings ([@Saumya40-codes](https://github.com/Saumya40-codes) in [#6091](https://github.com/jaegertracing/jaeger/pull/6091))
* Fix rebuild-ui.sh script ([@andreasgerstmayr](https://github.com/andreasgerstmayr) in [#6098](https://github.com/jaegertracing/jaeger/pull/6098))
* Use otel component host instead of no op host for prod code ([@chahatsagarmain](https://github.com/chahatsagarmain) in [#6085](https://github.com/jaegertracing/jaeger/pull/6085))
* [cassandra] prevent fallback to old schema for operation names table in case of db issues ([@arunvelsriram](https://github.com/arunvelsriram) in [#6061](https://github.com/jaegertracing/jaeger/pull/6061))

#### 🚧 Experimental Features

* Add otlp json support for kafka e2e integration tests ([@joeyyy09](https://github.com/joeyyy09) in [#5935](https://github.com/jaegertracing/jaeger/pull/5935))
* [v2] add es config comments ([@yurishkuro](https://github.com/yurishkuro) in [#6110](https://github.com/jaegertracing/jaeger/pull/6110))
* [chore][docs] add documentation to elasticsearch configuration ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6103](https://github.com/jaegertracing/jaeger/pull/6103))
* [jaeger-v2] refactor elasticsearch/opensearch configurations to have more logical groupings ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6090](https://github.com/jaegertracing/jaeger/pull/6090))
* [jaeger-v2] implement utf-8 sanitizer for otlp ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6078](https://github.com/jaegertracing/jaeger/pull/6078))
* [jaeger-v2] migrate elasticsearch/opensearch to use otel's tls configuration ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6079](https://github.com/jaegertracing/jaeger/pull/6079))
* [jaeger-v2] enable queueing configuration in storage exporter ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6080](https://github.com/jaegertracing/jaeger/pull/6080))
* [jaeger-v2] implement empty service name sanitizer for otlp ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6077](https://github.com/jaegertracing/jaeger/pull/6077))
* [jaeger-v2] refactor elasticsearch/opensearch storage configurations ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6060](https://github.com/jaegertracing/jaeger/pull/6060))

#### 👷 CI Improvements

* [v2] use health check in grpc e2e test ([@yurishkuro](https://github.com/yurishkuro) in [#6113](https://github.com/jaegertracing/jaeger/pull/6113))
* Update node.js github action to use npm lockfile, switch to latest jaeger ui ([@andreasgerstmayr](https://github.com/andreasgerstmayr) in [#6074](https://github.com/jaegertracing/jaeger/pull/6074))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Migrate from yarn v1 to npm ([@andreasgerstmayr](https://github.com/andreasgerstmayr) in [#2462](https://github.com/jaegertracing/jaeger-ui/pull/2462))

#### 👷 CI Improvements

* Run s390x build on push to main only ([@andreasgerstmayr](https://github.com/andreasgerstmayr) in [#2481](https://github.com/jaegertracing/jaeger-ui/pull/2481))

1.62.0 / 2.0.0-rc2 (2024-10-06)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* [query] change http and tls server configurations to use otel configurations ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6023](https://github.com/jaegertracing/jaeger/pull/6023))
* [fix][spm]: change default metrics namespace to match new default in spanmetricsconnector ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6007](https://github.com/jaegertracing/jaeger/pull/6007))

#### 🐞 Bug fixes, Minor Improvements

* [grpc storage]: propagate tenant to grpc backend ([@frzifus](https://github.com/frzifus) in [#6030](https://github.com/jaegertracing/jaeger/pull/6030))
* [feat] deduplicate spans based on their hashcode ([@cdanis](https://github.com/cdanis) in [#6009](https://github.com/jaegertracing/jaeger/pull/6009))

#### 🚧 Experimental Features

* [jaeger-v2] consolidate v1 and v2 configurations for grpc storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6042](https://github.com/jaegertracing/jaeger/pull/6042))
* [jaeger-v2] use environment variables in kafka config ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6028](https://github.com/jaegertracing/jaeger/pull/6028))
* [jaeger-v2] align cassandra storage config with otel ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5949](https://github.com/jaegertracing/jaeger/pull/5949))
* [jaeger-v2] refactor configuration for query service ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5998](https://github.com/jaegertracing/jaeger/pull/5998))
* [v2] add temporary expvar extension ([@yurishkuro](https://github.com/yurishkuro) in [#5986](https://github.com/jaegertracing/jaeger/pull/5986))

#### 👷 CI Improvements

* [ci] disable fail fast behaviour for ci workflows ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#6052](https://github.com/jaegertracing/jaeger/pull/6052))
* Testifylint: enable go-require ([@mmorel-35](https://github.com/mmorel-35) in [#5983](https://github.com/jaegertracing/jaeger/pull/5983))
* Fix regex for publishing v2 image ([@yurishkuro](https://github.com/yurishkuro) in [#5988](https://github.com/jaegertracing/jaeger/pull/5988))


### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Support uploads of .jsonl files ([@Saumya40-codes](https://github.com/Saumya40-codes) in [#2461](https://github.com/jaegertracing/jaeger-ui/pull/2461))


1.61.0 / 2.0.0-rc1 (2024-09-14)
-------------------------------

### Backend Changes

This release contains an official pre-release candidate of Jaeger v2, as binary and Docker image `jaeger`.

#### ⛔ Breaking Changes

* Remove support for cassandra 3.x and add cassandra 5.x ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5962](https://github.com/jaegertracing/jaeger/pull/5962))

#### 🐞 Bug fixes, Minor Improvements

* Fix: the 'tagtype' in es jaeger-span mapping tags.properties should be 'type' ([@chinaran](https://github.com/chinaran) in [#5980](https://github.com/jaegertracing/jaeger/pull/5980))
* Add readme for adaptive sampling ([@yurishkuro](https://github.com/yurishkuro) in [#5955](https://github.com/jaegertracing/jaeger/pull/5955))
* [adaptive sampling] clean-up after previous refactoring ([@yurishkuro](https://github.com/yurishkuro) in [#5954](https://github.com/jaegertracing/jaeger/pull/5954))
* [adaptive processor] remove redundant function ([@yurishkuro](https://github.com/yurishkuro) in [#5953](https://github.com/jaegertracing/jaeger/pull/5953))
* [jaeger-v2] consolidate options and namespaceconfig for badger storage ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5937](https://github.com/jaegertracing/jaeger/pull/5937))
* Remove unused "namespace" field from badger config ([@yurishkuro](https://github.com/yurishkuro) in [#5929](https://github.com/jaegertracing/jaeger/pull/5929))
* Simplify bundling of ui assets ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5917](https://github.com/jaegertracing/jaeger/pull/5917))
* Clean up grpc storage config ([@yurishkuro](https://github.com/yurishkuro) in [#5877](https://github.com/jaegertracing/jaeger/pull/5877))
* Add script to replace apache headers with spdx ([@thecaffeinedev](https://github.com/thecaffeinedev) in [#5808](https://github.com/jaegertracing/jaeger/pull/5808))
* Add copyright/license headers to script files ([@Zen-cronic](https://github.com/Zen-cronic) in [#5829](https://github.com/jaegertracing/jaeger/pull/5829))
* Clearer output from lint scripts ([@yurishkuro](https://github.com/yurishkuro) in [#5820](https://github.com/jaegertracing/jaeger/pull/5820))

#### 🚧 Experimental Features

* [jaeger-v2] add validation and comments to badger storage config ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5927](https://github.com/jaegertracing/jaeger/pull/5927))
* [jaeger-v2] add validation and comments to memory storage config ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5925](https://github.com/jaegertracing/jaeger/pull/5925))
* Support tail based sampling processor from otel collector extension ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5878](https://github.com/jaegertracing/jaeger/pull/5878))
* [v2] configure health check extension for all configs ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5861](https://github.com/jaegertracing/jaeger/pull/5861))
* [v2] add legacy formats into e2e kafka integration tests ([@joeyyy09](https://github.com/joeyyy09) in [#5824](https://github.com/jaegertracing/jaeger/pull/5824))
* [v2] configure healthcheck extension ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5831](https://github.com/jaegertracing/jaeger/pull/5831))
* Added _total suffix to otel counter metrics. ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5810](https://github.com/jaegertracing/jaeger/pull/5810))

#### 👷 CI Improvements

* Release v2 cleanup 3 ([@yurishkuro](https://github.com/yurishkuro) in [#5984](https://github.com/jaegertracing/jaeger/pull/5984))
* Replace loopvar linter ([@anishbista60](https://github.com/anishbista60) in [#5976](https://github.com/jaegertracing/jaeger/pull/5976))
* Stop using v1 and v1.x tags for docker images ([@yurishkuro](https://github.com/yurishkuro) in [#5956](https://github.com/jaegertracing/jaeger/pull/5956))
* V2 repease prep ([@yurishkuro](https://github.com/yurishkuro) in [#5932](https://github.com/jaegertracing/jaeger/pull/5932))
* Normalize build-binaries targets ([@yurishkuro](https://github.com/yurishkuro) in [#5924](https://github.com/jaegertracing/jaeger/pull/5924))
* Fix integration test log dumping for storage backends ([@mahadzaryab1](https://github.com/mahadzaryab1) in [#5915](https://github.com/jaegertracing/jaeger/pull/5915))
* Add jaeger-v2 binary as new release artifact ([@renovate-bot](https://github.com/renovate-bot) in [#5893](https://github.com/jaegertracing/jaeger/pull/5893))
* [ci] add support for v2 tags during build ([@yurishkuro](https://github.com/yurishkuro) in [#5890](https://github.com/jaegertracing/jaeger/pull/5890))
* Add hardcoded db password and username to cassandra integration test ([@Ali-Alnosairi](https://github.com/Ali-Alnosairi) in [#5805](https://github.com/jaegertracing/jaeger/pull/5805))
* Define contents permissions on "dependabot validate" workflow ([@mmorel-35](https://github.com/mmorel-35) in [#5874](https://github.com/jaegertracing/jaeger/pull/5874))
* [fix] print kafka logs on test failure ([@joeyyy09](https://github.com/joeyyy09) in [#5873](https://github.com/jaegertracing/jaeger/pull/5873))
* Pin github actions dependencies ([@harshitasao](https://github.com/harshitasao) in [#5860](https://github.com/jaegertracing/jaeger/pull/5860))
* Add go.mod for docker debug image ([@hellspawn679](https://github.com/hellspawn679) in [#5852](https://github.com/jaegertracing/jaeger/pull/5852))
* Enable lint rule: redefines-builtin-id ([@ZXYxc](https://github.com/ZXYxc) in [#5791](https://github.com/jaegertracing/jaeger/pull/5791))
* Require manual go version updates for patch versions ([@wasup-yash](https://github.com/wasup-yash) in [#5848](https://github.com/jaegertracing/jaeger/pull/5848))
* Clean up obselete 'version' tag from docker-compose files ([@vvs-personalstash](https://github.com/vvs-personalstash) in [#5826](https://github.com/jaegertracing/jaeger/pull/5826))
* Update expected codecov flags count to 19 ([@yurishkuro](https://github.com/yurishkuro) in [#5811](https://github.com/jaegertracing/jaeger/pull/5811))


### 📊 UI Changes

Dependencies upgrades only.


1.60.0 / 2.0.0-rc0 (2024-08-06)
-------------------------------

### Backend Changes

#### ⛔ Breaking Changes

* Completely remove "grpc-plugin" as storage type ([@yurishkuro](https://github.com/yurishkuro) in [#5741](https://github.com/jaegertracing/jaeger/pull/5741))

#### 🐞 Bug fixes, Minor Improvements

* Do not use image tag without version ([@yurishkuro](https://github.com/yurishkuro) in [#5783](https://github.com/jaegertracing/jaeger/pull/5783))
* Only attach :latest tag to versioned images from main ([@yurishkuro](https://github.com/yurishkuro) in [#5781](https://github.com/jaegertracing/jaeger/pull/5781))
* Add references to jaeger v2 ([@yurishkuro](https://github.com/yurishkuro) in [#5779](https://github.com/jaegertracing/jaeger/pull/5779))
* Ensure hotrod image is published at the end of e2e test ([@yurishkuro](https://github.com/yurishkuro) in [#5764](https://github.com/jaegertracing/jaeger/pull/5764))
* [bug] [hotrod] delay env var mapping until logger is initialized ([@yurishkuro](https://github.com/yurishkuro) in [#5760](https://github.com/jaegertracing/jaeger/pull/5760))
* Make otlp receiver listen on all ips again ([@yurishkuro](https://github.com/yurishkuro) in [#5739](https://github.com/jaegertracing/jaeger/pull/5739))
* [hotrod] fix connectivity in docker compose ([@yurishkuro](https://github.com/yurishkuro) in [#5734](https://github.com/jaegertracing/jaeger/pull/5734))

#### 🚧 Experimental Features

* [v2] enable remote sampling extension and include in e2e tests ([@yurishkuro](https://github.com/yurishkuro) in [#5802](https://github.com/jaegertracing/jaeger/pull/5802))
* Ensure similar naming for storage write metrics ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5798](https://github.com/jaegertracing/jaeger/pull/5798))
* [v2] ensure similar naming for query service metrics ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5785](https://github.com/jaegertracing/jaeger/pull/5785))
* Configure otel collector to observe internal telemetry ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5752](https://github.com/jaegertracing/jaeger/pull/5752))
* Add kafka exporter and receiver configuration ([@joeyyy09](https://github.com/joeyyy09) in [#5703](https://github.com/jaegertracing/jaeger/pull/5703))
* Enable spm in jaeger v2 ([@FlamingSaint](https://github.com/FlamingSaint) in [#5681](https://github.com/jaegertracing/jaeger/pull/5681))
* [jaeger-v2] add `remotesampling` extension ([@Pushkarm029](https://github.com/Pushkarm029) in [#5389](https://github.com/jaegertracing/jaeger/pull/5389))
* Created telset for remote-storage component ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5731](https://github.com/jaegertracing/jaeger/pull/5731))

#### 👷 CI Improvements

* Unpin codeql actions ([@yurishkuro](https://github.com/yurishkuro) in [#5787](https://github.com/jaegertracing/jaeger/pull/5787))
* Skip building hotrod for all platforms for pull requests ([@Manoramsharma](https://github.com/Manoramsharma) in [#5765](https://github.com/jaegertracing/jaeger/pull/5765))
* Add a threshold for expected zero values in the spm script ([@FlamingSaint](https://github.com/FlamingSaint) in [#5753](https://github.com/jaegertracing/jaeger/pull/5753))
* [v2] add e2e test with memory store ([@yurishkuro](https://github.com/yurishkuro) in [#5751](https://github.com/jaegertracing/jaeger/pull/5751))
* Rationalize naming of gha workflow files ([@yurishkuro](https://github.com/yurishkuro) in [#5750](https://github.com/jaegertracing/jaeger/pull/5750))


### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Allow uploading json-per-line otlp data ([@BenzeneAlcohol](https://github.com/BenzeneAlcohol) in [#2380](https://github.com/jaegertracing/jaeger-ui/pull/2380))


1.59.0 (2024-07-10)
-------------------

### Backend Changes

#### ⛔ Breaking Changes

* The OTEL Collector upgrade brought in a change where OTLP receivers started listening on `localhost` instead of `0.0.0.0` as before. As a result, when running in container environment the endpoints are likely unreachable from other containers (Issue [#5737](https://github.com/jaegertracing/jaeger/issues/5737)). The fix will be available in the next release. Meanwhile, the workaround is to instruct Jaeger to listen on `0.0.0.0`, as in [this fix](https://github.com/jaegertracing/jaeger/pull/5734/files#diff-299f817cc4ab077ddb763f1e6a023d9d042d714e2fd3736cc40af3f218d44f1eR15):
```
      - COLLECTOR_OTLP_GRPC_HOST_PORT=0.0.0.0:4317
      - COLLECTOR_OTLP_HTTP_HOST_PORT=0.0.0.0:4318
```
* Update opentelemetry-go to v1.28.0 and refactor references to semantic conventions ([@renovate-bot](https://github.com/renovate-bot) in [#5698](https://github.com/jaegertracing/jaeger/pull/5698))

#### ✨ New Features

* Run  jaeger-es-index-cleaner and jaeger-es-rollover locally ([@hellspawn679](https://github.com/hellspawn679) in [#5714](https://github.com/jaegertracing/jaeger/pull/5714))
* [tracegen] allow use of adaptive sampling ([@yurishkuro](https://github.com/yurishkuro) in [#5718](https://github.com/jaegertracing/jaeger/pull/5718))
* [v2] add v1 factory converter to v2 storage factory ([@james-ryans](https://github.com/james-ryans) in [#5497](https://github.com/jaegertracing/jaeger/pull/5497))
* Upgrade badger  v3->badger  v4 ([@hellspawn679](https://github.com/hellspawn679) in [#5619](https://github.com/jaegertracing/jaeger/pull/5619))

#### 🐞 Bug fixes, Minor Improvements

* Cleanup the prometheus config ([@FlamingSaint](https://github.com/FlamingSaint) in [#5720](https://github.com/jaegertracing/jaeger/pull/5720))
* Upgrade microsim to v0.4.1 ([@FlamingSaint](https://github.com/FlamingSaint) in [#5702](https://github.com/jaegertracing/jaeger/pull/5702))
* Add all mocks to mockery config file and regenerate ([@danish9039](https://github.com/danish9039) in [#5626](https://github.com/jaegertracing/jaeger/pull/5626))
* Add better logging options ([@yurishkuro](https://github.com/yurishkuro) in [#5675](https://github.com/jaegertracing/jaeger/pull/5675))
* Restore "operation" name in the metrics response ([@yurishkuro](https://github.com/yurishkuro) in [#5673](https://github.com/jaegertracing/jaeger/pull/5673))
* Add flag for custom authenticators in cassandra storage ([@hellspawn679](https://github.com/hellspawn679) in [#5628](https://github.com/jaegertracing/jaeger/pull/5628))
* Rename strategy store to sampling strategy provider ([@yurishkuro](https://github.com/yurishkuro) in [#5634](https://github.com/jaegertracing/jaeger/pull/5634))
* [query] avoid errors when closing shared listener ([@vermaaatul07](https://github.com/vermaaatul07) in [#5559](https://github.com/jaegertracing/jaeger/pull/5559))
* Bump github.com/golangci/golangci-lint from 1.55.2 to 1.59.1 and fix linter errors ([@FlamingSaint](https://github.com/FlamingSaint) in [#5579](https://github.com/jaegertracing/jaeger/pull/5579))
* Fix binary path in package-deploy.sh ([@yurishkuro](https://github.com/yurishkuro) in [#5561](https://github.com/jaegertracing/jaeger/pull/5561))

#### 🚧 Experimental Features

* Implement telemetry struct for v1 components initialization ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5695](https://github.com/jaegertracing/jaeger/pull/5695))
* Support default configs for storage backends ([@yurishkuro](https://github.com/yurishkuro) in [#5691](https://github.com/jaegertracing/jaeger/pull/5691))
* Simplify configs organization ([@yurishkuro](https://github.com/yurishkuro) in [#5690](https://github.com/jaegertracing/jaeger/pull/5690))
* Create metrics.factory adapter for otel metrics ([@Wise-Wizard](https://github.com/Wise-Wizard) in [#5661](https://github.com/jaegertracing/jaeger/pull/5661))

#### 👷 CI Improvements

* Apply 'latest' tag to latest published snapshot images ([@yurishkuro](https://github.com/yurishkuro) in [#5724](https://github.com/jaegertracing/jaeger/pull/5724))
* [bug] use correct argument as jaeger-version ([@yurishkuro](https://github.com/yurishkuro) in [#5716](https://github.com/jaegertracing/jaeger/pull/5716))
* Add spm integration tests ([@hellspawn679](https://github.com/hellspawn679) in [#5640](https://github.com/jaegertracing/jaeger/pull/5640))
* Add spm build to ci ([@yurishkuro](https://github.com/yurishkuro) in [#5663](https://github.com/jaegertracing/jaeger/pull/5663))
* Remove unnecessary .nocover files ([@yurishkuro](https://github.com/yurishkuro) in [#5642](https://github.com/jaegertracing/jaeger/pull/5642))
* Add tests for anonymizer/app/query. ([@shanukun](https://github.com/shanukun) in [#5638](https://github.com/jaegertracing/jaeger/pull/5638))
* Add alternate way to install gotip ([@EraKin575](https://github.com/EraKin575) in [#5618](https://github.com/jaegertracing/jaeger/pull/5618))
* Add semver to dependencies ([@danish9039](https://github.com/danish9039) in [#5590](https://github.com/jaegertracing/jaeger/pull/5590))
* Create config file for mockery instead of using explicit cli flags in the makefile ([@jesslourenco](https://github.com/jesslourenco) in [#5623](https://github.com/jaegertracing/jaeger/pull/5623))
* Update renovate bot to not apply patches to e2e test dependencies ([@DustyMMiller](https://github.com/DustyMMiller) in [#5622](https://github.com/jaegertracing/jaeger/pull/5622))
* Require renovate bot to run go mod tidy ([@yurishkuro](https://github.com/yurishkuro) in [#5612](https://github.com/jaegertracing/jaeger/pull/5612))
* Fix new warnings from the linter upgrade ([@WaterLemons2k](https://github.com/WaterLemons2k) in [#5589](https://github.com/jaegertracing/jaeger/pull/5589))
* [ci] validate that generated mocks are up to date ([@yurishkuro](https://github.com/yurishkuro) in [#5568](https://github.com/jaegertracing/jaeger/pull/5568))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Add escaped example to tag search help popup ([@yurishkuro](https://github.com/yurishkuro) in [#2354](https://github.com/jaegertracing/jaeger-ui/pull/2354))



1.58.1 (2024-06-22)
-------------------

### Backend Changes

#### 🐞 Bug fixes, Minor Improvements

* SPM: Restore "operation" name in the metrics response ([@yurishkuro](https://github.com/yurishkuro) in [#5673](https://github.com/jaegertracing/jaeger/pull/5673))


1.58.0 (2024-06-11)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

* Remove support for otel spanmetrics processor ([@varshith257](https://github.com/varshith257) in [#5539](https://github.com/jaegertracing/jaeger/pull/5539))
* Remove support for expvar-backed metrics ([@joeyyy09](https://github.com/joeyyy09) in [#5437](https://github.com/jaegertracing/jaeger/pull/5437))
* Remove support for elasticsearch v5/v6 ([@FlamingSaint](https://github.com/FlamingSaint) in [#5448](https://github.com/jaegertracing/jaeger/pull/5448))
* Remove support for gRPC-Plugin ([@h4shk4t](https://github.com/h4shk4t) in [#5388](https://github.com/jaegertracing/jaeger/pull/5388))

#### 🐞 Bug fixes, Minor Improvements

* Set desired providers/converters instead of relying on defaults ([@TylerHelmuth](https://github.com/TylerHelmuth) in [#5543](https://github.com/jaegertracing/jaeger/pull/5543))
* Add elasticsearch helper binaries to release ([@FlamingSaint](https://github.com/FlamingSaint) in [#5501](https://github.com/jaegertracing/jaeger/pull/5501))
* Replace internal metrics.factory usage with direct calls to expvar ([@prakrit55](https://github.com/prakrit55) in [#5496](https://github.com/jaegertracing/jaeger/pull/5496))
* [refactor] move root span handler into aggregator ([@Pushkarm029](https://github.com/Pushkarm029) in [#5478](https://github.com/jaegertracing/jaeger/pull/5478))
* Refactor adaptive sampling aggregator & strategy store ([@Pushkarm029](https://github.com/Pushkarm029) in [#5441](https://github.com/jaegertracing/jaeger/pull/5441))
* [remote-storage] add healthcheck to grpc server ([@james-ryans](https://github.com/james-ryans) in [#5461](https://github.com/jaegertracing/jaeger/pull/5461))
* Fix alpine image to 3.19.0 ([@prakrit55](https://github.com/prakrit55) in [#5454](https://github.com/jaegertracing/jaeger/pull/5454))
* Replace grpc-plugin storage type name with just grpc ([@yurishkuro](https://github.com/yurishkuro) in [#5442](https://github.com/jaegertracing/jaeger/pull/5442))
* [grpc-storage] use grpc.newclient ([@yurishkuro](https://github.com/yurishkuro) in [#5393](https://github.com/jaegertracing/jaeger/pull/5393))
* Replace public initfromoptions with private configurefromoptions ([@yurishkuro](https://github.com/yurishkuro) in [#5417](https://github.com/jaegertracing/jaeger/pull/5417))
* [jaeger-v2] fix e2e storage integration v0.99.0 otel col failing ([@james-ryans](https://github.com/james-ryans) in [#5419](https://github.com/jaegertracing/jaeger/pull/5419))
* Add purge method for cassandra ([@akagami-harsh](https://github.com/akagami-harsh) in [#5414](https://github.com/jaegertracing/jaeger/pull/5414))
* [v2] add diagrams to the docs ([@yurishkuro](https://github.com/yurishkuro) in [#5412](https://github.com/jaegertracing/jaeger/pull/5412))
* [es] remove unused indexcache ([@yurishkuro](https://github.com/yurishkuro) in [#5408](https://github.com/jaegertracing/jaeger/pull/5408))

#### 🚧 Experimental Features

* Create new grpc storage configuration to align with otel ([@akagami-harsh](https://github.com/akagami-harsh) in [#5331](https://github.com/jaegertracing/jaeger/pull/5331))

#### 👷 CI Improvements

* Add workflow to validate dependabot config ([@yurishkuro](https://github.com/yurishkuro) in [#5556](https://github.com/jaegertracing/jaeger/pull/5556))
* Create separate directories for docker compose files ([@FlamingSaint](https://github.com/FlamingSaint) in [#5554](https://github.com/jaegertracing/jaeger/pull/5554))
* [ci] use m.x version names in workflows ([@hellspawn679](https://github.com/hellspawn679) in [#5546](https://github.com/jaegertracing/jaeger/pull/5546))
* Enable lint rule: unused-parameter ([@FlamingSaint](https://github.com/FlamingSaint) in [#5534](https://github.com/jaegertracing/jaeger/pull/5534))
* Add a new workflow for testing the release workflow ([@yurishkuro](https://github.com/yurishkuro) in [#5532](https://github.com/jaegertracing/jaeger/pull/5532))
* Enable lint rules: struct-tag & unexported-return ([@FlamingSaint](https://github.com/FlamingSaint) in [#5533](https://github.com/jaegertracing/jaeger/pull/5533))
* Enable lint rules: early-return & indent-error-flow ([@FlamingSaint](https://github.com/FlamingSaint) in [#5526](https://github.com/jaegertracing/jaeger/pull/5526))
* Enable lint rule: exported ([@FlamingSaint](https://github.com/FlamingSaint) in [#5525](https://github.com/jaegertracing/jaeger/pull/5525))
* Enable lint rules: confusing-results & receiver-naming ([@FlamingSaint](https://github.com/FlamingSaint) in [#5524](https://github.com/jaegertracing/jaeger/pull/5524))
* Add manual dco check using python script dco-check ([@yurishkuro](https://github.com/yurishkuro) in [#5528](https://github.com/jaegertracing/jaeger/pull/5528))
* Use docker compose for cassandra integration tests ([@hellspawn679](https://github.com/hellspawn679) in [#5520](https://github.com/jaegertracing/jaeger/pull/5520))
* Enable lint rule: modifies-value-receiver ([@FlamingSaint](https://github.com/FlamingSaint) in [#5517](https://github.com/jaegertracing/jaeger/pull/5517))
* Enable lint rule: unused-receiver ([@FlamingSaint](https://github.com/FlamingSaint) in [#5521](https://github.com/jaegertracing/jaeger/pull/5521))
* Enable lint rule: dot-imports ([@FlamingSaint](https://github.com/FlamingSaint) in [#5513](https://github.com/jaegertracing/jaeger/pull/5513))
* Enable lint rules: bare-return & empty-lines ([@FlamingSaint](https://github.com/FlamingSaint) in [#5512](https://github.com/jaegertracing/jaeger/pull/5512))
* Manage 3rd party tools via dedicated go.mod ([@yurishkuro](https://github.com/yurishkuro) in [#5509](https://github.com/jaegertracing/jaeger/pull/5509))
* Enable lint rule: use-any ([@FlamingSaint](https://github.com/FlamingSaint) in [#5510](https://github.com/jaegertracing/jaeger/pull/5510))
* Enable lint rule: unexported-naming ([@FlamingSaint](https://github.com/FlamingSaint) in [#5511](https://github.com/jaegertracing/jaeger/pull/5511))
* Add nolintlint linter ([@yurishkuro](https://github.com/yurishkuro) in [#5508](https://github.com/jaegertracing/jaeger/pull/5508))
* Use docker compose for kafka integration tests ([@hellspawn679](https://github.com/hellspawn679) in [#5500](https://github.com/jaegertracing/jaeger/pull/5500))
* Use docker compose for elasticsearch/opensearch integration tests ([@hellspawn679](https://github.com/hellspawn679) in [#5490](https://github.com/jaegertracing/jaeger/pull/5490))
* Split v1 and v2 es/os integration tests ([@yurishkuro](https://github.com/yurishkuro) in [#5487](https://github.com/jaegertracing/jaeger/pull/5487))
* Remove args and depict the image in from directive ([@prakrit55](https://github.com/prakrit55) in [#5465](https://github.com/jaegertracing/jaeger/pull/5465))
* [v2] remove retries and increase timeout for e2e tests ([@james-ryans](https://github.com/james-ryans) in [#5460](https://github.com/jaegertracing/jaeger/pull/5460))
* Restore code coverage threshold back to 95% ([@varshith257](https://github.com/varshith257) in [#5457](https://github.com/jaegertracing/jaeger/pull/5457))
* [v2] add logging to read/write spans in e2e tests ([@james-ryans](https://github.com/james-ryans) in [#5456](https://github.com/jaegertracing/jaeger/pull/5456))
* Remove elasticsearch v5/v6 from tests ([@FlamingSaint](https://github.com/FlamingSaint) in [#5451](https://github.com/jaegertracing/jaeger/pull/5451))
* [v2] replace e2e span_reader grpc.dialcontext with newclient ([@james-ryans](https://github.com/james-ryans) in [#5443](https://github.com/jaegertracing/jaeger/pull/5443))
* Stop running integration tests for elasticsearch v5/v6 ([@yurishkuro](https://github.com/yurishkuro) in [#5440](https://github.com/jaegertracing/jaeger/pull/5440))
* [v2] remove temporary skipbinaryattrs flag from e2e tests ([@james-ryans](https://github.com/james-ryans) in [#5436](https://github.com/jaegertracing/jaeger/pull/5436))
* [v2] dump storage docker logs in github ci if e2e test failed ([@james-ryans](https://github.com/james-ryans) in [#5433](https://github.com/jaegertracing/jaeger/pull/5433))
* [v2] change e2e jaeger-v2 binary log output to temp file ([@james-ryans](https://github.com/james-ryans) in [#5431](https://github.com/jaegertracing/jaeger/pull/5431))
* [bug] fix syntax for invoking upload-codecov action ([@yurishkuro](https://github.com/yurishkuro) in [#5416](https://github.com/jaegertracing/jaeger/pull/5416))
* Use helper action to retry codecov uploads ([@yurishkuro](https://github.com/yurishkuro) in [#5411](https://github.com/jaegertracing/jaeger/pull/5411))
* Only build docker images for crossdock tests for linux/amd64 ([@varshith257](https://github.com/varshith257) in [#5410](https://github.com/jaegertracing/jaeger/pull/5410))

### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Document how to debug unit tests in vscode ([@RISHIKESHk07](https://github.com/RISHIKESHk07) in [#2297](https://github.com/jaegertracing/jaeger-ui/pull/2297))

#### 👷 CI Improvements

* Github actions added to block prs from fork/main branch ([@varshith257](https://github.com/varshith257) in [#2296](https://github.com/jaegertracing/jaeger-ui/pull/2296))

1.57.0 (2024-05-01)
-------------------

### Backend Changes

#### 🐞 Bug fixes, Minor Improvements

* [jaeger-v2] define an internal interface of storage v2 spanstore ([@james-ryans](https://github.com/james-ryans) in [#5399](https://github.com/jaegertracing/jaeger/pull/5399))
* Combine jaeger ui release notes with jaeger backend ([@albertteoh](https://github.com/albertteoh) in [#5405](https://github.com/jaegertracing/jaeger/pull/5405))
* [agent] use grpc.newclient ([@yurishkuro](https://github.com/yurishkuro) in [#5392](https://github.com/jaegertracing/jaeger/pull/5392))
* [sampling] fix merging of per-operation strategies into service strategies without them ([@kuujis](https://github.com/kuujis) in [#5277](https://github.com/jaegertracing/jaeger/pull/5277))
* Create sampling templates when creating sampling store ([@JaeguKim](https://github.com/JaeguKim) in [#5349](https://github.com/jaegertracing/jaeger/pull/5349))
* [kafka-consumer] set the rackid in consumer config ([@sappusaketh](https://github.com/sappusaketh) in [#5374](https://github.com/jaegertracing/jaeger/pull/5374))
* Adding best practices badge to readme.md ([@jkowall](https://github.com/jkowall) in [#5369](https://github.com/jaegertracing/jaeger/pull/5369))

#### 👷 CI Improvements

* Moving global write permissions down into the ci jobs ([@jkowall](https://github.com/jkowall) in [#5370](https://github.com/jaegertracing/jaeger/pull/5370))


### 📊 UI Changes

#### 🐞 Bug fixes, Minor Improvements

* Improve trace page title with data and unique emoji (fixes #2256) ([@nox](https://github.com/nox) in [#2275](https://github.com/jaegertracing/jaeger-ui/pull/2275))
* Require node version 20+ ([@Baalekshan](https://github.com/Baalekshan) in [#2274](https://github.com/jaegertracing/jaeger-ui/pull/2274))


1.56.0 (2024-04-02)
-------------------

### Backend Changes

#### ⛔ Breaking Changes

* Fix hotrod instructions ([@yurishkuro](https://github.com/yurishkuro) in [#5273](https://github.com/jaegertracing/jaeger/pull/5273))

#### 🐞 Bug fixes, Minor Improvements

* Refactor healthcheck signalling between server and service ([@WillSewell](https://github.com/WillSewell) in [#5308](https://github.com/jaegertracing/jaeger/pull/5308))
* Docs: badger file permission as non-root service ([@tico88612](https://github.com/tico88612) in [#5282](https://github.com/jaegertracing/jaeger/pull/5282))
* [kafka-consumer] add support for setting fetch message max bytes ([@sappusaketh](https://github.com/sappusaketh) in [#5283](https://github.com/jaegertracing/jaeger/pull/5283))
* [chore] remove repetitive words ([@tgolang](https://github.com/tgolang) in [#5265](https://github.com/jaegertracing/jaeger/pull/5265))
* Fix zipkin spanformat ([@fyuan1316](https://github.com/fyuan1316) in [#5261](https://github.com/jaegertracing/jaeger/pull/5261))
* [kafka-producer] support setting max message size ([@sappusaketh](https://github.com/sappusaketh) in [#5263](https://github.com/jaegertracing/jaeger/pull/5263))

#### 🚧 Experimental Features

* [jaeger-v2] add support for opensearch ([@akagami-harsh](https://github.com/akagami-harsh) in [#5257](https://github.com/jaegertracing/jaeger/pull/5257))
* [jaeger-v2] add support for cassandra ([@Pushkarm029](https://github.com/Pushkarm029) in [#5253](https://github.com/jaegertracing/jaeger/pull/5253))

#### 👷 CI Improvements

* Allow go-leak linter to fail ci ([@yurishkuro](https://github.com/yurishkuro) in [#5316](https://github.com/jaegertracing/jaeger/pull/5316))
* [jaeger-v2] add grpc storage backend integration test ([@james-ryans](https://github.com/james-ryans) in [#5259](https://github.com/jaegertracing/jaeger/pull/5259))
* Github actions added to block prs from fork/main branch ([@varshith257](https://github.com/varshith257) in [#5272](https://github.com/jaegertracing/jaeger/pull/5272))


### 📊 UI Changes

* UI pinned to version [1.40.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1400-2024-04-02).


1.55.0 (2024-03-04)
-------------------
### Backend Changes

#### ✨ New Features:

* Support uploading traces to UI in OpenTelemetry format (OTLP/JSON) ([@NavinShrinivas](https://github.com/NavinShrinivas) in [#5155](https://github.com/jaegertracing/jaeger/pull/5155))
* Add Elasticsearch storage support for adaptive sampling ([@Pushkarm029](https://github.com/Pushkarm029) in [#5158](https://github.com/jaegertracing/jaeger/pull/5158))

#### 🐞 Bug fixes, Minor Improvements:

* Add the `print-config` subcommand ([@gmafrac](https://github.com/gmafrac) in [#5200](https://github.com/jaegertracing/jaeger/pull/5200))
* Return more detailed errors from ES storage ([@yurishkuro](https://github.com/yurishkuro) in [#5209](https://github.com/jaegertracing/jaeger/pull/5209))
* Bump go version ([@yurishkuro](https://github.com/yurishkuro) in [#5180](https://github.com/jaegertracing/jaeger/pull/5180))

#### 🚧 Experimental Features:

* [jaeger-v2] Add support for gRPC storarge ([@james-ryans](https://github.com/james-ryans) in [#5228](https://github.com/jaegertracing/jaeger/pull/5228))
* [jaeger-v2] Add support for Elasticsearch ([@akagami-harsh](https://github.com/akagami-harsh) in [#5152](https://github.com/jaegertracing/jaeger/pull/5152))

### 📊 UI Changes

* UI pinned to version [1.39.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1390-2024-03-04).


1.54.0 (2024-02-06)
-------------------

### Backend Changes

#### ⛔ Breaking Changes:

* Remove remnants of internal otlp types ([@yurishkuro](https://github.com/yurishkuro) in [#5107](https://github.com/jaegertracing/jaeger/pull/5107))
* Use official otlp types in api_v3 and avoid triple-marshaling ([@yurishkuro](https://github.com/yurishkuro) in [#5098](https://github.com/jaegertracing/jaeger/pull/5098))

#### ✨ New Features:

* [jaeger-v2] add support for badger ([@akagami-harsh](https://github.com/akagami-harsh) in [#5112](https://github.com/jaegertracing/jaeger/pull/5112))

#### 🐞 Bug fixes, Minor Improvements:

* [jaeger-v2] streamline storage initialization ([@yurishkuro](https://github.com/yurishkuro) in [#5171](https://github.com/jaegertracing/jaeger/pull/5171))
* Replace security self-assesment with one from cncf/tag-security ([@jkowall](https://github.com/jkowall) in [#5142](https://github.com/jaegertracing/jaeger/pull/5142))
* Avoid changing a correct order of span references ([@sherwoodwang](https://github.com/sherwoodwang) in [#5121](https://github.com/jaegertracing/jaeger/pull/5121))

#### 👷 CI Improvements:

* Remove test summary reports ([@albertteoh](https://github.com/albertteoh) in [#5126](https://github.com/jaegertracing/jaeger/pull/5126))

### UI Changes

* UI pinned to version [1.38.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1380-2024-02-06).

1.53.0 (2024-01-08)
-------------------

### Backend Changes

#### ⛔ Breaking Changes:

* 💤 swap zipkin server for zipkin receiver from otel collector contrib ([@yurishkuro](https://github.com/yurishkuro) in [#5045](https://github.com/jaegertracing/jaeger/pull/5045))
* Make all-in-one metric names match metrics from standalone components ([@yurishkuro](https://github.com/yurishkuro) in [#5008](https://github.com/jaegertracing/jaeger/pull/5008))

#### 🐞 Bug fixes, Minor Improvements:

* Upgrade thrift compiler to v0.19 and regenerate types ([@yurishkuro](https://github.com/yurishkuro) in [#5080](https://github.com/jaegertracing/jaeger/pull/5080))
* Add gogo/protobuf to opentelemetry otlp data model ([@yurishkuro](https://github.com/yurishkuro) in [#5067](https://github.com/jaegertracing/jaeger/pull/5067))
* Remove grpc-gateway dependency ([@yurishkuro](https://github.com/yurishkuro) in [#5060](https://github.com/jaegertracing/jaeger/pull/5060))
* Add manual implementation of apiv3 http endpoints ([@yurishkuro](https://github.com/yurishkuro) in [#5054](https://github.com/jaegertracing/jaeger/pull/5054))
* Allow specifying version for hotrod docker-compose ([@yurishkuro](https://github.com/yurishkuro) in [#5011](https://github.com/jaegertracing/jaeger/pull/5011))

#### 👷 CI Improvements:

* Publish go tip test report ([@albertteoh](https://github.com/albertteoh) in [#5082](https://github.com/jaegertracing/jaeger/pull/5082))
* Upload test report ([@albertteoh](https://github.com/albertteoh) in [#5035](https://github.com/jaegertracing/jaeger/pull/5035))
* Separate test report collection from the main target ([@yurishkuro](https://github.com/yurishkuro) in [#5061](https://github.com/jaegertracing/jaeger/pull/5061))
* Bugfix: set pipefail when running unit tests to preserve exit code ([@yurishkuro](https://github.com/yurishkuro) in [#5057](https://github.com/jaegertracing/jaeger/pull/5057))
* Regenerate thrift types and enable thrift check ([@yurishkuro](https://github.com/yurishkuro) in [#5039](https://github.com/jaegertracing/jaeger/pull/5039))
* Regenerate hotrod proto ([@yurishkuro](https://github.com/yurishkuro) in [#5040](https://github.com/jaegertracing/jaeger/pull/5040))
* Fix permission failed on checks-run ([@albertteoh](https://github.com/albertteoh) in [#5041](https://github.com/jaegertracing/jaeger/pull/5041))
* Refactor protobuf types generation ([@yurishkuro](https://github.com/yurishkuro) in [#5037](https://github.com/jaegertracing/jaeger/pull/5037))
* Publish test report ([@albertteoh](https://github.com/albertteoh) in [#5030](https://github.com/jaegertracing/jaeger/pull/5030))
* Ci: simplify check-label workflow ([@EshaanAgg](https://github.com/EshaanAgg) in [#5033](https://github.com/jaegertracing/jaeger/pull/5033))
* Fix goroutine leaks in several packages ([@yurishkuro](https://github.com/yurishkuro) in [#5026](https://github.com/jaegertracing/jaeger/pull/5026))
* Add goleak check in more tests that do not fail ([@akagami-harsh](https://github.com/akagami-harsh) in [#5025](https://github.com/jaegertracing/jaeger/pull/5025))
* Ci: add retry logic in the install go tip github action ([@akagami-harsh](https://github.com/akagami-harsh) in [#5022](https://github.com/jaegertracing/jaeger/pull/5022))
* Move go tip installation into sub-action ([@yurishkuro](https://github.com/yurishkuro) in [#5020](https://github.com/jaegertracing/jaeger/pull/5020))
* Add goleak check to packages with empty tests ([@yurishkuro](https://github.com/yurishkuro) in [#5017](https://github.com/jaegertracing/jaeger/pull/5017))
* Add goleak check to cmd/agent/app/configmanager ([@yurishkuro](https://github.com/yurishkuro) in [#5015](https://github.com/jaegertracing/jaeger/pull/5015))
* Feature: add goleak to check goroutine leak in tests ([@akagami-harsh](https://github.com/akagami-harsh) in [#5010](https://github.com/jaegertracing/jaeger/pull/5010))
* Remove custom gocache location ([@yurishkuro](https://github.com/yurishkuro) in [#4995](https://github.com/jaegertracing/jaeger/pull/4995))

1.52.0 (2023-12-05)
-------------------

### Backend Changes

#### ✨ New Features:

* Support Elasticsearch 8.x ([@pmuls99](https://github.com/pmuls99) in [#4829](https://github.com/jaegertracing/jaeger/pull/4829))
* Make ArchiveTrace button auto-configurable ([@thecoons](https://github.com/thecoons) in [#4913](https://github.com/jaegertracing/jaeger/pull/4913))

#### 🐞 Bug fixes, Minor Improvements:

* [SPM] differentiate null from no error data ([@albertteoh](https://github.com/albertteoh) in [#4985](https://github.com/jaegertracing/jaeger/pull/4985))
* Fix example/grafana-integration ([@angristan](https://github.com/angristan) in [#4980](https://github.com/jaegertracing/jaeger/pull/4980))
* Fix (badger): add missing SamplingStoreFactory.CreateLock method ([@slayer321](https://github.com/slayer321) in [#4966](https://github.com/jaegertracing/jaeger/pull/4966))
* Normalize metric names due to breaking change ([@albertteoh](https://github.com/albertteoh) in [#4957](https://github.com/jaegertracing/jaeger/pull/4957))
* [kafka-consumer] add topic name as a tag to offset manager metrics ([@abliqo](https://github.com/abliqo) in [#4951](https://github.com/jaegertracing/jaeger/pull/4951))
* Make UI placeholder more descriptive ([@yurishkuro](https://github.com/yurishkuro) in [#4937](https://github.com/jaegertracing/jaeger/pull/4937))
* Remove google.golang.org/protobuf dependency from model & storage apis ([@akagami-harsh](https://github.com/akagami-harsh) in [#4917](https://github.com/jaegertracing/jaeger/pull/4917))
* Read OTEL env vars for resource attributes ([@yurishkuro](https://github.com/yurishkuro) in [#4932](https://github.com/jaegertracing/jaeger/pull/4932))

#### 🚧 Experimental Features:

* Exp: rename jaeger-v2 binary to just jaeger ([@yurishkuro](https://github.com/yurishkuro) in [#4918](https://github.com/jaegertracing/jaeger/pull/4918))

#### 👷 CI Improvements:

* [CI]: improve kafka integration test self-sufficiency ([@RipulHandoo](https://github.com/RipulHandoo) in [#4989](https://github.com/jaegertracing/jaeger/pull/4989))
* Separate all-in-one integration tests for v1 and v2 ([@yurishkuro](https://github.com/yurishkuro) in [#4968](https://github.com/jaegertracing/jaeger/pull/4968))
* Collect code coverage from integration tests and upload to codecov ([@yurishkuro](https://github.com/yurishkuro) in [#4964](https://github.com/jaegertracing/jaeger/pull/4964))
* [CI/ES] use default template priorities ([@yurishkuro](https://github.com/yurishkuro) in [#4962](https://github.com/jaegertracing/jaeger/pull/4962))
* Unleash dependabot on docker files and add dependency review workflow ([@step-security-bot](https://github.com/step-security-bot) in [#4945](https://github.com/jaegertracing/jaeger/pull/4945))
* Split unit-test workflow into tests and lint ([@MeenuyD](https://github.com/MeenuyD) in [#4933](https://github.com/jaegertracing/jaeger/pull/4933))
* [CI]: harden github actions ([@step-security-bot](https://github.com/step-security-bot) in [#4923](https://github.com/jaegertracing/jaeger/pull/4923))
* [CI]: build jaeger v2 image on main branch runs ([@yurishkuro](https://github.com/yurishkuro) in [#4920](https://github.com/jaegertracing/jaeger/pull/4920))
* Exp: publish jaeger v2 image ([@yurishkuro](https://github.com/yurishkuro) in [#4919](https://github.com/jaegertracing/jaeger/pull/4919))
* [CI]: set default to fix 'unbound variable' error on main ([@yurishkuro](https://github.com/yurishkuro) in [#4916](https://github.com/jaegertracing/jaeger/pull/4916))
* [CI]: test jaeger-v2 as all-in-one in ci ([@yurishkuro](https://github.com/yurishkuro) in [#4890](https://github.com/jaegertracing/jaeger/pull/4890))
* Fix release script broken by recent linting cleanup ([@yurishkuro](https://github.com/yurishkuro) in [#4915](https://github.com/jaegertracing/jaeger/pull/4915))

### UI Changes

* UI pinned to version [1.36.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1360-2023-12-05).

1.51.0 (2023-11-04)
-------------------

### Backend Changes

#### ✨ New Features:

* Feat: add sampling store support to badger ([@slayer321](https://github.com/slayer321) in [#4834](https://github.com/jaegertracing/jaeger/pull/4834))
* Feat: add span adjuster that moves some otel resource attributes to span.process ([@james-ryans](https://github.com/james-ryans) in [#4844](https://github.com/jaegertracing/jaeger/pull/4844))
* Add product/file version in windows executables ([@ResamVi](https://github.com/ResamVi) in [#4811](https://github.com/jaegertracing/jaeger/pull/4811))

#### 🐞 Bug fixes, Minor Improvements:

* Fix dependency policy and add to security-insights.yml ([@jkowall](https://github.com/jkowall) in [#4907](https://github.com/jaegertracing/jaeger/pull/4907))
* Add reload interval to otel server certificates ([@james-ryans](https://github.com/james-ryans) in [#4898](https://github.com/jaegertracing/jaeger/pull/4898))
* Feat: add blackhole storage, for benchmarking ([@yurishkuro](https://github.com/yurishkuro) in [#4896](https://github.com/jaegertracing/jaeger/pull/4896))
* Add otel resource detector to jaeger components ([@james-ryans](https://github.com/james-ryans) in [#4864](https://github.com/jaegertracing/jaeger/pull/4864))
* Fix batchprocessor to set correct span format flags ([@k0zl](https://github.com/k0zl) in [#4796](https://github.com/jaegertracing/jaeger/pull/4796))
* Expose collector ports in docker images ([@arunvelsriram](https://github.com/arunvelsriram) in [#4810](https://github.com/jaegertracing/jaeger/pull/4810))

#### 🚧 Experimental Features:

* Exp(jaeger-v2): simplify all-in-one configuration ([@yurishkuro](https://github.com/yurishkuro) in [#4875](https://github.com/jaegertracing/jaeger/pull/4875))
* Exp: support primary and archive storage ([@yurishkuro](https://github.com/yurishkuro) in [#4873](https://github.com/jaegertracing/jaeger/pull/4873))
* Feat(jaeger-v2): create default config for all-in-one ([@yurishkuro](https://github.com/yurishkuro) in [#4842](https://github.com/jaegertracing/jaeger/pull/4842))

#### 👷 CI Improvements:

* Ci: split the install-tools into test/build groups ([@MeenuyD](https://github.com/MeenuyD) in [#4878](https://github.com/jaegertracing/jaeger/pull/4878))
* Simplify binary building in makefile ([@yurishkuro](https://github.com/yurishkuro) in [#4885](https://github.com/jaegertracing/jaeger/pull/4885))
* Ci: pass variable instead of calling make build-xxx-debug ([@yurishkuro](https://github.com/yurishkuro) in [#4883](https://github.com/jaegertracing/jaeger/pull/4883))
* Simplify makefile ([@yurishkuro](https://github.com/yurishkuro) in [#4882](https://github.com/jaegertracing/jaeger/pull/4882))
* Test: add more linters ([@yurishkuro](https://github.com/yurishkuro) in [#4881](https://github.com/jaegertracing/jaeger/pull/4881))
* Ci: enable linting of code in examples/ ([@yurishkuro](https://github.com/yurishkuro) in [#4880](https://github.com/jaegertracing/jaeger/pull/4880))
* Ci: keep the ui asset's .gz file timestamps the same as the original file ([@yurishkuro](https://github.com/yurishkuro) in [#4879](https://github.com/jaegertracing/jaeger/pull/4879))
* Add first pass at the security-insights.yml ([@jkowall](https://github.com/jkowall) in [#4872](https://github.com/jaegertracing/jaeger/pull/4872))
* Create scorecard.yml for ossf implementation ([@jkowall](https://github.com/jkowall) in [#4870](https://github.com/jaegertracing/jaeger/pull/4870))
* Add ci validation of shell scripts using shellcheck ([@akagami-harsh](https://github.com/akagami-harsh) in [#4826](https://github.com/jaegertracing/jaeger/pull/4826))
* Chore: add dynamic loading bar functionality to release-notes.py ([@anshgoyalevil](https://github.com/anshgoyalevil) in [#4857](https://github.com/jaegertracing/jaeger/pull/4857))
* Ci: add the label-check workflow to verify changelog labels on each pr ([@anshgoyalevil](https://github.com/anshgoyalevil) in [#4847](https://github.com/jaegertracing/jaeger/pull/4847))
* Ci(hotrod): print hotrod container logs in case of test failure ([@yurishkuro](https://github.com/yurishkuro) in [#4845](https://github.com/jaegertracing/jaeger/pull/4845))
* Ci: drop -v from ci unit tests to make failures easier to see ([@yurishkuro](https://github.com/yurishkuro) in [#4839](https://github.com/jaegertracing/jaeger/pull/4839))
* Use commit hash as image label when building & integration-testing ([@yurishkuro](https://github.com/yurishkuro) in [#4824](https://github.com/jaegertracing/jaeger/pull/4824))
* Clean-up some linter warnings in build scripts ([@yurishkuro](https://github.com/yurishkuro) in [#4823](https://github.com/jaegertracing/jaeger/pull/4823))
* Fix build-all-in-one-image script ([@albertteoh](https://github.com/albertteoh) in [#4819](https://github.com/jaegertracing/jaeger/pull/4819))
* [ci-release] improve release workflow for manual runs ([@yurishkuro](https://github.com/yurishkuro) in [#4818](https://github.com/jaegertracing/jaeger/pull/4818))
* Add --force to docker commands ([@albertteoh](https://github.com/albertteoh) in [#4820](https://github.com/jaegertracing/jaeger/pull/4820))
* Use setup-node.js for publish release ([@albertteoh](https://github.com/albertteoh) in [#4816](https://github.com/jaegertracing/jaeger/pull/4816))
* Clean up ci scripts and prune docker images between builds ([@yurishkuro](https://github.com/yurishkuro) in [#4815](https://github.com/jaegertracing/jaeger/pull/4815))
* Clean-up & fortify ci-release ([@yurishkuro](https://github.com/yurishkuro) in [#4813](https://github.com/jaegertracing/jaeger/pull/4813))

### UI Changes

* UI pinned to version [1.35.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1350-2023-11-02).

1.50.0 (2023-10-06)
-------------------

### Backend Changes

#### ⛔ Breaking Changes

* [sampling] Remove support for SAMPLING_TYPE env var and 'static' value ([@yurishkuro](https://github.com/yurishkuro) in [#4735](https://github.com/jaegertracing/jaeger/pull/4735))
* Use non-root user in built containers ([@nikzayn](https://github.com/nikzayn) in [#4783](https://github.com/jaegertracing/jaeger/pull/4783)) - this change may cause issues with existing installations using Badger storage, because the existing files would be owned by a different user and would not be writeable after Jaeger upgrade. The workaround is to manually chown the files to the new user (uid=10001).

#### New Features

* Add cassandra schema compaction window configuration ([@sameersecond](https://github.com/sameersecond) in [#4767](https://github.com/jaegertracing/jaeger/pull/4767))
* Add jaeger-v2 single binary based on otel collector ([@yurishkuro](https://github.com/yurishkuro) in [#4766](https://github.com/jaegertracing/jaeger/pull/4766))
* [kafka-consumer] Consumer metrics should have a tag with topic name ([@abliqo](https://github.com/abliqo) in [#4778](https://github.com/jaegertracing/jaeger/pull/4778))
* Support http proxy env variables ([@pavolloffay](https://github.com/pavolloffay) in [#4769](https://github.com/jaegertracing/jaeger/pull/4769))
* Support reloading es client's password from file ([@haanhvu](https://github.com/haanhvu) in [#4342](https://github.com/jaegertracing/jaeger/pull/4342))

#### Bug fixes, Minor Improvements

* Fix jaegerqueryreqsfailing alert rule missing 'operation' in query ([@james-ryans](https://github.com/james-ryans) in [#4797](https://github.com/jaegertracing/jaeger/pull/4797))
* Add e2e test for sampling storage ([@slayer321](https://github.com/slayer321) in [#4772](https://github.com/jaegertracing/jaeger/pull/4772))
* [tests] Simplify cassandra e2e test cleanup ([@yurishkuro](https://github.com/yurishkuro) in [#4794](https://github.com/jaegertracing/jaeger/pull/4794))
* [tests] Fix failing e2e test for cassandra storage ([@slayer321](https://github.com/slayer321) in [#4776](https://github.com/jaegertracing/jaeger/pull/4776))
* Remove unneeded references to opentracing ([@yurishkuro](https://github.com/yurishkuro) in [#4790](https://github.com/jaegertracing/jaeger/pull/4790))
* Use non-root user in built containers ([@nikzayn](https://github.com/nikzayn) in [#4783](https://github.com/jaegertracing/jaeger/pull/4783))
* Run all integration tests against cassandra ([@yurishkuro](https://github.com/yurishkuro) in [#4773](https://github.com/jaegertracing/jaeger/pull/4773))
* [hotrod] Log driver locations as json to demo respective ui capability ([@yurishkuro](https://github.com/yurishkuro) in [#4765](https://github.com/jaegertracing/jaeger/pull/4765))
* Replace python script with tracegen ([@albertteoh](https://github.com/albertteoh) in [#4753](https://github.com/jaegertracing/jaeger/pull/4753))
* [fix] Close elasticsearch client properly ([@Lauquik](https://github.com/Lauquik) in [#4754](https://github.com/jaegertracing/jaeger/pull/4754))
* Add deprecation warning to jaeger-agent ([@yurishkuro](https://github.com/yurishkuro) in [#4749](https://github.com/jaegertracing/jaeger/pull/4749))
* Deprecate grpc-storage-plugin sidecar model ([@yurishkuro](https://github.com/yurishkuro) in [#4744](https://github.com/jaegertracing/jaeger/pull/4744))
* Upgrade query api v3 to official opentelemetry format ([@yurishkuro](https://github.com/yurishkuro) in [#4742](https://github.com/jaegertracing/jaeger/pull/4742))
* [SPM] Deprecate support for spanmetrics processor naming convention ([@yurishkuro](https://github.com/yurishkuro) in [#4741](https://github.com/jaegertracing/jaeger/pull/4741))
* Deprecate expvar metrics backend ([@yurishkuro](https://github.com/yurishkuro) in [#4740](https://github.com/jaegertracing/jaeger/pull/4740))
* Fix flaky testgetroundtripper* tests ([@albertteoh](https://github.com/albertteoh) in [#4738](https://github.com/jaegertracing/jaeger/pull/4738))

### UI Changes

* UI pinned to version [1.34.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1340-2023-10-04).

1.49.0 (2023-09-07)
-------------------

### Backend Changes

#### ⛔ Breaking Changes

* [hotrod] Make OTLP the default exporter in HotROD ([@yurishkuro](https://github.com/yurishkuro) in [#4698](https://github.com/jaegertracing/jaeger/pull/4698))
* [SPM] Support spanmetrics connector by default ([@albertteoh](https://github.com/albertteoh) in [#4704](https://github.com/jaegertracing/jaeger/pull/4704))
* [tracegen] Stop supporting -trace-exporter=jaeger ([@yurishkuro](https://github.com/yurishkuro) in [#4717](https://github.com/jaegertracing/jaeger/pull/4717))
* [hotrod] Stop supporting -otel-exporter=jaeger ([@yurishkuro](https://github.com/yurishkuro) in [#4719](https://github.com/jaegertracing/jaeger/pull/4719))
* [hotrod] Metrics endpoints moved from route service (:8083) to frontend service (:8080) ([@yurishkuro](https://github.com/yurishkuro) in [#4720](https://github.com/jaegertracing/jaeger/pull/4720))

#### Bug fixes, Minor Improvements

* Allow disabling brearer token override from request in metrics store ([@pavolloffay](https://github.com/pavolloffay) in [#4726](https://github.com/jaegertracing/jaeger/pull/4726))
* Add the enable tracing opt-in flag ([@albertteoh](https://github.com/albertteoh) in [#4685](https://github.com/jaegertracing/jaeger/pull/4685))
* [tracegen] Add build info during compilation ([@yurishkuro](https://github.com/yurishkuro) in [#4727](https://github.com/jaegertracing/jaeger/pull/4727))
* Log version/build info on startup ([@yurishkuro](https://github.com/yurishkuro) in [#4723](https://github.com/jaegertracing/jaeger/pull/4723))
* [zipkin] Replace zipkin exporter from jaeger sdk with otel zipkin exp ([@afzal442](https://github.com/afzal442) in [#4674](https://github.com/jaegertracing/jaeger/pull/4674))

### UI Changes

* UI pinned to version [1.33.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1330-2023-08-06).

1.48.0 (2023-08-15)
-------------------

### Backend Changes

#### Bug fixes, Minor Improvements

* [fix] Disable tracing of OTLP Receiver ([@yurishkuro](https://github.com/yurishkuro) in [#4662](https://github.com/jaegertracing/jaeger/pull/4662))
* [hotrod/observer_test] Switch to OpenTelemetry ([@afzal442](https://github.com/afzal442) in [#4635](https://github.com/jaegertracing/jaeger/pull/4635))
* [memstore-plugin]Switch to OpenTelemetry SDK ([@afzal442](https://github.com/afzal442) in [#4643](https://github.com/jaegertracing/jaeger/pull/4643))
* [tracegen] Allow to control cardinality of attribute keys ([@yurishkuro](https://github.com/yurishkuro) in [#4634](https://github.com/jaegertracing/jaeger/pull/4634))
* Replace OT const wth OTEL trace.span for zipkin comp ([@afzal442](https://github.com/afzal442) in [#4625](https://github.com/jaegertracing/jaeger/pull/4625))
* Replace OpenTracing instrumentation with OpenTelemetry in grpc storage plugin ([@afzal442](https://github.com/afzal442) in [#4611](https://github.com/jaegertracing/jaeger/pull/4611))
* Replace OT trace with `otel trace` spans type to span model ([@afzal442](https://github.com/afzal442) in [#4614](https://github.com/jaegertracing/jaeger/pull/4614))
* Replace cassandra-spanstore tracing instrumentation with`OTEL` ([@afzal442](https://github.com/afzal442) in [#4599](https://github.com/jaegertracing/jaeger/pull/4599))
* Replace es-spanstore tracing instrumentation with OpenTelemetry ([@afzal442](https://github.com/afzal442) in [#4596](https://github.com/jaegertracing/jaeger/pull/4596))
* Replace metricsstore/reader tracing instrumentation with OpenTelemetry ([@afzal442](https://github.com/afzal442) in [#4595](https://github.com/jaegertracing/jaeger/pull/4595))
* Replace Jaeger SDK with OTEL SDK + OT Bridge ([@afzal442](https://github.com/afzal442) in [#4574](https://github.com/jaegertracing/jaeger/pull/4574))
* [kafka-consumer] Ingester should use topic name from actual Kafka consumer instead of configuration ([@abliqo](https://github.com/abliqo) in [#4593](https://github.com/jaegertracing/jaeger/pull/4593))
* Enable CORS settings on OTLP HTTP endpoint ([@pmuls99](https://github.com/pmuls99) in [#4586](https://github.com/jaegertracing/jaeger/pull/4586))
* [hotrod] Return trace ID via traceresponse header ([@yurishkuro](https://github.com/yurishkuro) in [#4584](https://github.com/jaegertracing/jaeger/pull/4584))
* [hotrod] Remove most references to OpenTracing ([@yurishkuro](https://github.com/yurishkuro) in [#4585](https://github.com/jaegertracing/jaeger/pull/4585))
* [hotrod] Validate user input to avoid security warnings from code scanning ([@yurishkuro](https://github.com/yurishkuro) in [#4583](https://github.com/jaegertracing/jaeger/pull/4583))
* [hotrod] Upgrade HotROD to use OpenTelemetry instrumentation ([@afzal442](https://github.com/afzal442) in [#4548](https://github.com/jaegertracing/jaeger/pull/4548))
* [kafka-consumer] Use wait group to ensure goroutine is finished before returning from Close ([@kennyaz](https://github.com/kennyaz) in [#4582](https://github.com/jaegertracing/jaeger/pull/4582))
* [tracegen] Enable BlockOnQueueFull in OTel SDK to avoid dropped spans ([@haanhvu](https://github.com/haanhvu) in [#4578](https://github.com/jaegertracing/jaeger/pull/4578))
* [hotrod] Handle both OT and OTEL baggage ([@yurishkuro](https://github.com/yurishkuro) in [#4572](https://github.com/jaegertracing/jaeger/pull/4572))

### UI Changes

* UI pinned to version [1.32.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1320-2023-08-14).

1.47.0 (2023-07-05)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

* [SPM] Due to a breaking change in OpenTelemetry's prometheus exporter ([details](https://github.com/open-telemetry/opentelemetry-collector-contrib/releases/tag/v0.80.0))
  metric names will no longer be normalized by default, meaning that the expected metric names would be `calls` and
  `duration_*`. Backwards compatibility with older OpenTelemetry Collector versions can be achieved through the following flags:
  * `prometheus.query.normalize-calls`: If true, normalizes the "calls" metric name. e.g. "calls_total".
  * `prometheus.query.normalize-duration`: If true, normalizes the "duration" metric name to include the duration units. e.g. "duration_milliseconds_bucket".

#### New Features
* [Cassandra] Add Configuration.Close() to ensure TLS cert watcher is closed ([@kennyaz](https://github.com/kennyaz) in [#4515](https://github.com/jaegertracing/jaeger/pull/4515))
* Add *.kerberos.disable-fast-negotiation option to Kafka consumer ([@pmuls99](https://github.com/pmuls99) in [#4520](https://github.com/jaegertracing/jaeger/pull/4520))
* Support Prometheus normalization for specific metrics related to OpenTelemetry compatibility ([@albertteoh](https://github.com/albertteoh) in [#4555](https://github.com/jaegertracing/jaeger/pull/4555))

#### Bug fixes, Minor Improvements

* Add readme for memstore plugin ([@yurishkuro](https://github.com/yurishkuro) in [283bdd9](https://github.com/jaegertracing/jaeger/commit/283bdd93cbb4a467842625d8eb320722fcb83494))
* Pass a wrapper instead of `opentracing.Tracer` to ease migration to OTEL in the future [part 1] ([@afzalbin64](https://github.com/afzalbin64) in [#4529](https://github.com/jaegertracing/jaeger/pull/4529))
* [hotROD] Add OTEL instrumentation to customer svc ([@afzal442](https://github.com/afzal442) in [#4559](https://github.com/jaegertracing/jaeger/pull/4559))
* [hotROD] Replace gRPC instrumentation with OTEL ([@afzal442](https://github.com/afzal442) in [#4558](https://github.com/jaegertracing/jaeger/pull/4558))
* [hotROD]: Upgrade `redis` service to use native OTEL instrumentation ([@afzal442](https://github.com/afzal442) in [#4533](https://github.com/jaegertracing/jaeger/pull/4533))
* [hotROD] Fix OTEL logging in HotRod example ([@albertteoh](https://github.com/albertteoh) in [#4556](https://github.com/jaegertracing/jaeger/pull/4556))
* [hotrod] Reduce span exporter's batch timeout to let the spans be exported sooner ([@GLVSKiriti](https://github.com/GLVSKiriti) in [#4518](https://github.com/jaegertracing/jaeger/pull/4518))
* [tracegen] Add options to generate more spans and attributes for additional use cases ([@yurishkuro](https://github.com/yurishkuro) in [#4535](https://github.com/jaegertracing/jaeger/pull/4535))
* Build improvement to rebuild jaeger-ui if the tree does not match any tag ([@bobrik](https://github.com/bobrik) in [#4553](https://github.com/jaegertracing/jaeger/pull/4553))
* [Test] Fixed a race condition causing unit test failure by changing the logging  ([@yurishkuro](https://github.com/yurishkuro) in [#4546](https://github.com/jaegertracing/jaeger/pull/4546)) resolves [#4497](https://github.com/jaegertracing/jaeger/issues/4497)


### UI Changes

* UI pinned to version [1.31.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1310-2023-07-05).


1.46.0 (2023-06-05)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

OTLP receiver is now enabled by default ([#4494](https://github.com/jaegertracing/jaeger/pull/4494)). This change follows the Jaeger's strategic direction of aligning closely with the OpenTelemetry project. This may cause port conflicts if  `jaeger-collector` is depoyed in host network namespace. The original `--collector.otlp.enabled` option is still available and MUST be set to `false` if OTLP receiver is not desired.

#### New Features

* Make OTLP receiver enabled by default ([@yurishkuro](https://github.com/yurishkuro) in [#4494](https://github.com/jaegertracing/jaeger/pull/4494))
* [SPM] Add support for OpenTelemetry SpanMetrics Connector ([@albertteoh](https://github.com/albertteoh) in [#4452](https://github.com/jaegertracing/jaeger/pull/4452)). See [Migration README](https://github.com/jaegertracing/jaeger/blob/main/docker-compose/monitor/README.md#migrating).

#### Bug fixes, Minor Improvements

* Log processor error in Kafka consumer ([@pavolloffay](https://github.com/pavolloffay) in [#4399](https://github.com/jaegertracing/jaeger/pull/4399))
* [bug] Remove TerminateAfter from Elasticsearch/Opensearch query resulting in incomplete span count/list ([@Jakob3xD](https://github.com/Jakob3xD) in [#4336](https://github.com/jaegertracing/jaeger/pull/4336))
* [agent] Use RawConn.Control to get fd instead of Fd() to prevent deadlock on shutdown ([@ChenX1993](https://github.com/ChenX1993) in [#4449](https://github.com/jaegertracing/jaeger/pull/4449))
* [SPM] Fix docker compose command ([@tqi-raurora](https://github.com/tqi-raurora) in [#4444](https://github.com/jaegertracing/jaeger/pull/4444))

#### Maintenance

* [test] Fix flaky test - TestSpanProcessorWithOnDroppedSpanOption ([@yurishkuro](https://github.com/yurishkuro) in [#4489](https://github.com/jaegertracing/jaeger/pull/4489))
* [ci] Skip debug builds when not making a release ([@psk001](https://github.com/psk001) in [#4496](https://github.com/jaegertracing/jaeger/pull/4496))
* Fix some function comments ([@cuishuang](https://github.com/cuishuang) in [#4410](https://github.com/jaegertracing/jaeger/pull/4410))
* Increase dependabot open-pull-requests-limit=10 ([@yurishkuro](https://github.com/yurishkuro) in [04548fc](https://github.com/jaegertracing/jaeger/commit/04548fc339689f970da2de36b964fd3abfca41c2))
* Add jkowall as release manger for July ([@jkowall](https://github.com/jkowall) in [#4446](https://github.com/jaegertracing/jaeger/pull/4446))
* Fix versions in release schedule ([@yurishkuro](https://github.com/yurishkuro) in [8a9d13a](https://github.com/jaegertracing/jaeger/commit/8a9d13a31b477707cec73a9b7bf6242b27cec0ea))

### UI Changes

* UI pinned to version [1.30.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1300-2023-06-05).


1.45.0 (2023-05-03)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* Add HTTP post port mapping to docker command ([@albertteoh](https://github.com/albertteoh) in [#4407](https://github.com/jaegertracing/jaeger/pull/4407))
* Simplify ES config and factory ([@yurishkuro](https://github.com/yurishkuro) in [#4396](https://github.com/jaegertracing/jaeger/pull/4396))
* Add otlp-grpc for tracegen's trace-exporter ([@boysusu](https://github.com/boysusu) in [#4374](https://github.com/jaegertracing/jaeger/pull/4374))
* Allow follows-from reference as a parent span id ([@kubarydz](https://github.com/kubarydz) in [#4382](https://github.com/jaegertracing/jaeger/pull/4382))
* Expose drop span hook as an option in Collector SpanProcessor ([@ChenX1993](https://github.com/ChenX1993) in [#4387](https://github.com/jaegertracing/jaeger/pull/4387))

### UI Changes

* UI pinned to version [1.29.1](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1291-2023-05-03).


1.44.0 (2023-04-10)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* Store span warnings as tags in Cassandra ([@utsavoza](https://github.com/utsavoza) in [#4313](https://github.com/jaegertracing/jaeger/pull/4313))
* Add Keep-Alive flag for Zipkin HTTP server ([@topjung3](https://github.com/topjung3) in [#4366](https://github.com/jaegertracing/jaeger/pull/4366))
* Log access to static assets; remove favicon test ([@yurishkuro](https://github.com/yurishkuro) in [#4302](https://github.com/jaegertracing/jaeger/pull/4302))

### UI Changes

* UI pinned to version [1.29.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1290-2023-04-10).

1.43.0 (2023-03-15)
-------------------
### Backend Changes

#### New Features

* Support secure OTLP exporter config for hotrod ([@graphaelli](https://github.com/graphaelli) in [#4231](https://github.com/jaegertracing/jaeger/pull/4231))
* Jaeger YugabyteDB(YCQL) support ([@HarshDaryani896](https://github.com/HarshDaryani896) in [#4220](https://github.com/jaegertracing/jaeger/pull/4220))

#### Bug fixes, Minor Improvements

* Replace pkg/multierror with standard errors.Join ([@ClementRepo](https://github.com/ClementRepo) in [#4293](https://github.com/jaegertracing/jaeger/pull/4293))
* Remove pkg/multicloser ([@yurishkuro](https://github.com/yurishkuro) in [#4291](https://github.com/jaegertracing/jaeger/pull/4291))
* Refactor build linux artifacts only for PR ([@Eileen-Yu](https://github.com/Eileen-Yu) in [#4286](https://github.com/jaegertracing/jaeger/pull/4286))
* Speed-up CI by using published UI artifacts ([@shubbham1219](https://github.com/shubbham1219) in [#4251](https://github.com/jaegertracing/jaeger/pull/4251))
* Update Go version to 1.20 ([@SaarthakMaini](https://github.com/SaarthakMaini) in [#4206](https://github.com/jaegertracing/jaeger/pull/4206))
* Use http.MethodGet instead of "GET" ([@my-git9](https://github.com/my-git9) in [#4248](https://github.com/jaegertracing/jaeger/pull/4248))
* Updating all-in-one path ([@bigfleet](https://github.com/bigfleet) in [#4234](https://github.com/jaegertracing/jaeger/pull/4234))
* Migrate the use of fsnotify to fswatcher in cert_watcher.go ([@haanhvu](https://github.com/haanhvu) in [#4232](https://github.com/jaegertracing/jaeger/pull/4232))
* Restore baggage support in HotROD 🚗 ([@yurishkuro](https://github.com/yurishkuro) in [#4225](https://github.com/jaegertracing/jaeger/pull/4225))

### UI Changes

* UI pinned to version [1.28.1](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1281-2023-03-15).

1.42.0 (2023-02-05)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

* HotROD 🚗 application is switched from Jaeger SDK to OpenTelemetry SDK ([@yurishkuro](https://github.com/yurishkuro) in [#4187](https://github.com/jaegertracing/jaeger/pull/4187)). Some environment variables previously accepted are no longer recognized. See PR for details.
  * Map old env vars from Jaeger SDK to OTel SDK vars ([@yurishkuro](https://github.com/yurishkuro) in [#4200](https://github.com/jaegertracing/jaeger/pull/4200))
  * Use patched version of github.com/opentracing-contrib/go-grpc in HotROD ([@yurishkuro](https://github.com/yurishkuro) in [#4210](https://github.com/jaegertracing/jaeger/pull/4210))
* `tracegen` utility is switched from Jaeger SDK to OpenTelemetry SDK ([@yurishkuro](https://github.com/yurishkuro) in [#4189](https://github.com/jaegertracing/jaeger/pull/4189)). Some environment variables previously accepted are no longer recognized. See PR for details.

#### New Features

* Add CLI flags for controlling HTTP server timeouts ([@yurishkuro](https://github.com/yurishkuro) in [#4167](https://github.com/jaegertracing/jaeger/pull/4167))
* Watch directories for certificate hot-reload ([@tsaarni](https://github.com/tsaarni) in [#4159](https://github.com/jaegertracing/jaeger/pull/4159))
* Support tenant header propagation in query service and grpc-plugin ([@pavolloffay](https://github.com/pavolloffay) in [#4151](https://github.com/jaegertracing/jaeger/pull/4151))


#### Bug fixes, Minor Improvements

* Replace Thrift-gen with Proto-gen types for sampling strategies ([@yurishkuro](https://github.com/yurishkuro) in [#4181](https://github.com/jaegertracing/jaeger/pull/4181))
* Use Protobuf-based JSON output for sampling strategies ([@yurishkuro](https://github.com/yurishkuro) in [#4173](https://github.com/jaegertracing/jaeger/pull/4173))
* [tests]: Use `t.Setenv` to set env vars in tests ([@Juneezee](https://github.com/Juneezee) in [#4169](https://github.com/jaegertracing/jaeger/pull/4169))
* ci(lint): bump golangci-lint to v1.50.1 ([@mmorel-35](https://github.com/mmorel-35) in [#4195](https://github.com/jaegertracing/jaeger/pull/4195))
* Convert token propagation integration test to plain unit test ([@yurishkuro](https://github.com/yurishkuro) in [#4162](https://github.com/jaegertracing/jaeger/pull/4162))
* Refactor scripts/es-integration-test.sh ([@yurishkuro](https://github.com/yurishkuro) in [#4161](https://github.com/jaegertracing/jaeger/pull/4161))
* Fix "index out of range" when receiving a dual client/server Zipkin span ([@yurishkuro](https://github.com/yurishkuro) in [#4160](https://github.com/jaegertracing/jaeger/pull/4160))

### UI Changes

* No changes.

1.41.0 (2023-01-04)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* Remove global platform arg in cassandra schema dockerfile ([@jkandasa](https://github.com/jkandasa) in [#4123](https://github.com/jaegertracing/jaeger/pull/4123))
* Add multi arch support to cassandra-schema container ([@jkandasa](https://github.com/jkandasa) in [#4122](https://github.com/jaegertracing/jaeger/pull/4122))

### UI

* No changes.

1.40.0 (2022-12-07)
-------------------
### Backend Changes

#### New Features

* Release signing ([@jkowall](https://github.com/jkowall) in [#4033](https://github.com/jaegertracing/jaeger/pull/4033))

#### Bug fixes, Minor Improvements

* Fix cassandra schema scripts to be able to run from a remote node ([@yurishkuro](https://github.com/yurishkuro) in [#4087](https://github.com/jaegertracing/jaeger/pull/4087))
* Catch panic from Prometheus client on invalid label strings ([@yurishkuro](https://github.com/yurishkuro) in [#4051](https://github.com/jaegertracing/jaeger/pull/4051))
* Span tags of type int64 may lose precision ([@shubbham1215](https://github.com/shubbham1215) in [#4034](https://github.com/jaegertracing/jaeger/pull/4034))

### UI

* UI pinned to version [1.27.3](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1273-2022-12-07).


1.39.0 (2022-11-01)
-------------------
### Backend Changes

#### New Features

* Add support for OpenSearch 2.x ([@gaurav-05](https://github.com/gaurav-05) in [#3966](https://github.com/jaegertracing/jaeger/pull/3966))

#### Bug fixes, Minor Improvements

* Pin SBOM action to commit ([@yurishkuro](https://github.com/yurishkuro) in [bb49249](https://github.com/jaegertracing/jaeger/commit/bb492490594c9d9321ed9242862ac2a8864ff771))
* Remove auth requirement on monitor demo ([@joe-elliott](https://github.com/joe-elliott) in [#4005](https://github.com/jaegertracing/jaeger/pull/4005))
* Increase sleep time when waiting for ES/OS backend ([@yurishkuro](https://github.com/yurishkuro) in [b9805f7](https://github.com/jaegertracing/jaeger/commit/b9805f7bc075224cfab37abab9df24ca51f38683))
* Fix CVE-2022-32149 for gotlang.org/x/text package ([@mehta-ankit](https://github.com/mehta-ankit) in [#3992](https://github.com/jaegertracing/jaeger/pull/3992))
* Expose otel configured thrift http port ([@albertteoh](https://github.com/albertteoh) in [#3986](https://github.com/jaegertracing/jaeger/pull/3986))
* Adding anchore for SBOM signing during release ([@jkowall](https://github.com/jkowall) in [#3987](https://github.com/jaegertracing/jaeger/pull/3987))
* Bump sarama to 1.33.0 ([@pavolloffay](https://github.com/pavolloffay) in [#3983](https://github.com/jaegertracing/jaeger/pull/3983))
* Add note on jaeger grpc storage compliance ([@arajkumar](https://github.com/arajkumar) in [#3985](https://github.com/jaegertracing/jaeger/pull/3985))
* Added link to FOSSA and Artifact Hub to README ([@jkowall](https://github.com/jkowall) in [#3964](https://github.com/jaegertracing/jaeger/pull/3964))
* Add grafana container to monitor docker-compose ([@albertteoh](https://github.com/albertteoh) in [#3955](https://github.com/jaegertracing/jaeger/pull/3955))
* Expose storage integration helpers as go pkg ([@arajkumar](https://github.com/arajkumar) in [#3944](https://github.com/jaegertracing/jaeger/pull/3944))

### UI

* UI pinned to version [1.27.2](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1272-2022-11-01).

1.38.1 (2022-10-04)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* Bump github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3939](https://github.com/jaegertracing/jaeger/pull/3939))
* Bump github.com/apache/thrift from 0.16.0 to 0.17.0 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3936](https://github.com/jaegertracing/jaeger/pull/3936))
* Bump github.com/hashicorp/go-hclog from 1.2.2 to 1.3.1 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3934](https://github.com/jaegertracing/jaeger/pull/3934))
* Bump go.opentelemetry.io/otel from 1.9.0 to 1.10.0 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3932](https://github.com/jaegertracing/jaeger/pull/3932))
* Bump github.com/hashicorp/go-plugin from 1.4.4 to 1.4.5 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3930](https://github.com/jaegertracing/jaeger/pull/3930))
* Bump github.com/spf13/viper from 1.12.0 to 1.13.0 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3931](https://github.com/jaegertracing/jaeger/pull/3931))
* Bump OTEL dependencies => v0.60.0 and grpc => v1.49.0 ([@yurishkuro](https://github.com/yurishkuro) in [#3928](https://github.com/jaegertracing/jaeger/pull/3928))
* Bump golang.org/x/net to 2022-09-26 ([@yurishkuro](https://github.com/yurishkuro) in [#3927](https://github.com/jaegertracing/jaeger/pull/3927))
* Bump codecov/codecov-action from 3.1.0 to 3.1.1 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3917](https://github.com/jaegertracing/jaeger/pull/3917))

### UI Changes

* UI pinned to version [1.27.1](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1271-2022-10-04) to bump dependencies.


1.38.0 (2022-09-16)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* fix: jaeger-agent sampling endpoint returns backwards incompatible JSON ([@vprithvi](https://github.com/vprithvi) in [#3897](https://github.com/jaegertracing/jaeger/pull/3897))
* fix: streaming span writer is not working in grpc based remote storage plugin ([@arajkumar](https://github.com/arajkumar) in [#3887](https://github.com/jaegertracing/jaeger/pull/3887))
* Fix race condition when adding collector tags ([@yurishkuro](https://github.com/yurishkuro) in [#3886](https://github.com/jaegertracing/jaeger/pull/3886))
* Change build info date to commit timestamp ([@TripleDogDare](https://github.com/TripleDogDare) in [#3876](https://github.com/jaegertracing/jaeger/pull/3876))
* Add 🚗 ([@yurishkuro](https://github.com/yurishkuro) in [55a8ca9](https://github.com/jaegertracing/jaeger/commit/55a8ca97e3772579b395ffbe4b937a4f5993b008))
* Add AdditionalDialOptions to ConnBuilder ([@vprithvi](https://github.com/vprithvi) in [#3865](https://github.com/jaegertracing/jaeger/pull/3865))
* Add sample docker-compose configuration using Kafka ([@yurishkuro](https://github.com/yurishkuro) in [7006e9f](https://github.com/jaegertracing/jaeger/commit/7006e9fe50c8467ad6b84f2072a3cf136bfbe4ec))

### UI Changes

* UI pinned to version [1.27.0 - see the changelog](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1270-2022-09-15).

1.37.0 (2022-08-03)
-------------------
### Backend Changes

* Add remote-storage service ([@yurishkuro](https://github.com/yurishkuro) in [#3836](https://github.com/jaegertracing/jaeger/pull/3836))

#### Bug fixes, Minor Improvements

* Fix ingester panic when span.process=nil ([@locmai](https://github.com/locmai) in [#3819](https://github.com/jaegertracing/jaeger/pull/3819))
* Added windows zip file generation ([@adhithyasrinivasan](https://github.com/adhithyasrinivasan) in [#3817](https://github.com/jaegertracing/jaeger/pull/3817))
* Refactor gRPC storage plugin for better composability ([@yurishkuro](https://github.com/yurishkuro) in [#3833](https://github.com/jaegertracing/jaeger/pull/3833))

### UI Changes

* UI pinned to version [1.26.0 - see the changelog](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1260-2022-08-03).

1.36.0 (2022-07-05)
-------------------
### Backend Changes

#### New Features
* Add flag to enable span size metrics reporting ([@ymtdzzz](https://github.com/ymtdzzz) in [#3782](https://github.com/jaegertracing/jaeger/pull/3782))
  * Span size metrics are enabled via the `--collector.enable-span-size-metrics` flag (even if `--collector.queue-size-memory` is disabled).
* Add multi-tenancy support ([@esnible](https://github.com/esnible) in [#3688](https://github.com/jaegertracing/jaeger/pull/3688))
  * Enabled when `--multi_tenancy.enabled=true` is passed to the collector.
  * The header carrying the tenants can be specified via the `--multi_tenancy.header` flag, which defaults to `x-tenant`.
  * The list of allowed tenants can be set via the `--multi_tenancy.tenants` flag, which defaults to an unrestricted list of tenants.

#### Bug fixes, Minor Improvements
* Introduce ParentReference adjuster ([@bobrik](https://github.com/bobrik) in [#3786](https://github.com/jaegertracing/jaeger/pull/3786))
* Ignore the problem of self-reported spans when multi-tenant enabled ([@esnible](https://github.com/esnible) in [#3787](https://github.com/jaegertracing/jaeger/pull/3787))
* Copy expvar metrics implementation from jaeger-lib ([@yurishkuro](https://github.com/yurishkuro) in [#3772](https://github.com/jaegertracing/jaeger/pull/3772))
* Copy Prometheus metrics implementation from jaeger-lib ([@yurishkuro](https://github.com/yurishkuro) in [#3771](https://github.com/jaegertracing/jaeger/pull/3771))
* Copy metrics API from jaeger-lib ([@yurishkuro](https://github.com/yurishkuro) in [#3767](https://github.com/jaegertracing/jaeger/pull/3767))
* Use file move instead of overwriting content ([@yurishkuro](https://github.com/yurishkuro) in [#3726](https://github.com/jaegertracing/jaeger/pull/3726))
* Refactor tenancy checking from gRPC to gRPC batch consumer ([@esnible](https://github.com/esnible) in [#3718](https://github.com/jaegertracing/jaeger/pull/3718))
* Fix ineffectual `--skip-dependencies` flag in es-rollover ([@frzifus](https://github.com/frzifus) in [#3724](https://github.com/jaegertracing/jaeger/pull/3724))
* Fix custom gogo codec to allow OTLP data ([@yurishkuro](https://github.com/yurishkuro) in [#3719](https://github.com/jaegertracing/jaeger/pull/3719))

### UI Changes

* UI pinned to version [1.25.0 - see the changelog](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1250-2022-07-03).

1.35.2 (2022-06-15)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

* fix: dependency flag not recognized ([@frzifus](https://github.com/frzifus) in [#3724](https://github.com/jaegertracing/jaeger/pull/3724))

1.35.1 (2022-06-01)
-------------------
### Backend Changes

#### Bug fixes, Minor Improvements

- Fix custom gogo codec to allow OTLP data ([@yurishkuro](https://github.com/yurishkuro) in [#3719](https://github.com/jaegertracing/jaeger/pull/3719))

1.35.0 (2022-06-01)
-------------------
### Backend Changes

#### New Features

- Introduce OTLP receiver configuration flags ([@yurishkuro](https://github.com/yurishkuro) in [#3710](https://github.com/jaegertracing/jaeger/pull/3710))
- Define Health Server for GRPC servers ([@mmorel-35](https://github.com/mmorel-35) in [#3712](https://github.com/jaegertracing/jaeger/pull/3712))
- Add OTLP receiver to collector ([@yurishkuro](https://github.com/yurishkuro) in [#3701](https://github.com/jaegertracing/jaeger/pull/3701))
- Add flag to enable/disable dependencies on rollover ([@rubenvp8510](https://github.com/rubenvp8510) in [#3705](https://github.com/jaegertracing/jaeger/pull/3705))
- Add TLS configuration for Admin Server ([@mmorel-35](https://github.com/mmorel-35) in [#3679](https://github.com/jaegertracing/jaeger/pull/3679))
- Add TLS configuration for Zipkin ([@mmorel-35](https://github.com/mmorel-35) in [#3676](https://github.com/jaegertracing/jaeger/pull/3676))

#### Bug fixes, Minor Improvements

- Fix for WithTenant() lost orig context ([@esnible](https://github.com/esnible) in [#3685](https://github.com/jaegertracing/jaeger/pull/3685))
- Add entries to env command for new storage types ([@yurishkuro](https://github.com/yurishkuro) in [#3678](https://github.com/jaegertracing/jaeger/pull/3678))
- Fix Prometheus factory signature ([@yurishkuro](https://github.com/yurishkuro) in [#3681](https://github.com/jaegertracing/jaeger/pull/3681))

### UI Changes

* UI pinned to version [1.24.0 - see the changelog](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1240-2022-06-01).


1.34.0 (2022-05-11)
-------------------
### Backend Changes

#### New Features

* Add kubernetes example for hotrod app ([@highb](https://github.com/highb) in [#3645](https://github.com/jaegertracing/jaeger/pull/3645))
* Support writing via gRPC stream in storage plugin ([@vuuihc](https://github.com/vuuihc) in [#3640](https://github.com/jaegertracing/jaeger/pull/3640))
* Instrument MetricsReader with metrics ([@albertteoh](https://github.com/albertteoh) in [#3667](https://github.com/jaegertracing/jaeger/pull/3667))

#### Bug fixes, Minor Improvements

* Sanitize spans with null process or empty service name ([@yurishkuro](https://github.com/yurishkuro) in [#3631](https://github.com/jaegertracing/jaeger/pull/3631))
* Flow tenant through processors as a string ([@esnible](https://github.com/esnible) in [#3661](https://github.com/jaegertracing/jaeger/pull/3661))
* Fix es.log-level behaviour ([@albertteoh](https://github.com/albertteoh) in [#3664](https://github.com/jaegertracing/jaeger/pull/3664))
* Mention remote gRPC storage API ([@yurishkuro](https://github.com/yurishkuro) in [cb6853d](https://github.com/jaegertracing/jaeger/commit/cb6853d4eea1ab5430f5e8db6b603cd7de5a16c3))
* Perform log.fatal if TLS flags are used when tls.enabled=false #2893 ([@clock21am](https://github.com/clock21am) in [#3030](https://github.com/jaegertracing/jaeger/pull/3030))
* Upgrade to Go 1.18 ([@yurishkuro](https://github.com/yurishkuro) in [#3624](https://github.com/jaegertracing/jaeger/pull/3624))

### UI Changes

* UI pinned to version 1.23.0. The changelog is available here [v1.23.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1230-2022-05-10).

1.33.0 (2022-04-11)
-------------------
#### New Features

* Add SAMPLING_STORAGE_TYPE environment variable ([@joe-elliott](https://github.com/joe-elliott) in [#3573](https://github.com/jaegertracing/jaeger/pull/3573))
* Support min/max TLS version in TLS config ([@Ashmita152](https://github.com/Ashmita152) in [#3567](https://github.com/jaegertracing/jaeger/pull/3567))
* Add support for ciphersuites in tls config ([@Ashmita152](https://github.com/Ashmita152) in [#3564](https://github.com/jaegertracing/jaeger/pull/3564))

#### Bug fixes, Minor Improvements

* Fix: exit on grpc plugin crash ([@johanneswuerbach](https://github.com/johanneswuerbach) in [#3604](https://github.com/jaegertracing/jaeger/pull/3604))
* Bump go.opentelemetry.io/collector/model from 0.47.0 to 0.48.0 ([@dependabot[bot]](https://github.com/apps/dependabot) in [#3608](https://github.com/jaegertracing/jaeger/pull/3608))
* Fix format string for go1.18 ([@bobrik](https://github.com/bobrik) in [#3596](https://github.com/jaegertracing/jaeger/pull/3596))
* Fix favicon returning 500 inside container ([@Ashmita152](https://github.com/Ashmita152) in [#3569](https://github.com/jaegertracing/jaeger/pull/3569))
* Elasticsearch: Do not create index templates if ILM is enabled. ([@rbizos](https://github.com/rbizos) in [#3610](https://github.com/jaegertracing/jaeger/pull/3610))


### UI Changes

* UI pinned to version 1.22.0. The changelog is available here [v1.22.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1220-2022-04-11).

1.32.0 (2022-03-06)
-------------------

### Backend Changes

#### New Features

* Enable gRPC reflection service on collector/query ([@yurishkuro](https://github.com/yurishkuro) in [#3526](https://github.com/jaegertracing/jaeger/pull/3526))

#### Bug fixes, Minor Improvements

* Fix query for latency metrics ([@Ashmita152](https://github.com/Ashmita152) in [#3559](https://github.com/jaegertracing/jaeger/pull/3559))
* Fix integration tests containing spans in the future ([@johanneswuerbach](https://github.com/johanneswuerbach) in [#3538](https://github.com/jaegertracing/jaeger/pull/3538))
* Add system diagram using mermaid markdown ([@yurishkuro](https://github.com/yurishkuro) in [#3529](https://github.com/jaegertracing/jaeger/pull/3529))
* Fix indexDateLayout for elasticsearch dependencies #3523 ([@ilyamor](https://github.com/ilyamor) in [#3524](https://github.com/jaegertracing/jaeger/pull/3524))
* Fix builds due to upstream OTEL proto path change ([@albertteoh](https://github.com/albertteoh) in [#3525](https://github.com/jaegertracing/jaeger/pull/3525))

### UI Changes

* UI pinned to version 1.21.0. The changelog is available here [v1.21.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1210-2022-03-06).

1.31.0 (2022-02-04)
-------------------
### Backend Changes
#### Bug fixes, Minor Improvements

* Bump Go compiler in CI to 1.17.6 ([@yurishkuro](https://github.com/yurishkuro) in [#3516](https://github.com/jaegertracing/jaeger/pull/3516))
* Add support for ES index aliases / rollover to the dependency store (Resolves #2143) ([@frittentheke](https://github.com/frittentheke) in [#2144](https://github.com/jaegertracing/jaeger/pull/2144))
* Use existing functions from xdg-go/scram pkg ([@yurishkuro](https://github.com/yurishkuro) in [#3481](https://github.com/jaegertracing/jaeger/pull/3481))

### UI Changes

* UI pinned to version 1.20.1. The changelog is available here [v1.20.1](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1201-2022-02-04).


1.30.0 (2022-01-11)
-------------------
### Backend Changes

#### New Features

* Add remote gRPC option for storage plugin ([@cevian](https://github.com/cevian) in [#3383](https://github.com/jaegertracing/jaeger/pull/3383))
* Build binaries for darwin/arm64 ([@jhchabran](https://github.com/jhchabran) in [#3431](https://github.com/jaegertracing/jaeger/pull/3431))
* Add MaxConnectionAge[Grace] to collector's gRPC server ([@jpkrohling](https://github.com/jpkrohling) in [#3422](https://github.com/jaegertracing/jaeger/pull/3422))

#### Bug fixes, Minor Improvements

* Fix prefixed index rollover ([@albertteoh](https://github.com/albertteoh) in [#3457](https://github.com/jaegertracing/jaeger/pull/3457))
* Log problems communicating with Elasticsearch ([@esnible](https://github.com/esnible) in [#3451](https://github.com/jaegertracing/jaeger/pull/3451))

### UI Changes

* UI pinned to version 1.20.0. The changelog is available here [v1.20.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1200-jan-11-2022)

1.29.0 (2021-12-01)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

* Remove deprecated `--badger.truncate` CLI flag ([@yurishkuro](https://github.com/yurishkuro) in [#3410](https://github.com/jaegertracing/jaeger/pull/3410))

#### New Features

* Expose rackID option in ingester ([@shyimo](https://github.com/shyimo) in [#3395](https://github.com/jaegertracing/jaeger/pull/3395))

#### Bug fixes, Minor Improvements

* Fix debug image builds by installing `build-base` to enable GCC ([@yurishkuro](https://github.com/yurishkuro) in [#3400](https://github.com/jaegertracing/jaeger/pull/3400))
* Limit URL size in Elasticsearch index delete request ([@jkandasa](https://github.com/jkandasa) in [#3375](https://github.com/jaegertracing/jaeger/pull/3375))

### UI Changes

* UI pinned to version 1.19.0. The changelog is available here [v1.19.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1190-dec-1-2021)

1.28.0 (2021-11-06)
-------------------
### Backend Changes
* Add auth token propagation for metrics reader ([@albertteoh](https://github.com/albertteoh) in [#3341](https://github.com/jaegertracing/jaeger/pull/3341))

#### New Features
* Add in-memory storage support for adaptive sampling ([@lonewolf3739](https://github.com/lonewolf3739) in [#3335](https://github.com/jaegertracing/jaeger/pull/3335))

#### Bug fixes, Minor Improvements

* Do not throw error on empty indices in Elasticsach rollover lookback ([@jkandasa](https://github.com/jkandasa) in [#3369](https://github.com/jaegertracing/jaeger/pull/3369))
* Treat input throughput data as immutable ([@rbroggi](https://github.com/rbroggi) in [#3360](https://github.com/jaegertracing/jaeger/pull/3360))
* Remove dependencies on unused tools, install tools explicitly instead of via go.mod ([@rbroggi](https://github.com/rbroggi) in [#3355](https://github.com/jaegertracing/jaeger/pull/3355))
* Update mockery to version 2 and adapt to `install-tools` approach ([@rbroggi](https://github.com/rbroggi) in [#3358](https://github.com/jaegertracing/jaeger/pull/3358))
* Control lightweight storage integration tests via build tags ([@rbroggi](https://github.com/rbroggi) in [#3346](https://github.com/jaegertracing/jaeger/pull/3346))
* Remove package `integration` from coverage reports ([@rbroggi](https://github.com/rbroggi) in [#3357](https://github.com/jaegertracing/jaeger/pull/3357))
* Remove outdated reference to cover.sh ([@rbroggi](https://github.com/rbroggi) in [#3348](https://github.com/jaegertracing/jaeger/pull/3348))
* Update monitoring mixin ([@jpkrohling](https://github.com/jpkrohling) in [#3331](https://github.com/jaegertracing/jaeger/pull/3331))
* Update Jaeger chart link ([@isbang](https://github.com/isbang) in [#3328](https://github.com/jaegertracing/jaeger/pull/3328))
* Fix args order in strings.Contains in es-rollover ([@pavolloffay](https://github.com/pavolloffay) in [#3324](https://github.com/jaegertracing/jaeger/pull/3324))
* Use `(TB).TempDir` instead of non-portable `/mnt/*` in Badger ([@pavolloffay](https://github.com/pavolloffay) in [#3325](https://github.com/jaegertracing/jaeger/pull/3325))
* Fix `peer.service` retrieval from Zipkin's `MESSAGE_ADDR` annotation ([@Git-Jiro](https://github.com/Git-Jiro) in [#3312](https://github.com/jaegertracing/jaeger/pull/3312))

### UI Changes

* UI pinned to version 1.18.0. The changelog is available here [v1.18.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1180-nov-6-2021)


1.27.0 (2021-10-06)
-------------------
### Backend Changes
* Migrate elasticsearch rollover to go ([#3242](https://github.com/jaegertracing/jaeger/pull/3242), [@rubenvp8510](https://github.com/rubenvp8510))
* Add 'opensearch' as a supported value for SPAN_STORAGE_TYPE ([#3255](https://github.com/jaegertracing/jaeger/pull/3255), [@yurishkuro](https://github.com/yurishkuro))

#### New Features
* Add support for adaptive sampling with a Cassandra backend. ([#2966](https://github.com/jaegertracing/jaeger/pull/2966), [@joe-elliott](https://github.com/joe-elliott), [@Ashmita152](https://github.com/Ashmita152))

#### Bug fixes, Minor Improvements
* Support graceful shutdown in grpc plugin ([#3249](https://github.com/jaegertracing/jaeger/pull/3249), [@slon](https://github.com/slon))
* Enable gzip compression for collector grpc endpoint. ([#3236](https://github.com/jaegertracing/jaeger/pull/3236), [@slon](https://github.com/slon))
* Use UTC in es-index-cleaner ([#3261](https://github.com/jaegertracing/jaeger/pull/3261), [@pavolloffay](https://github.com/pavolloffay))
* Upgrade to alpine-3.14 ([#3304](https://github.com/jaegertracing/jaeger/pull/3304), [@nontw](https://github.com/nontw))
* refactor: move from io/ioutil to io and os package ([#3294](https://github.com/jaegertracing/jaeger/pull/3294), [@Juneezee](https://github.com/Juneezee))
* Changed sampling type env var and updated collector help text ([#3302](https://github.com/jaegertracing/jaeger/pull/3302), [@joe-elliott](https://github.com/joe-elliott))
* Close #3270: Prevent rollover lookback from passing the Unix epoch ([#3299](https://github.com/jaegertracing/jaeger/pull/3299), [@ctreatma](https://github.com/ctreatma))
* Fixing otel configuration in docker compose ([#3286](https://github.com/jaegertracing/jaeger/pull/3286), [@Ashmita152](https://github.com/Ashmita152))
* Added ability to pass config file to grpc plugin in integration tests ([#3253](https://github.com/jaegertracing/jaeger/pull/3253), [@EinKrebs](https://github.com/EinKrebs))

### UI Changes

* UI pinned to version 1.17.0. The changelog is available here [v1.17.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1170-oct-6-2021)


1.26.0 (2021-09-06)
-------------------
### Backend Changes
#### New Features
* Add cassandra v4 support ([@Ashmita152](https://github.com/Ashmita152) in [#3225](https://github.com/jaegertracing/jaeger/pull/3225)).
  * **Note**: Users running on existing schema version [v3](https://github.com/jaegertracing/jaeger/blob/1d563f964da4b472a6f7e4cfdb85fe1167c30129/plugin/storage/cassandra/schema/v003.cql.tmpl) without issues, need not run an upgrade to [v4](https://github.com/jaegertracing/jaeger/blob/1d563f964da4b472a6f7e4cfdb85fe1167c30129/plugin/storage/cassandra/schema/v004.cql.tmpl). The new schema simply updates the table metadata, primarily removing `dclocal_read_repair_chance` deprecated table option, [which automatically gets ignored/removed when upgrading to Cassandra 4.0](https://issues.apache.org/jira/browse/CASSANDRA-13910).
* Add es-index-cleaner golang implementation ([@pavolloffay](https://github.com/pavolloffay) in [#3204](https://github.com/jaegertracing/jaeger/pull/3204))
* Add CLI Option for gRPC Max Receive Message Size ([@js8080](https://github.com/js8080) in [#3214](https://github.com/jaegertracing/jaeger/pull/3214) and [#3192](https://github.com/jaegertracing/jaeger/pull/3192))
* Automatically detect OpenSearch version ([@pavolloffay](https://github.com/pavolloffay) in [#3198](https://github.com/jaegertracing/jaeger/pull/3198))
* Add SendGetBodyAs on elasticsearch ([@NatMarchand](https://github.com/NatMarchand) in [#3193](https://github.com/jaegertracing/jaeger/pull/3193))
* Set lookback in ES rollover to distant past ([@pavolloffay](https://github.com/pavolloffay) in [#3169](https://github.com/jaegertracing/jaeger/pull/3169))

#### Bug fixes, Minor Improvements
* Check for invalid --agent.tags ([@esnible](https://github.com/esnible) in [#3246](https://github.com/jaegertracing/jaeger/pull/3246))
* Replace old linters with golangci-lint ([@albertteoh](https://github.com/albertteoh) in [#3237](https://github.com/jaegertracing/jaeger/pull/3237))
* Fix panic on empty findTraces query ([@akuzni2](https://github.com/akuzni2) in [#3232](https://github.com/jaegertracing/jaeger/pull/3232))
* Upgrade to Go v1.17 ([@Ashmita152](https://github.com/Ashmita152) in [#3220](https://github.com/jaegertracing/jaeger/pull/3220))
* Add docker buildx make target ([@pavolloffay](https://github.com/pavolloffay) in [#3213](https://github.com/jaegertracing/jaeger/pull/3213))
* Fix the name for elasticsearch integration tests ([@Ashmita152](https://github.com/Ashmita152) in [#3208](https://github.com/jaegertracing/jaeger/pull/3208))
* Upgrade ES images in integration tests ([@pavolloffay](https://github.com/pavolloffay) in [#3185](https://github.com/jaegertracing/jaeger/pull/3185))

### UI Changes
* UI pinned to version 1.16.0 - https://github.com/jaegertracing/jaeger-ui/releases/tag/v1.16.0

1.25.0 (2021-08-04)
-------------------
#### New Features
* Add query service with OTLP ([#3086](https://github.com/jaegertracing/jaeger/pull/3086), [@pavolloffay](https://github.com/pavolloffay))
* Add ppc64le support on multiarch docker images ([#3160](https://github.com/jaegertracing/jaeger/pull/3160), [@krishvoor](https://github.com/krishvoor))

#### Bug fixes, Minor Improvements
* Fix base path in grpc gateway for api_v3 ([#3139](https://github.com/jaegertracing/jaeger/pull/3139), [@pavolloffay](https://github.com/pavolloffay))
* Add /api prefix for /v3 API ([#3178](https://github.com/jaegertracing/jaeger/pull/3178), [@pavolloffay](https://github.com/pavolloffay))
* Define `http.Server.ErrorLog` to forward logs to Zap ([#3157](https://github.com/jaegertracing/jaeger/pull/3157), [@yurishkuro](https://github.com/yurishkuro))
* Add ATM dev environment docker-compose and API doc ([#3171](https://github.com/jaegertracing/jaeger/pull/3171), [@albertteoh](https://github.com/albertteoh))
* Log the source of sampling strategies ([#3166](https://github.com/jaegertracing/jaeger/pull/3166), [@yurishkuro](https://github.com/yurishkuro))
* Pin elasticsearch-py to older version without elastic.co product check ([#3180](https://github.com/jaegertracing/jaeger/pull/3180), [@pavolloffay](https://github.com/pavolloffay))

1.24.0 (2021-07-07)
-------------------
### Backend Changes

#### ⛔ Breaking Changes

* Upgrade Badger from v1.6.2 to v3.2103.0 ([#3096](https://github.com/jaegertracing/jaeger/pull/3096), [@Ashmita152](https://github.com/Ashmita152)):
  * Deprecated `--badger.truncate` flag.
  * All badger related expvar prefix has changed from `badger` to `badger_v3`.

#### New Features

* Add docker images for linux/arm64 ([#3124](https://github.com/jaegertracing/jaeger/pull/3124), [@GaruGaru](https://github.com/GaruGaru))
* Add s390x support on multiarch docker images ([#2948](https://github.com/jaegertracing/jaeger/pull/2948), [@kun-lu20](https://github.com/kun-lu20))
* Add TLS support for Prometheus reader ([#3096](https://github.com/jaegertracing/jaeger/pull/3096), [@albertteoh](https://github.com/albertteoh))

##### [Monitor tab for service health metrics](https://github.com/jaegertracing/jaeger/issues/2954)

* Add HTTP handler for metrics querying [#3095](https://github.com/jaegertracing/jaeger/pull/3095), [@albertteoh](https://github.com/albertteoh))
* Add MetricsQueryService grcp handler [#3091](https://github.com/jaegertracing/jaeger/pull/3091), [@albertteoh](https://github.com/albertteoh))
* Hook up MetricsQueryService to main funcs ([#3079](https://github.com/jaegertracing/jaeger/pull/3079), [@albertteoh](https://github.com/albertteoh))
* Add metrics query capability to query service ([#3061](https://github.com/jaegertracing/jaeger/pull/3061), [@albertteoh](https://github.com/albertteoh))

#### Bug fixes, Minor Improvements

* Add build info metrics to Jaeger components ([#3087](https://github.com/jaegertracing/jaeger/pull/3087), [@Ashmita152](https://github.com/Ashmita152))
* Upgrade gRPC to 1.38.x ([#3096](https://github.com/jaegertracing/jaeger/pull/3096), [@pavolloffay](https://github.com/pavolloffay))

1.23.0 (2021-06-04)
-------------------
### Backend Changes

#### New Features

#### ⛔ Breaking Changes

* Remove unused `--es-archive.max-span-age` flag ([#2865](https://github.com/jaegertracing/jaeger/pull/2865), [@albertteoh](https://github.com/albertteoh)):

#### New Features

* Inject trace context to grpc metadata ([#2870](https://github.com/jaegertracing/jaeger/pull/2870), [@lujiajing1126](https://github.com/lujiajing1126))
* Passing default sampling strategies file as environment variable ([#3027](https://github.com/jaegertracing/jaeger/pull/3027), [@Ashmita152](https://github.com/Ashmita152))
* [es] Add index rollover mode that can choose day and hour ([#2965](https://github.com/jaegertracing/jaeger/pull/2965), [@WalkerWang731](https://github.com/WalkerWang731))
* Add a TIMEOUT environment variable for es rollover ([#2938](https://github.com/jaegertracing/jaeger/pull/2938), [@ediezh](https://github.com/ediezh))
* Allow the ILM policy name to be configurable ([#2971](https://github.com/jaegertracing/jaeger/pull/2971), [@jrRibeiro](https://github.com/jrRibeiro))
* [es] Add remote read clusters option for cross-cluster querying ([#2874](https://github.com/jaegertracing/jaeger/pull/2874), [@dgrizzanti](https://github.com/dgrizzanti))
* Enable logging in ES client ([#2862](https://github.com/jaegertracing/jaeger/pull/2862), [@albertteoh](https://github.com/albertteoh))

#### Bug fixes, Minor Improvements

* Fix jaeger-agent reproducible memory leak ([#3050](https://github.com/jaegertracing/jaeger/pull/3050), [@jpkrohling](https://github.com/jpkrohling))
* Changed Range Query to use startTimeMillis date field instead of startTime field ([#2980](https://github.com/jaegertracing/jaeger/pull/2980), [@Sreevani871](https://github.com/Sreevani871))
* Verify FindTraces() received a query ([#2979](https://github.com/jaegertracing/jaeger/pull/2979), [@esnible](https://github.com/esnible))
* Set Content-Type in healthcheck's http response ([#2926](https://github.com/jaegertracing/jaeger/pull/2926), [@logeable](https://github.com/logeable))
* Add jaeger-query HTTP handler diagnostic logging ([#2906](https://github.com/jaegertracing/jaeger/pull/2906), [@albertteoh](https://github.com/albertteoh))
* Fix es-archive namespace default values ([#2865](https://github.com/jaegertracing/jaeger/pull/2865), [@albertteoh](https://github.com/albertteoh))

1.22.0 (2021-02-23)
-------------------

### Backend Changes

#### ⛔ Breaking Changes

* Remove deprecated TLS flags ([#2790](https://github.com/jaegertracing/jaeger/issues/2790), [@albertteoh](https://github.com/albertteoh)):
    * `--cassandra.tls` is replaced by `--cassandra.tls.enabled`
    * `--cassandra-archive.tls` is replaced by `--cassandra-archive.tls.enabled`
    * `--collector.grpc.tls` is replaced by `--collector.grpc.tls.enabled`
    * `--collector.grpc.tls.client.ca` is replaced by `--collector.grpc.tls.client-ca`
    * `--es.tls` is replaced by `--es.tls.enabled`
    * `--es-archive.tls` is replaced by `--es-archive.tls.enabled`
    * `--kafka.consumer.tls` is replaced by `--kafka.consumer.tls.enabled`
    * `--kafka.producer.tls` is replaced by `--kafka.producer.tls.enabled`
    * `--reporter.grpc.tls` is replaced by `--reporter.grpc.tls.enabled`

* Remove deprecated flags of Query Server  `--query.port` and `--query.host-port`, please use dedicated HTTP `--query.http-server.host-port` (defaults to `:16686`) and gRPC `--query.grpc-server.host-port` (defaults to `:16685`)  host-ports flags instead ([#2772](https://github.com/jaegertracing/jaeger/pull/2772), [@rjs211](https://github.com/rjs211))
    * By default, if no flags are set, the query server starts on the dedicated ports.  To use common port for gRPC and  HTTP endpoints, the host-port flags have to be explicitly set

* Remove deprecated CLI flags ([#2751](https://github.com/jaegertracing/jaeger/issues/2751), [@LostLaser](https://github.com/LostLaser)):
    * `--collector.http-port` is replaced by `--collector.http-server.host-port`
    * `--collector.grpc-port` is replaced by `--collector.grpc-server.host-port`
    * `--collector.zipkin.http-port` is replaced by `--collector.zipkin.host-port`

* Remove deprecated flags `--health-check-http-port` & `--admin-http-port`, please use `--admin.http.host-port` ([#2752](https://github.com/jaegertracing/jaeger/pull/2752), [@pradeepnnv](https://github.com/pradeepnnv))

* Remove deprecated flag `--es.max-num-spans`, please use `--es.max-doc-count` ([#2482](https://github.com/jaegertracing/jaeger/pull/2482), [@BernardTolosajr](https://github.com/BernardTolosajr))

* Remove deprecated flag `--jaeger.tags`, please use `--agent.tags` ([#2753](https://github.com/jaegertracing/jaeger/pull/2753), [@yurishkuro](https://github.com/yurishkuro))

* Remove deprecated Cassandra flags ([#2789](https://github.com/jaegertracing/jaeger/pull/2789), [@albertteoh](https://github.com/albertteoh)):
    * `--cassandra.enable-dependencies-v2` - Jaeger will automatically detect the version of the dependencies table
    * `--cassandra.tls.verify-host` - please use `--cassandra.tls.skip-host-verify` instead

* Remove incorrectly scoped downsample flags from the query service ([#2782](https://github.com/jaegertracing/jaeger/pull/2782), [@joe-elliott](https://github.com/joe-elliott))
    * `--downsampling.hashsalt` removed from jaeger-query
    * `--downsampling.ratio` removed from jaeger-query

#### New Features

* Add TLS Support for gRPC and HTTP endpoints of the Query and Collector servers ([#2337](https://github.com/jaegertracing/jaeger/pull/2337), [#2772](https://github.com/jaegertracing/jaeger/pull/2772), [#2798](https://github.com/jaegertracing/jaeger/pull/2798), [@rjs211](https://github.com/rjs211))

    *  If TLS in enabled on either or both of gRPC or HTTP endpoints, the gRPC host-port and the HTTP host-port have to be different
    *  If TLS is disabled on both endpoints, common HTTP and gRPC host-port can be explicitly set using the following host-port flags respectively:
        * Query: `--query.http-server.host-port` and  `--query.grpc-server.host-port`
        * Collector: `--collector.http-server.host-port` and `--collector.grpc-server.host-port`
* Add support for Kafka SASL/PLAIN authentication via SCRAM-SHA-256 or SCRAM-SHA-512 mechanism ([#2724](https://github.com/jaegertracing/jaeger/pull/2724), [@WalkerWang731](https://github.com/WalkerWang731))
* [agent] Add metrics to show connections status between agent and collectors ([#2657](https://github.com/jaegertracing/jaeger/pull/2657), [@WalkerWang731](https://github.com/WalkerWang731))
* Add plaintext as a supported kafka auth option ([#2721](https://github.com/jaegertracing/jaeger/pull/2721), [@pdepaepe](https://github.com/pdepaepe))
* Add ability to use JS file for UI configuration (#123 from jaeger-ui) ([#2707](https://github.com/jaegertracing/jaeger/pull/2707), [@th3M1ke](https://github.com/th3M1ke))
* Support Elasticsearch ILM for managing jaeger indices ([#2796](https://github.com/jaegertracing/jaeger/pull/2796), [@bhiravabhatla](https://github.com/bhiravabhatla))
* Push official images to quay.io, in addition to Docker Hub ([#2783](https://github.com/jaegertracing/jaeger/pull/2783), [@Ashmita152](https://github.com/Ashmita152))
* Add status command ([#2684](https://github.com/jaegertracing/jaeger/pull/2684), [@sniperking1234](https://github.com/sniperking1234))
    * Usage:
      ```bash
      $ ./cmd/collector/collector-darwin-amd64 status
      {"status":"Server available","upSince":"2021-02-19T17:57:12.671902+11:00","uptime":"25.241233383s"}
      ```
* Support configurable date separator for Elasticsearch index names ([#2637](https://github.com/jaegertracing/jaeger/pull/2637), [@sniperking1234](https://github.com/sniperking1234))

#### Bug fixes, Minor Improvements

* Use workaround for windows x509.SystemCertPool() ([#2756](https://github.com/jaegertracing/jaeger/pull/2756), [@Ashmita152](https://github.com/Ashmita152))
* Guard against mal-formed payloads received by the agent, potentially causing high memory utilization ([#2780](https://github.com/jaegertracing/jaeger/pull/2780), [@jpkrohling](https://github.com/jpkrohling))
* Expose cache TTL for ES span writer index+service ([#2737](https://github.com/jaegertracing/jaeger/pull/2737), [@necrolyte2](https://github.com/necrolyte2))
* Copy spans from memory store ([#2720](https://github.com/jaegertracing/jaeger/pull/2720), [@bobrik](https://github.com/bobrik))
* [pkg/queue] Add `StartConsumersWithFactory` function ([#2714](https://github.com/jaegertracing/jaeger/pull/2714), [@mx-psi](https://github.com/mx-psi))
* Fix potential cross-site scripting issue ([#2697](https://github.com/jaegertracing/jaeger/pull/2697), [@yurishkuro](https://github.com/yurishkuro))
* Updated gRPC Storage Plugin README with example ([#2687](https://github.com/jaegertracing/jaeger/pull/2687), [@js8080](https://github.com/js8080))
* Deduplicate collector tags ([#2658](https://github.com/jaegertracing/jaeger/pull/2658), [@Betula-L](https://github.com/Betula-L))
* Add latency metrics on collector HTTP endpoints ([#2664](https://github.com/jaegertracing/jaeger/pull/2664), [@dimitarvdimitrov](https://github.com/dimitarvdimitrov))
* Fix collector panic due to sarama sdk ([#2654](https://github.com/jaegertracing/jaeger/pull/2654), [@Betula-L](https://github.com/Betula-L))
* Handle collector Start error ([#2647](https://github.com/jaegertracing/jaeger/pull/2647), [@albertteoh](https://github.com/albertteoh))
* [anonymizer] Save trace in UI format ([#2629](https://github.com/jaegertracing/jaeger/pull/2629), [@yurishkuro](https://github.com/yurishkuro))

### UI Changes

* UI pinned to version 1.13.0. The changelog is available here [v1.13.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1130-february-20-2021)


1.21.0 (2020-11-13)
-------------------

### Backend Changes

#### New Features

* Add trace anonymizer utility ([#2621](https://github.com/jaegertracing/jaeger/pull/2621), [#2585](https://github.com/jaegertracing/jaeger/pull/2585), [@Ashmita152](https://github.com/Ashmita152))
* Add URL option for sampling strategies file ([#2519](https://github.com/jaegertracing/jaeger/pull/2519), [@goku321](https://github.com/goku321))
* Expose tunning options via expvar ([#2496](https://github.com/jaegertracing/jaeger/pull/2496), [@dstdfx](https://github.com/dstdfx))
* Support more encodings for Kafka in OTel Ingester ([#2580](https://github.com/jaegertracing/jaeger/pull/2580), [@XSAM](https://github.com/XSAM))
* Create debug docker images for jaeger backends ([#2545](https://github.com/jaegertracing/jaeger/pull/2545), [@Ashmita152](https://github.com/Ashmita152))
* Display backend & UI versions in Jaeger UI
  * Inject version info into index.html ([#2547](https://github.com/jaegertracing/jaeger/pull/2547), [@yurishkuro](https://github.com/yurishkuro))
  * Added jaeger ui version to about menu ([#606](https://github.com/jaegertracing/jaeger-ui/pull/606), [@alanisaac](https://github.com/alanisaac))

#### Bug fixes, Minor Improvements

* Update x/text to v0.3.4 ([#2625](https://github.com/jaegertracing/jaeger/pull/2625), [@objectiser](https://github.com/objectiser))
* Update CodeQL to latest best practices ([#2615](https://github.com/jaegertracing/jaeger/pull/2615), [@jhutchings1](https://github.com/jhutchings1))
* Bump opentelemetry-collector to v0.14.0 ([#2617](https://github.com/jaegertracing/jaeger/pull/2617), [@Vemmy124](https://github.com/Vemmy124))
* Bump Badger to v1.6.2 ([#2613](https://github.com/jaegertracing/jaeger/pull/2613), [@Ackar](https://github.com/Ackar))
* Fix sarama consumer deadlock ([#2587](https://github.com/jaegertracing/jaeger/pull/2587), [@albertteoh](https://github.com/albertteoh))
* Avoid deadlock if Stop is called before Serve ([#2608](https://github.com/jaegertracing/jaeger/pull/2608), [@chlunde](https://github.com/chlunde))
* Return buffers to pool on network errors or queue overflow ([#2609](https://github.com/jaegertracing/jaeger/pull/2609), [@chlunde](https://github.com/chlunde))
* Clarify deadlock panic message ([#2605](https://github.com/jaegertracing/jaeger/pull/2605), [@yurishkuro](https://github.com/yurishkuro))
* fix: don't create tags w/ empty name for internal zipkin spans ([#2596](https://github.com/jaegertracing/jaeger/pull/2596), [@mzahor](https://github.com/mzahor))
* TBufferedServer: Avoid channel close/send race on Stop ([#2583](https://github.com/jaegertracing/jaeger/pull/2583), [@chlunde](https://github.com/chlunde))
* Bumped OpenTelemetry Collector to v0.12.0 ([#2562](https://github.com/jaegertracing/jaeger/pull/2562), [@jpkrohling](https://github.com/jpkrohling))
* Disable Zipkin server if port/address is not configured ([#2554](https://github.com/jaegertracing/jaeger/pull/2554), [@yurishkuro](https://github.com/yurishkuro))
* [hotrod] Add links to traces ([#2536](https://github.com/jaegertracing/jaeger/pull/2536), [@yurishkuro](https://github.com/yurishkuro))
* OTel Cassandra/Elasticsearch Exporter queue defaults ([#2533](https://github.com/jaegertracing/jaeger/pull/2533), [@joe-elliott](https://github.com/joe-elliott))
* [otel] Update jaeger-lib to v2.4.0 ([#2538](https://github.com/jaegertracing/jaeger/pull/2538), [@dstdfx](https://github.com/dstdfx))
* Remove unnecessary ServiceName index seek if tags query is available ([#2535](https://github.com/jaegertracing/jaeger/pull/2535), [@burmanm](https://github.com/burmanm))
* Update static UI assets path in contrib doc ([#2548](https://github.com/jaegertracing/jaeger/pull/2548), [@albertteoh](https://github.com/albertteoh))

### UI Changes

* UI pinned to version 1.12.0. The changelog is available here [v1.12.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1120-november-14-2020)


1.20.0 (2020-09-29)
-------------------

### Backend Changes

#### ⛔ Breaking Changes
* Configurable ES doc count ([#2453](https://github.com/jaegertracing/jaeger/pull/2453), [@albertteoh](https://github.com/albertteoh))

    The `--es.max-num-spans` flag has been deprecated in favour of `--es.max-doc-count`.
    `--es.max-num-spans` is marked for removal in v1.21.0 as indicated in the flag description.

    If both `--es.max-num-spans` and `--es.max-doc-count` are set, the lesser of the two will be used.

    The use of `--es.max-doc-count` (which defaults to 10,000) will limit the results from all Elasticsearch
    queries by the configured value, limiting counts for Jaeger UI:

    * Services
    * Operations
    * Dependencies (edges in a dependency graph)
    * Span fetch size for a trace

* The default value for the flag `query.max-clock-skew-adjustment` has changed to `0s`, meaning that the clock skew adjustment is now disabled by default. See [#1459](https://github.com/jaegertracing/jaeger/issues/1459).

#### New Features

* Grpc plugin archive storage support ([#2317](https://github.com/jaegertracing/jaeger/pull/2317), [@m8rge](https://github.com/m8rge))
* Separate Ports for GRPC and HTTP requests in Query Server ([#2387](https://github.com/jaegertracing/jaeger/pull/2387), [@rjs211](https://github.com/rjs211))
* Configurable ES doc count ([#2453](https://github.com/jaegertracing/jaeger/pull/2453), [@albertteoh](https://github.com/albertteoh))
* Add storage metrics to OTEL, metrics by span service name ([#2431](https://github.com/jaegertracing/jaeger/pull/2431), [@pavolloffay](https://github.com/pavolloffay))

#### Bug fixes, Minor Improvements

* Increase coverage on otel/app/defaultconfig and otel/app/defaultcomponents ([#2515](https://github.com/jaegertracing/jaeger/pull/2515), [@joe-elliott](https://github.com/joe-elliott))
* Use OTEL Kafka Exporter/Receiver Instead of Jaeger Core ([#2494](https://github.com/jaegertracing/jaeger/pull/2494), [@joe-elliott](https://github.com/joe-elliott))
* Fix OTEL kafka receiver/ingester panic ([#2512](https://github.com/jaegertracing/jaeger/pull/2512), [@pavolloffay](https://github.com/pavolloffay))
* Disable clock-skew-adjustment by default. ([#2513](https://github.com/jaegertracing/jaeger/pull/2513), [@jpkrohling](https://github.com/jpkrohling))
* Fix ES OTEL status code ([#2501](https://github.com/jaegertracing/jaeger/pull/2501), [@pavolloffay](https://github.com/pavolloffay))
* OTel: Factored out Config Factory ([#2495](https://github.com/jaegertracing/jaeger/pull/2495), [@joe-elliott](https://github.com/joe-elliott))
* Fix failing ServerInUseHostPort test on MacOS ([#2477](https://github.com/jaegertracing/jaeger/pull/2477), [@albertteoh](https://github.com/albertteoh))
* Fix unmarshalling in OTEL badger ([#2488](https://github.com/jaegertracing/jaeger/pull/2488), [@pavolloffay](https://github.com/pavolloffay))
* Improve UI placeholder message ([#2487](https://github.com/jaegertracing/jaeger/pull/2487), [@yurishkuro](https://github.com/yurishkuro))
* Translate OTEL instrumentation library to ES DB model ([#2484](https://github.com/jaegertracing/jaeger/pull/2484), [@pavolloffay](https://github.com/pavolloffay))
* Add partial retry capability to OTEL ES exporter. ([#2456](https://github.com/jaegertracing/jaeger/pull/2456), [@pavolloffay](https://github.com/pavolloffay))
* Log deprecation warning only when deprecated flags are set ([#2479](https://github.com/jaegertracing/jaeger/pull/2479), [@pavolloffay](https://github.com/pavolloffay))
* Clean-up Badger's trace-not-found check ([#2481](https://github.com/jaegertracing/jaeger/pull/2481), [@yurishkuro](https://github.com/yurishkuro))
* Run the jaeger-agent as a non-root user by default ([#2466](https://github.com/jaegertracing/jaeger/pull/2466), [@chgl](https://github.com/chgl))
* Regenerate certificates to use SANs instead of Common Name ([#2461](https://github.com/jaegertracing/jaeger/pull/2461), [@albertteoh](https://github.com/albertteoh))
* Support custom port in cassandra schema creation ([#2472](https://github.com/jaegertracing/jaeger/pull/2472), [@MarianZoll](https://github.com/MarianZoll))
* Consolidated OTel ES IndexNameProviders ([#2458](https://github.com/jaegertracing/jaeger/pull/2458), [@joe-elliott](https://github.com/joe-elliott))
* Add positive confirmation that Agent made a connection to Collector (… ([#2423](https://github.com/jaegertracing/jaeger/pull/2423), [@BernardTolosajr](https://github.com/BernardTolosajr))
* Propagate TraceNotFound error from grpc storage plugins ([#2455](https://github.com/jaegertracing/jaeger/pull/2455), [@joe-elliott](https://github.com/joe-elliott))
* Use new ES reader implementation in OTEL ([#2441](https://github.com/jaegertracing/jaeger/pull/2441), [@pavolloffay](https://github.com/pavolloffay))
* Updated grpc-go to v1.29.1 ([#2445](https://github.com/jaegertracing/jaeger/pull/2445), [@jpkrohling](https://github.com/jpkrohling))
* Remove olivere elastic client from OTEL ([#2448](https://github.com/jaegertracing/jaeger/pull/2448), [@pavolloffay](https://github.com/pavolloffay))
* Use queue retry per exporter ([#2444](https://github.com/jaegertracing/jaeger/pull/2444), [@pavolloffay](https://github.com/pavolloffay))
* Add context.Context to WriteSpan ([#2436](https://github.com/jaegertracing/jaeger/pull/2436), [@yurishkuro](https://github.com/yurishkuro))
* Fix mutex unlock in storage exporters ([#2442](https://github.com/jaegertracing/jaeger/pull/2442), [@pavolloffay](https://github.com/pavolloffay))
* Add Grafana integration example ([#2408](https://github.com/jaegertracing/jaeger/pull/2408), [@fktkrt](https://github.com/fktkrt))
* Fix TLS flags settings in jaeger OTEL receiver ([#2438](https://github.com/jaegertracing/jaeger/pull/2438), [@pavolloffay](https://github.com/pavolloffay))
* Add context to dependencies endpoint ([#2434](https://github.com/jaegertracing/jaeger/pull/2434), [@yoave23](https://github.com/yoave23))
* Fix error equals ([#2429](https://github.com/jaegertracing/jaeger/pull/2429), [@albertteoh](https://github.com/albertteoh))

### UI Changes

* UI pinned to version 1.11.0. The changelog is available here [v1.11.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1110-september-28-2020)

1.19.2 (2020-08-26)
-------------------

Upgrade to a working UI version before React refactoring.


1.19.1 (2020-08-26)
-------------------

Revert UI back to 1.9 due to a bug https://github.com/jaegertracing/jaeger-ui/issues/628


1.19.0 (2020-08-26)
-------------------

### Known Issues

The pull request [#2297](https://github.com/jaegertracing/jaeger/pull/2297) aimed to add TLS support for the gRPC Query server but the flag registration is missing, so that this feature can't be used at the moment. A fix is planned for the next Jaeger version (1.20).

### Backend Changes

#### New Features

* Reload TLS certificates on change ([#2389](https://github.com/jaegertracing/jaeger/pull/2389), [@pavolloffay](https://github.com/pavolloffay))
* Add --kafka.producer.batch-min-messages collector flag ([#2371](https://github.com/jaegertracing/jaeger/pull/2371), [@prymitive](https://github.com/prymitive))
* Make UDP socket buffer size configurable ([#2336](https://github.com/jaegertracing/jaeger/pull/2336), [@kbarukhov](https://github.com/kbarukhov))
* Enable batch and queued retry processors by default ([#2330](https://github.com/jaegertracing/jaeger/pull/2330), [@pavolloffay](https://github.com/pavolloffay))
* Add trace anonymizer prototype ([#2328](https://github.com/jaegertracing/jaeger/pull/2328), [@yurishkuro](https://github.com/yurishkuro))
* Add native OTEL ES exporter ([#2295](https://github.com/jaegertracing/jaeger/pull/2295), [@pavolloffay](https://github.com/pavolloffay))
* Define busy error type in processor ([#2314](https://github.com/jaegertracing/jaeger/pull/2314), [@pavolloffay](https://github.com/pavolloffay))
* Use gRPC instead of tchannel in hotrod ([#2307](https://github.com/jaegertracing/jaeger/pull/2307), [@pavolloffay](https://github.com/pavolloffay))
* TLS support for gRPC Query server ([#2297](https://github.com/jaegertracing/jaeger/pull/2297), [@jan25](https://github.com/jan25))

#### Bug fixes, Minor Improvements

* Check missing server URL in ES OTEL client ([#2411](https://github.com/jaegertracing/jaeger/pull/2411), [@pavolloffay](https://github.com/pavolloffay))
* Fix Elasticsearch version in ES OTEL writer ([#2409](https://github.com/jaegertracing/jaeger/pull/2409), [@pavolloffay](https://github.com/pavolloffay))
* Fix go vet warnings on Go 1.15 ([#2401](https://github.com/jaegertracing/jaeger/pull/2401), [@prymitive](https://github.com/prymitive))
* Add new Elasticsearch reader implementation ([#2364](https://github.com/jaegertracing/jaeger/pull/2364), [@pavolloffay](https://github.com/pavolloffay))
* Only add the collector port if it was not explicitly set ([#2396](https://github.com/jaegertracing/jaeger/pull/2396), [@joe-elliott](https://github.com/joe-elliott))
* Fix panic in collector when Zipkin server is shutdown  ([#2392](https://github.com/jaegertracing/jaeger/pull/2392), [@Sreevani871](https://github.com/Sreevani871))
* Update validity of TLS certificates to 3650 days ([#2390](https://github.com/jaegertracing/jaeger/pull/2390), [@rjs211](https://github.com/rjs211))
* Added span and trace id to hotrod logs ([#2384](https://github.com/jaegertracing/jaeger/pull/2384), [@joe-elliott](https://github.com/joe-elliott))
* Jaeger agent logs "0" whenever sampling strategies are requested ([#2382](https://github.com/jaegertracing/jaeger/pull/2382), [@jpkrohling](https://github.com/jpkrohling))
* Fix shutdown order for collector components ([#2381](https://github.com/jaegertracing/jaeger/pull/2381), [@Sreevani871](https://github.com/Sreevani871))
* Serve jquery-3.1.1.min.js locally ([#2378](https://github.com/jaegertracing/jaeger/pull/2378), [@chaseSpace](https://github.com/chaseSpace))
* Use a single shared set of CA, client & server keys/certs for testing ([#2343](https://github.com/jaegertracing/jaeger/pull/2343), [@rjs211](https://github.com/rjs211))
* fix null pointer in jaeger-spark-dependencies ([#2327](https://github.com/jaegertracing/jaeger/pull/2327), [@moolen](https://github.com/moolen))
* Rename ARCH to TARGETARCH for multi platform build by docker buildx ([#2320](https://github.com/jaegertracing/jaeger/pull/2320), [@morlay](https://github.com/morlay))
* Mask passwords when written as json ([#2302](https://github.com/jaegertracing/jaeger/pull/2302), [@objectiser](https://github.com/objectiser))

### UI Changes

* UI pinned to version 1.10.0. The changelog is available here [v1.10.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v1100-august-25-2020)

1.18.1 (2020-06-19)
------------------

### Backend Changes

#### Security Fixes

* CVE-2020-10750: jaegertracing/jaeger: credentials leaked to container logs ([@chlunde](https://github.com/chlunde))

#### ⛔ Breaking Changes

#### New Features
* Add ppc64le support ([#2293](https://github.com/jaegertracing/jaeger/pull/2293), [@Siddhesh-Ghadi](https://github.com/Siddhesh-Ghadi))
* Expose option to enable TLS when sniffing an Elasticsearch Cluster ([#2263](https://github.com/jaegertracing/jaeger/pull/2263), [@jennynilsen](https://github.com/jennynilsen))
* Enable OTEL receiver by default ([#2279](https://github.com/jaegertracing/jaeger/pull/2279), [@pavolloffay](https://github.com/pavolloffay))
* Add Badger OTEL exporter ([#2269](https://github.com/jaegertracing/jaeger/pull/2269), [@pavolloffay](https://github.com/pavolloffay))
* Add all-in-one OTEL component ([#2262](https://github.com/jaegertracing/jaeger/pull/2262), [@pavolloffay](https://github.com/pavolloffay))
* Add arm64 support for binaries and docker images ([#2176](https://github.com/jaegertracing/jaeger/pull/2176), [@MrXinWang](https://github.com/MrXinWang))
* Add Zipkin OTEL receiver ([#2247](https://github.com/jaegertracing/jaeger/pull/2247), [@pavolloffay](https://github.com/pavolloffay))

#### Bug fixes, Minor Improvements
* Remove experimental flag from rollover ([#2290](https://github.com/jaegertracing/jaeger/pull/2290), [@pavolloffay](https://github.com/pavolloffay))
* Move ES dependencies index mapping to JSON template file ([#2285](https://github.com/jaegertracing/jaeger/pull/2285), [@pavolloffay](https://github.com/pavolloffay))
* Bump go-plugin to 1.3 ([#2261](https://github.com/jaegertracing/jaeger/pull/2261), [@yurishkuro](https://github.com/yurishkuro))
* Ignore chmod events on UI config watcher. ([#2254](https://github.com/jaegertracing/jaeger/pull/2254), [@rubenvp8510](https://github.com/rubenvp8510))
* Normalize CLI flags to use host:port addresses ([#2212](https://github.com/jaegertracing/jaeger/pull/2212), [@ohdearaugustin](https://github.com/ohdearaugustin))
* Add kafka receiver flags to ingester ([#2251](https://github.com/jaegertracing/jaeger/pull/2251), [@pavolloffay](https://github.com/pavolloffay))
* Configure Jaeger receiver and exporter by flags ([#2241](https://github.com/jaegertracing/jaeger/pull/2241), [@pavolloffay](https://github.com/pavolloffay))

### UI Changes

1.18.0 (2020-05-14)
------------------

### Backend Changes

#### ⛔ Breaking Changes

* Remove Tchannel between agent and collector ([#2115](https://github.com/jaegertracing/jaeger/pull/2115), [#2112](https://github.com/jaegertracing/jaeger/pull/2112), [@pavolloffay](https://github.com/pavolloffay))

    Remove `Tchannel` port (14267) from collector and `Tchannel` reporter from agent.

    These flags were removed from agent:
    ```
    --collector.host-port
    --reporter.tchannel.discovery.conn-check-timeout
    --reporter.tchannel.discovery.min-peers
    --reporter.tchannel.host-port
    --reporter.tchannel.report-timeout
    ```

    These flags were removed from collector:
    ```
    --collector.port
    ```

* Normalize CLI flags to use host:port addresses ([#1827](https://github.com/jaegertracing/jaeger/pull/1827), [@annanay25](https://github.com/annanay25))

  Flags previously accepting listen addresses in any other format have been deprecated:

  * `collector.port` is superseded by `collector.tchan-server.host-port`
  * `collector.http-port` is superseded by `collector.http-server.host-port`
  * `collector.grpc-port` is superseded by `collector.grpc-server.host-port`
  * `collector.zipkin.http-port` is superseded by `collector.zipkin.host-port`
  * `admin-http-port` is superseded by `admin.http.host-port`

#### New Features

* Add grpc storage plugin OTEL exporter ([#2229](https://github.com/jaegertracing/jaeger/pull/2229), [@pavolloffay](https://github.com/pavolloffay))
* Add OTEL ingester component ([#2225](https://github.com/jaegertracing/jaeger/pull/2225), [@pavolloffay](https://github.com/pavolloffay))
* Add Kafka OTEL receiver/ingester ([#2221](https://github.com/jaegertracing/jaeger/pull/2221), [@pavolloffay](https://github.com/pavolloffay))
* Create Jaeger OTEL agent component  ([#2216](https://github.com/jaegertracing/jaeger/pull/2216), [@pavolloffay](https://github.com/pavolloffay))
* Merge hardcoded/default configuration with OTEL config file ([#2211](https://github.com/jaegertracing/jaeger/pull/2211), [@pavolloffay](https://github.com/pavolloffay))
* Support periodic refresh of sampling strategies ([#2188](https://github.com/jaegertracing/jaeger/pull/2188), [@defool](https://github.com/defool))
* Add Elasticsearch OTEL exporter ([#2140](https://github.com/jaegertracing/jaeger/pull/2140), [@pavolloffay](https://github.com/pavolloffay))
* Add Cassandra OTEL exporter ([#2139](https://github.com/jaegertracing/jaeger/pull/2139), [@pavolloffay](https://github.com/pavolloffay))
* Add Kafka OTEL storage exporter ([#2135](https://github.com/jaegertracing/jaeger/pull/2135), [@pavolloffay](https://github.com/pavolloffay))
* Clock skew config ([#2119](https://github.com/jaegertracing/jaeger/pull/2119), [@joe-elliott](https://github.com/joe-elliott))
* Introduce OpenTelemetry collector ([#2086](https://github.com/jaegertracing/jaeger/pull/2086), [@pavolloffay](https://github.com/pavolloffay))
* Support regex tags search for Elasticseach backend ([#2049](https://github.com/jaegertracing/jaeger/pull/2049), [@annanay25](https://github.com/annanay25))

#### Bug fixes, Minor Improvements

* Do not skip service/operation indexing for firehose spans ([#2242](https://github.com/jaegertracing/jaeger/pull/2242), [@yurishkuro](https://github.com/yurishkuro))
* Add build information to OTEL binaries ([#2237](https://github.com/jaegertracing/jaeger/pull/2237), [@pavolloffay](https://github.com/pavolloffay))
* Enable service default sampling param ([#2230](https://github.com/jaegertracing/jaeger/pull/2230), [@defool](https://github.com/defool))
* Add Jaeger OTEL agent to docker image upload ([#2227](https://github.com/jaegertracing/jaeger/pull/2227), [@ning2008wisc](https://github.com/ning2008wisc))
* Support adding process tags in OTEL via env variable ([#2220](https://github.com/jaegertracing/jaeger/pull/2220), [@pavolloffay](https://github.com/pavolloffay))
* Bump OTEL version and update exporters to use new API ([#2196](https://github.com/jaegertracing/jaeger/pull/2196), [@pavolloffay](https://github.com/pavolloffay))
* Support sampling strategies file flag in OTEL collector ([#2195](https://github.com/jaegertracing/jaeger/pull/2195), [@pavolloffay](https://github.com/pavolloffay))
* Add zipkin receiver to OTEL collector ([#2181](https://github.com/jaegertracing/jaeger/pull/2181), [@pavolloffay](https://github.com/pavolloffay))
* Add Dockerfile for OTEL collector and publish latest tag ([#2167](https://github.com/jaegertracing/jaeger/pull/2167), [@pavolloffay](https://github.com/pavolloffay))
* Run OTEL collector without configuration file ([#2148](https://github.com/jaegertracing/jaeger/pull/2148), [@pavolloffay](https://github.com/pavolloffay))
* Update gocql to support AWS MCS ([#2133](https://github.com/jaegertracing/jaeger/pull/2133), [@johanneswuerbach](https://github.com/johanneswuerbach))
* Return appropriate gRPC errors/codes to indicate request status ([#2132](https://github.com/jaegertracing/jaeger/pull/2132), [@yurishkuro](https://github.com/yurishkuro))
* Remove tchannel port from dockerfile and test ([#2118](https://github.com/jaegertracing/jaeger/pull/2118), [@pavolloffay](https://github.com/pavolloffay))
* Remove tchannel between agent and collector ([#2115](https://github.com/jaegertracing/jaeger/pull/2115), [@pavolloffay](https://github.com/pavolloffay))
* Move all tchannel packages to a single top level package ([#2112](https://github.com/jaegertracing/jaeger/pull/2112), [@pavolloffay](https://github.com/pavolloffay))

### UI Changes

* UI pinned to version 1.9.0. The changelog is available here [v1.9.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v190-may-14-2020)

1.17.1 (2020-03-13)
------------------

#### Bug fixes, Minor Improvements

* Fix enable Kafka TLS when TLS auth is specified [#2107](https://github.com/jaegertracing/jaeger/pull/2107), [@pavolloffay](https://github.com/pavoloffay))
* Migrate project to go modules [#2098](https://github.com/jaegertracing/jaeger/pull/2098), [@pavolloffay](https://github.com/pavoloffay))
* Do not skip service/operation indexing for firehose spans [#2090](https://github.com/jaegertracing/jaeger/pull/2090), [@yurishkuro](https://github.com/yurishkuro))
* Close the span writer on main ([#2096](https://github.com/jaegertracing/jaeger/pull/2096), [@jpkrohling](https://github.com/jpkrohling))
* Improved graceful shutdown - Collector ([#2076](https://github.com/jaegertracing/jaeger/pull/2076), [@jpkrohling](https://github.com/jpkrohling))
* Improved graceful shutdown - Agent ([#2031](https://github.com/jaegertracing/jaeger/pull/2031), [@jpkrohling](https://github.com/jpkrohling))

### UI Changes

* UI pinned to version 1.8.0. The changelog is available here [v1.8.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v180-march-12-2020)

1.17.0 (2020-02-24)
------------------

### Backend Changes

#### New Features

* [tracegen] Add service name as a command line option ([#2080](https://github.com/jaegertracing/jaeger/pull/2080), [@kevinearls](https://github.com/kevinearls))
* Allow the Configuration of Additional Headers on the Jaeger Query HTTP API ([#2056](https://github.com/jaegertracing/jaeger/pull/2056), [@joe-elliott](https://github.com/joe-elliott))
* Warn about time adjustment in tags ([#2052](https://github.com/jaegertracing/jaeger/pull/2052), [@bobrik](https://github.com/bobrik))
* Add CLI flags for Kafka batching params ([#2047](https://github.com/jaegertracing/jaeger/pull/2047), [@apm-opentt](https://github.com/apm-opentt))
* Added support for dynamic queue sizes ([#1985](https://github.com/jaegertracing/jaeger/pull/1985), [@jpkrohling](https://github.com/jpkrohling))
* [agent] Process data loss stats from clients ([#2010](https://github.com/jaegertracing/jaeger/pull/2010), [@yurishkuro](https://github.com/yurishkuro))
* Add /api/sampling endpoint in collector ([#1990](https://github.com/jaegertracing/jaeger/pull/1990), [@RickyRajinder](https://github.com/RickyRajinder))
* Add basic authentication to Kafka storage ([#1983](https://github.com/jaegertracing/jaeger/pull/1983), [@chandresh-pancholi](https://github.com/chandresh-pancholi))
* Make operation_strategies part also be part of default_strategy  ([#1749](https://github.com/jaegertracing/jaeger/pull/1749), [@rutgerbrf](https://github.com/rutgerbrf))

#### Bug fixes, Minor Improvements

* Upgrade gRPC to ^1.26 ([#2077](https://github.com/jaegertracing/jaeger/pull/2077), [@yurishkuro](https://github.com/yurishkuro))
* Remove pkg/errors from dependencies ([#2073](https://github.com/jaegertracing/jaeger/pull/2073), [@yurishkuro](https://github.com/yurishkuro))
* Update dependencies, pin grpc<1.27 ([#2072](https://github.com/jaegertracing/jaeger/pull/2072), [@yurishkuro](https://github.com/yurishkuro))
* Refactor collector mains ([#2060](https://github.com/jaegertracing/jaeger/pull/2060), [@jpkrohling](https://github.com/jpkrohling))
* Clarify that "kafka" is not a real storage backend ([#2066](https://github.com/jaegertracing/jaeger/pull/2066), [@yurishkuro](https://github.com/yurishkuro))
* Added -trimpath option to go build ([#2064](https://github.com/jaegertracing/jaeger/pull/2064), [@kadern0](https://github.com/kadern0))
* Use memory size flag to activate dyn queue size feature ([#2059](https://github.com/jaegertracing/jaeger/pull/2059), [@jpkrohling](https://github.com/jpkrohling))
* Add before you push to the queue to prevent race condition on size ([#2044](https://github.com/jaegertracing/jaeger/pull/2044), [@joe-elliott](https://github.com/joe-elliott))
* Count received batches from conforming clients ([#2030](https://github.com/jaegertracing/jaeger/pull/2030), [@yurishkuro](https://github.com/yurishkuro))
* [agent] Do not increment data loss counters on the first client batch ([#2028](https://github.com/jaegertracing/jaeger/pull/2028), [@yurishkuro](https://github.com/yurishkuro))
* Allow raw port numbers for UDP servers ([#2025](https://github.com/jaegertracing/jaeger/pull/2025), [@yurishkuro](https://github.com/yurishkuro))
* Publish tracegen ([#2022](https://github.com/jaegertracing/jaeger/pull/2022), [@jpkrohling](https://github.com/jpkrohling))
* Build binaries for Linux on IBM Z / s390x architecture ([#1982](https://github.com/jaegertracing/jaeger/pull/1982), [@prankkelkar](https://github.com/prankkelkar))
* Admin/Query: Log the real port instead of the provided one to enable the use of port 0 ([#2002](https://github.com/jaegertracing/jaeger/pull/2002), [@ChadiEM](https://github.com/ChadiEM))
* Split agent's HTTP server and handler ([#1996](https://github.com/jaegertracing/jaeger/pull/1996), [@yurishkuro](https://github.com/yurishkuro))
* Clean-up collector handlers builder ([#1991](https://github.com/jaegertracing/jaeger/pull/1991), [@yurishkuro](https://github.com/yurishkuro))
* Added 'resize' operation to BoundedQueue ([#1948](https://github.com/jaegertracing/jaeger/pull/1949), [@jpkrohling](https://github.com/jpkrohling))
* Add common TLS configuration ([#1838](https://github.com/jaegertracing/jaeger/pull/1838), [@backjo](https://github.com/backjo))

### UI Changes

* UI pinned to version 1.7.0. The changelog is available here [v1.7.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v170-february-21-2020)

1.16.0 (2019-12-17)
------------------

### Backend Changes

#### ⛔ Breaking Changes

##### List of service operations can be classified by span kinds ([#1943](https://github.com/jaegertracing/jaeger/pull/1943), [#1942](https://github.com/jaegertracing/jaeger/pull/1942), [#1937](https://github.com/jaegertracing/jaeger/pull/1937), [@guo0693](https://github.com/guo0693))

* Endpoint changes:
    * Both Http & gRPC servers now take new optional parameter `spanKind` in addition to `service`. When spanKind
     is absent or empty, operations from all kinds of spans will be returned.
    * Instead of returning a list of string, both Http & gRPC servers return a list of operation struct. Please
    update your client code to process the new response. Example response:
        ```
        curl 'http://localhost:6686/api/operations?service=UserService&spanKind=server' | jq
        {
            "data": [{
                "name": "UserService::getExtendedUser",
                "spanKind": "server"
            },
            {
                "name": "UserService::getUserProfile",
                "spanKind": "server"
            }],
            "total": 2,
            "limit": 0,
            "offset": 0,
            "errors": null
        }
        ```
    * The legacy http endpoint stay untouched:
        ```
        /services/{%s}/operations
        ```
* Storage plugin changes:
    * Memory updated to support spanKind on write & read, no migration is required.
    * [Badger](https://github.com/jaegertracing/jaeger/issues/1922) & [ElasticSearch](https://github.com/jaegertracing/jaeger/issues/1923)
    to be implemented:
    For now `spanKind` will be set as empty string during read & write, only `name` will be valid operation name.
    * Cassandra updated to support spanKind on write & read ([#1937](https://github.com/jaegertracing/jaeger/pull/1937), [@guo0693](https://github.com/guo0693)):
        If you don't run the migration script, nothing will break, the system will use the old table
        `operation_names` and set empty `spanKind` in the response.
        Steps to get the updated functionality:
        1.  You will need to run the command below on the host where you can use `cqlsh` to connect to Cassandra:
            ```
            KEYSPACE=jaeger_v1 CQL_CMD='cqlsh host 9042 -u test_user -p test_password --request-timeout=3000'
            bash ./v002tov003.sh
            ```
            The script will create new table `operation_names_v2` and migrate data from the old table.
            `spanKind` column will be empty for those data.
            At the end, it will ask you whether you want to drop the old table or not.
        2. Restart ingester & query services so that they begin to use the new table

##### Trace and Span IDs are always padded to 32 or 16 hex characters with leading zeros ([#1956](https://github.com/jaegertracing/jaeger/pull/1956), [@yurishkuro](https://github.com/yurishkuro))

Previously, Jaeger backend always rendered trace and span IDs as  the shortest possible hex string, e.g. an ID
with numeric value 255 would be rendered as a string `ff`. This change makes the IDs to always render as 16 or 32
characters long hex string, e.g. the same id=255 would render as `00000000000000ff`. It mostly affects how UI
displays the IDs, the URLs, and the JSON returned from `jaeger-query` service.

Motivation: Among randomly generated and uniformly distributed trace IDs, only 1/16th of them start with 0
followed by a significant digit, 1/256th start with two 0s, and so on in decreasing geometric progression.
Therefore, trimming the leading 0s is a very modest optimization on the size of the data being transmitted or stored.

However, trimming 0s leads to ambiguities when the IDs are used as correlations with other monitoring systems,
such as logging, that treat the IDs as opaque strings and cannot establish the equivalence between padded and
unpadded IDs. It is also incompatible with W3C Trace Context and Zipkin B3 formats, both of which include all
leading 0s, so an application instrumented with OpenTelemetry SDKs may be logging different trace ID strings
than application instrumented with Jaeger SDKs (related issue #1657).

Overall, the change is backward compatible:
  * links with non-padded IDs in the UI will still work
  * data stored in Elasticsearch (where IDs are represented as strings) is still readable

However, some custom integration that rely on exact string matches of trace IDs may be broken.

##### Change default rollover conditions to 2 days ([#1963](https://github.com/jaegertracing/jaeger/pull/1963), [@pavolloffay](https://github.com/pavolloffay))

Change default rollover conditions from 7 days to 2 days.

Given that by default Jaeger uses daily indices and some organizations do not keep data longer than 7 days
the default of 7 days seems unreasonable - it might result in a too big index and
running curator would immediately remove the old index.

#### New Features

* Support collector tags, similar to agent tags ([#1854](https://github.com/jaegertracing/jaeger/pull/1854), [@radekg](https://github.com/radekg))
* Support insecure TLS and only CA cert for Elasticsearch ([#1918](https://github.com/jaegertracing/jaeger/pull/1918), [@pavolloffay](https://github.com/pavolloffay))
* Allow tracer config via env vars ([#1919](https://github.com/jaegertracing/jaeger/pull/1919), [@yurishkuro](https://github.com/yurishkuro))
* Allow turning off tags/logs indexing in Cassandra ([#1915](https://github.com/jaegertracing/jaeger/pull/1915), [@joe-elliott](https://github.com/joe-elliott))
* Blacklisting/Whitelisting tags for Cassandra indexing  ([#1904](https://github.com/jaegertracing/jaeger/pull/1904), [@joe-elliott](https://github.com/joe-elliott))

#### Bug fixes, Minor Improvements

* Support custom basepath in HotROD ([#1894](https://github.com/jaegertracing/jaeger/pull/1894), [@jan25](https://github.com/jan25))
* Deprecate tchannel reporter flags ([#1978](https://github.com/jaegertracing/jaeger/pull/1978), [@objectiser](https://github.com/objectiser))
* Do not truncate tags in Elasticsearch ([#1970](https://github.com/jaegertracing/jaeger/pull/1970), [@pavolloffay](https://github.com/pavolloffay))
* Export SaveSpan to enable multiplexing ([#1968](https://github.com/jaegertracing/jaeger/pull/1968), [@albertteoh](https://github.com/albertteoh))
* Make rollover init step idempotent ([#1964](https://github.com/jaegertracing/jaeger/pull/1964), [@pavolloffay](https://github.com/pavolloffay))
* Update python urllib3 version required by curator ([#1965](https://github.com/jaegertracing/jaeger/pull/1965), [@pavolloffay](https://github.com/pavolloffay))
* Allow changing max log level for gRPC storage plugins ([#1962](https://github.com/jaegertracing/jaeger/pull/1962), [@yyyogev](https://github.com/yyyogev))
* Fix the bug that operation_name table can not be init more than once ([#1961](https://github.com/jaegertracing/jaeger/pull/1961), [@guo0693](https://github.com/guo0693))
* Improve migration script ([#1946](https://github.com/jaegertracing/jaeger/pull/1946), [@guo0693](https://github.com/guo0693))
* Fix order of the returned results from badger backend.  ([#1939](https://github.com/jaegertracing/jaeger/pull/1939), [@burmanm](https://github.com/burmanm))
* Update python pathlib to pathlib2 ([#1930](https://github.com/jaegertracing/jaeger/pull/1930), [@objectiser](https://github.com/objectiser))
* Use proxy env vars if they're configured ([#1910](https://github.com/jaegertracing/jaeger/pull/1910), [@zoidbergwill](https://github.com/zoidbergwill))

### UI Changes

* UI pinned to version 1.6.0. The changelog is available here [v1.6.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v160-december-16-2019)

1.15.1 (2019-11-07)
------------------

##### Bug fixes, Minor Improvements

* Build platform binaries as part of CI ([#1909](https://github.com/jaegertracing/jaeger/pull/1909), [@yurishkuro](https://github.com/yurishkuro))
* Upgrade and fix dependencies ([#1907](https://github.com/jaegertracing/jaeger/pull/1907), [@yurishkuro](https://github.com/yurishkuro))


1.15.0 (2019-11-07)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

* The default value for the Ingester's flag `ingester.deadlockInterval` has been changed to `0` ([#1868](https://github.com/jaegertracing/jaeger/pull/1868), [@jpkrohling](https://github.com/jpkrohling))

  With the new default, the ingester won't `panic` if there are no messages for the last minute. To restore the previous behavior, set the flag's value to `1m`.

* Mark `--collector.grpc.tls.client.ca` flag as deprecated for jaeger-collector. ([#1840](https://github.com/jaegertracing/jaeger/pull/1840), [@yurishkuro](https://github.com/yurishkuro))

  The deprecated flag will still work until being removed, it's recommended to use `--collector.grpc.tls.client-ca` instead.

##### New Features

* Support TLS for Kafka ([#1414](https://github.com/jaegertracing/jaeger/pull/1414), [@MichaHoffmann](https://github.com/MichaHoffmann))
* Add ack and compression parameters for Kafka #1359 ([#1712](https://github.com/jaegertracing/jaeger/pull/1712), [@chandresh-pancholi](https://github.com/chandresh-pancholi))
* Propagate the bearer token to the gRPC plugin server ([#1822](https://github.com/jaegertracing/jaeger/pull/1822), [@radekg](https://github.com/radekg))
* Add Truncate and ReadOnly options for badger ([#1842](https://github.com/jaegertracing/jaeger/pull/1842), [@burmanm](https://github.com/burmanm))

##### Bug fixes, Minor Improvements

* Use correct context on ES search methods ([#1850](https://github.com/jaegertracing/jaeger/pull/1850), [@rubenvp8510](https://github.com/rubenvp8510))
* Handling of expected error codes coming from grpc storage plugins #1741 ([#1814](https://github.com/jaegertracing/jaeger/pull/1814), [@chandresh-pancholi](https://github.com/chandresh-pancholi))
* Fix ordering of indexScanKeys after TraceID parsing ([#1809](https://github.com/jaegertracing/jaeger/pull/1809), [@burmanm](https://github.com/burmanm))
* Small memory optimizations in badger write-path ([#1771](https://github.com/jaegertracing/jaeger/pull/1771), [@burmanm](https://github.com/burmanm))
* Set an empty value when a default env var value is missing ([#1777](https://github.com/jaegertracing/jaeger/pull/1777), [@jpkrohling](https://github.com/jpkrohling))
* Decouple storage dependencies and bump Go to 1.13.x ([#1886](https://github.com/jaegertracing/jaeger/pull/1886), [@yurishkuro](https://github.com/yurishkuro))
* Update gopkg.in/yaml.v2 dependency to v2.2.4 ([#1865](https://github.com/jaegertracing/jaeger/pull/1865), [@objectiser](https://github.com/objectiser))
* Upgrade jaeger-client 2.19 and jaeger-lib 2.2 and prom client 1.x ([#1810](https://github.com/jaegertracing/jaeger/pull/1810), [@yurishkuro](https://github.com/yurishkuro))
* Unpin grpc version and use serviceConfig to set the load balancer  ([#1786](https://github.com/jaegertracing/jaeger/pull/1786), [@guanw](https://github.com/guanw))

#### UI Changes

* UI pinned to version 1.5.0. The changelog is available here [v1.5.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v150-november-4-2019)

1.14.0 (2019-09-02)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

* Create ES index templates instead of indices ([#1627](https://github.com/jaegertracing/jaeger/pull/1627), [@pavolloffay](https://github.com/pavolloffay))

  This can break existing Elasticsearch deployments if security policies are applied.
  For instance Jaeger `X-Pack` configuration now requires permission to create index templates - `manage_index_templates`.

##### New Features

* Add Elasticsearch version configuration to rollover script ([#1769](https://github.com/jaegertracing/jaeger/pull/1769), [@pavolloffay](https://github.com/pavolloffay))
* Add Elasticsearch version flag ([#1753](https://github.com/jaegertracing/jaeger/pull/1753), [@pavolloffay](https://github.com/pavolloffay))
* Add Elasticsearch 7 support ([#1690](https://github.com/jaegertracing/jaeger/pull/1690), [@gregoryfranklin](https://github.com/gregoryfranklin))

  The index mappings in Elasticsearch 7 are not backwards compatible with the older versions.
  Therefore using Elasticsearch 7 with data created with older version would not work.
  Elasticsearch 6.8 supports 7.x, 6.x, 5.x compatible mappings. The upgrade has to be done
  first to ES 6.8, then apply data migration or wait until old daily indices are removed (this requires
  to start Jaeger with `--es.version=7` to force using ES 7.x mappings for newly created indices).

  Jaeger by default uses Elasticsearch ping endpoint (`/`) to derive the version which is used
  for index mappings selection. The version can be overridden by flag `--es.version`.

* Support for Zipkin Protobuf spans over HTTP ([#1695](https://github.com/jaegertracing/jaeger/pull/1695), [@jan25](https://github.com/jan25))
* Added support for hot reload of UI config ([#1688](https://github.com/jaegertracing/jaeger/pull/1688), [@jpkrohling](https://github.com/jpkrohling))
* Added base Grafana dashboard and Alert rules ([#1745](https://github.com/jaegertracing/jaeger/pull/1745), [@jpkrohling](https://github.com/jpkrohling))
* Add the jaeger-mixin for monitoring ([#1668](https://github.com/jaegertracing/jaeger/pull/1668), [@gouthamve](https://github.com/gouthamve))
* Added flags for driving cassandra connection compression through config ([#1675](https://github.com/jaegertracing/jaeger/pull/1675), [@sagaranand015](https://github.com/sagaranand015))
* Support index cleaner for rollover indices and add integration tests ([#1689](https://github.com/jaegertracing/jaeger/pull/1689), [@pavolloffay](https://github.com/pavolloffay))
* Add client TLS auth to gRPC reporter ([#1591](https://github.com/jaegertracing/jaeger/pull/1591), [@tcolgate](https://github.com/tcolgate))
* Collector kafka producer protocol version config ([#1658](https://github.com/jaegertracing/jaeger/pull/1658), [@marqc](https://github.com/marqc))
* Configurable kafka protocol version for msg consuming by jaeger ingester ([#1640](https://github.com/jaegertracing/jaeger/pull/1640), [@marqc](https://github.com/marqc))
* Use credentials when describing keyspaces in cassandra schema builder ([#1655](https://github.com/jaegertracing/jaeger/pull/1655), [@MiLk](https://github.com/MiLk))
* Add connect-timeout for Cassandra ([#1647](https://github.com/jaegertracing/jaeger/pull/1647), [@sagaranand015](https://github.com/sagaranand015))

##### Bug fixes, Minor Improvements

* Fix gRPC over cmux and add unit tests ([#1758](https://github.com/jaegertracing/jaeger/pull/1758), [@yurishkuro](https://github.com/yurishkuro))
* Add CA certificates to agent image ([#1764](https://github.com/jaegertracing/jaeger/pull/1764), [@yurishkuro](https://github.com/yurishkuro))
* Fix badger merge-join algorithm to correctly filter indexes ([#1721](https://github.com/jaegertracing/jaeger/pull/1721), [@burmanm](https://github.com/burmanm))
* Change Zipkin CORS origins and headers to comma separated list ([#1556](https://github.com/jaegertracing/jaeger/pull/1556), [@JonasVerhofste](https://github.com/JonasVerhofste))
* Added null guards to 'Process' when processing an incoming span ([#1723](https://github.com/jaegertracing/jaeger/pull/1723), [@jpkrohling](https://github.com/jpkrohling))
* Export expvar metrics of badger to the metricsFactory ([#1704](https://github.com/jaegertracing/jaeger/pull/1704), [@burmanm](https://github.com/burmanm))
* Pass TTL as int, not as float64 ([#1710](https://github.com/jaegertracing/jaeger/pull/1710), [@yurishkuro](https://github.com/yurishkuro))
* Use find by regex for archive index in index cleaner ([#1693](https://github.com/jaegertracing/jaeger/pull/1693), [@pavolloffay](https://github.com/pavolloffay))
* Allow token propagation if token type is not specified ([#1685](https://github.com/jaegertracing/jaeger/pull/1685), [@rubenvp8510](https://github.com/rubenvp8510))
* Fix duplicated spans when querying Elasticsearch ([#1677](https://github.com/jaegertracing/jaeger/pull/1677), [@pavolloffay](https://github.com/pavolloffay))
* Fix the threshold precision issue ([#1665](https://github.com/jaegertracing/jaeger/pull/1665), [@guanw](https://github.com/guanw))
* Badger filter duplicate results from a single indexSeek ([#1649](https://github.com/jaegertracing/jaeger/pull/1649), [@burmanm](https://github.com/burmanm))
* Badger make default dirs work in Windows ([#1653](https://github.com/jaegertracing/jaeger/pull/1653), [@burmanm](https://github.com/burmanm))

#### UI Changes

* UI pinned to version 1.4.0. The changelog is available here [v1.4.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v130-june-21-2019)

1.13.1 (2019-06-28)
------------------

#### Backend Changes

##### Bug fixes, Minor Improvements

* Change default for bearer-token-propagation to false ([#1642](https://github.com/jaegertracing/jaeger/pull/1642), [@wsoula](https://github.com/wsoula))

#### UI Changes

1.13.0 (2019-06-27)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

* The traces related metrics on collector now have a new tag `sampler_type` ([#1576](https://github.com/jaegertracing/jaeger/pull/1576), [@guanw](https://github.com/guanw))

  This might break some existing metrics dashboard (if so, users need to update query to aggregate over this new tag).

  The list of metrics affected: `traces.received`, `traces.rejected`, `traces.saved-by-svc`.

* Remove deprecated index prefix separator `:` from Elastic ([#1620](https://github.com/jaegertracing/jaeger/pull/1620), [@pavolloffay](https://github.com/pavolloffay))

  In Jaeger 1.9.0 release the Elasticsearch index separator was changed from `:` to `-`. To keep backwards
  compatibility the query service kept querying indices with `:` separator, however the new indices
  were created only with `-`. This release of Jaeger removes the query capability for indices containing `:`,
  therefore it's recommended to keep using older version until indices containing old separator are
  not queried anymore.

##### New Features

* Passthrough OAuth bearer token supplied to Query service through to ES storage ([#1599](https://github.com/jaegertracing/jaeger/pull/1599), [@rubenvp8510](https://github.com/rubenvp8510))
* Kafka kerberos authentication support for collector/ingester ([#1589](https://github.com/jaegertracing/jaeger/pull/1589), [@rubenvp8510](https://github.com/rubenvp8510))
* Allow Cassandra schema builder to use credentials ([#1635](https://github.com/jaegertracing/jaeger/pull/1635), [@PS-EGHornbostel](https://github.com/PS-EGHornbostel))
* Add docs generation command ([#1572](https://github.com/jaegertracing/jaeger/pull/1572), [@pavolloffay](https://github.com/pavolloffay))

##### Bug fixes, Minor Improvements

* Fix data race between `Agent.Run()` and `Agent.Stop()` ([#1625](https://github.com/jaegertracing/jaeger/pull/1625), [@tigrannajaryan](https://github.com/tigrannajaryan))
* Use json number when unmarshalling data from ES ([#1618](https://github.com/jaegertracing/jaeger/pull/1618), [@pavolloffay](https://github.com/pavolloffay))
* Define logs as nested data type ([#1622](https://github.com/jaegertracing/jaeger/pull/1622), [@pavolloffay](https://github.com/pavolloffay))
* Fix archive storage not querying old spans older than maxSpanAge ([#1617](https://github.com/jaegertracing/jaeger/pull/1617), [@pavolloffay](https://github.com/pavolloffay))
* Query service: fix logging errors on SIGINT ([#1601](https://github.com/jaegertracing/jaeger/pull/1601), [@jan25](https://github.com/jan25))
* Direct grpc logs to Zap logger ([#1606](https://github.com/jaegertracing/jaeger/pull/1606), [@yurishkuro](https://github.com/yurishkuro))
* Fix sending status to health check channel in Query service ([#1598](https://github.com/jaegertracing/jaeger/pull/1598), [@jan25](https://github.com/jan25))
* Add tmp-volume to all-in-one image to fix badger storage ([#1571](https://github.com/jaegertracing/jaeger/pull/1571), [@burmanm](https://github.com/burmanm))
* Do not fail es-cleaner if there are no jaeger indices ([#1569](https://github.com/jaegertracing/jaeger/pull/1569), [@pavolloffay](https://github.com/pavolloffay))
* Automatically set `GOMAXPROCS` ([#1560](https://github.com/jaegertracing/jaeger/pull/1560), [@rubenvp8510](https://github.com/rubenvp8510))
* Add CA certs to all-in-one image ([#1554](https://github.com/jaegertracing/jaeger/pull/1554), [@chandresh-pancholi](https://github.com/chandresh-pancholi))

#### UI Changes

* UI pinned to version 1.3.0. The changelog is available here [v1.3.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v130-june-21-2019)

1.12.0 (2019-05-16)
------------------

#### Backend Changes

##### ⛔ Breaking Changes
- The `kafka` flags were removed in favor of `kafka.producer` and `kafka.consumer` flags ([#1424](https://github.com/jaegertracing/jaeger/pull/1424), [@ledor473](https://github.com/ledor473))

    The following flags have been **removed** in the Collector and the Ingester:
    ```
    --kafka.brokers
    --kafka.encoding
    --kafka.topic
    --ingester.brokers
    --ingester.encoding
    --ingester.topic
    --ingester.group-id
    ```

    In the Collector, they are replaced by:
    ```
    --kafka.producer.brokers
    --kafka.producer.encoding
    --kafka.producer.topic
    ```

    In the Ingester, they are replaced by:
    ```
    --kafka.consumer.brokers
    --kafka.consumer.encoding
    --kafka.consumer.topic
    --kafka.consumer.group-id
    ```

* Add Admin port and group all ports in one file ([#1442](https://github.com/jaegertracing/jaeger/pull/1442), [@yurishkuro](https://github.com/yurishkuro))

    This change fixes issues [#1428](https://github.com/jaegertracing/jaeger/issues/1428), [#1332](https://github.com/jaegertracing/jaeger/issues/1332) and moves all metrics endpoints from API ports to **admin ports**. It requires re-configuring Prometheus scraping rules. Each Jaeger binary has its own admin port that can be found under `--admin-http-port` command line flag by running the `${binary} help` command.

##### New Features

* Add gRPC resolver using external discovery service ([#1498](https://github.com/jaegertracing/jaeger/pull/1498), [@guanw](https://github.com/guanw))
* gRPC storage plugin framework ([#1461](https://github.com/jaegertracing/jaeger/pull/1461), [@chvck](https://github.com/chvck))
* Supports customized kafka client id ([#1507](https://github.com/jaegertracing/jaeger/pull/1507), [@newly12](https://github.com/newly12))
* Support gRPC for query service ([#1307](https://github.com/jaegertracing/jaeger/pull/1307), [@annanay25](https://github.com/annanay25))
* Expose tls.InsecureSkipVerify to es.tls.* CLI flags ([#1473](https://github.com/jaegertracing/jaeger/pull/1473), [@stefanvassilev](https://github.com/stefanvassilev))
* Return info msg for `/health` endpoint ([#1465](https://github.com/jaegertracing/jaeger/pull/1465), [@stefanvassilev](https://github.com/stefanvassilev))
* Add pprof endpoint to admin endpoint ([#1375](https://github.com/jaegertracing/jaeger/pull/1375), [@konradgaluszka](https://github.com/konradgaluszka))
* Add inbound transport as label to collector metrics [#1446](https://github.com/jaegertracing/jaeger/pull/1446) ([guanw](https://github.com/guanw))
* Sorted key/value store `badger` backed storage plugin ([#760](https://github.com/jaegertracing/jaeger/pull/760), [@burmanm](https://github.com/burmanm))
* Add Admin port and group all ports in one file ([#1442](https://github.com/jaegertracing/jaeger/pull/1442), [@yurishkuro](https://github.com/yurishkuro))
* Adds support for agent level tag ([#1396](https://github.com/jaegertracing/jaeger/pull/1396), [@annanay25](https://github.com/annanay25))
* Add a Downsampling writer that drop a percentage of spans ([#1353](https://github.com/jaegertracing/jaeger/pull/1353), [@guanw](https://github.com/guanw))

##### Bug fixes, Minor Improvements

* Sort traces in memory store to return most recent traces ([#1394](https://github.com/jaegertracing/jaeger/pull/1394), [@jacobmarble](https://github.com/jacobmarble))
* Add span format tag for jaeger-collector ([#1493](https://github.com/jaegertracing/jaeger/pull/1493), [@guo0693](https://github.com/guo0693))
* Upgrade gRPC to 1.20.1 ([#1492](https://github.com/jaegertracing/jaeger/pull/1492), [@guanw](https://github.com/guanw))
* Switch from counter to a gauge for partitions held ([#1485](https://github.com/jaegertracing/jaeger/pull/1485), [@bobrik](https://github.com/bobrik))
* Add CORS handling for Zipkin collector service ([#1463](https://github.com/jaegertracing/jaeger/pull/1463), [@JonasVerhofste](https://github.com/JonasVerhofste))
* Check elasticsearch nil response ([#1467](https://github.com/jaegertracing/jaeger/pull/1467), [@YEXINGZHE54](https://github.com/YEXINGZHE54))
* Disable sampling in logger - `zap`([#1460](https://github.com/jaegertracing/jaeger/pull/1460), [@psinghal20](https://github.com/psinghal20))
* New layout for proto definitions and generated files ([#1427](https://github.com/jaegertracing/jaeger/pull/1427), [@annanay25](https://github.com/annanay25))
* Upgrade Go to 1.12.1 ([#1437](https://github.com/jaegertracing/jaeger/pull/1437) ,[@yurishkuro](https://github.com/yurishkuro))

#### UI Changes

* UI pinned to version 1.2.0. The changelog is available here [v1.2.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v120-may-14-2019)

1.11.0 (2019-03-07)
------------------

#### Backend Changes

##### ⛔ Breaking Changes
- Introduce `kafka.producer` and `kafka.consumer` flags to replace `kafka` flags ([#1360](https://github.com/jaegertracing/jaeger/pull/1360), [@ledor473](https://github.com/ledor473))

    The following flags have been deprecated in the Collector and the Ingester:
    ```
    --kafka.brokers
    --kafka.encoding
    --kafka.topic
    ```

    In the Collector, they are replaced by:
    ```
    --kafka.producer.brokers
    --kafka.producer.encoding
    --kafka.producer.topic
    ```

    In the Ingester, they are replaced by:
    ```
    --kafka.consumer.brokers
    --kafka.consumer.encoding
    --kafka.consumer.group-id
    ```
##### New Features

- Support secure gRPC channel between agent and collector ([#1391](https://github.com/jaegertracing/jaeger/pull/1391), [@ghouscht](https://github.com/ghouscht), [@yurishkuro](https://github.com/yurishkuro))
- Allow to use TLS with ES basic auth ([#1388](https://github.com/jaegertracing/jaeger/pull/1388), [@pavolloffay](https://github.com/pavolloffay))

##### Bug fixes, Minor Improvements

- Make `esRollover.py init` idempotent ([#1407](https://github.com/jaegertracing/jaeger/pull/1407) and [#1408](https://github.com/jaegertracing/jaeger/pull/1408), [@pavolloffay](https://github.com/pavolloffay))
- Allow thrift reporter if grpc hosts are not provided ([#1400](https://github.com/jaegertracing/jaeger/pull/1400), [@pavolloffay](https://github.com/pavolloffay))
- Deprecate colon in index prefix in ES dependency store ([#1386](https://github.com/jaegertracing/jaeger/pull/1386), [@pavolloffay](https://github.com/pavolloffay))
- Make grpc reporter default and add retry ([#1384](https://github.com/jaegertracing/jaeger/pull/1384), [@pavolloffay](https://github.com/pavolloffay))
- Use `CQLSH_HOST` in final call to `cqlsh` ([#1372](https://github.com/jaegertracing/jaeger/pull/1372), [@funny-falcon](https://github.com/funny-falcon))

#### UI Changes

* UI pinned to version 1.1.0. The changelog is available here [v1.1.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v110-march-3-2019)


1.10.1 (2019-02-21)
------------------

#### Backend Changes

- Discover dependencies table version automatically ([#1364](https://github.com/jaegertracing/jaeger/pull/1364), [@black-adder](https://github.com/black-adder))

##### Bug fixes, Minor Improvements

- Separate query-service functionality from http handler ([#1312](https://github.com/jaegertracing/jaeger/pull/1312), [@annanay25](https://github.com/annanay25))

#### UI Changes


1.10.0 (2019-02-15)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

- Remove cassandra SASI indices ([#1328](https://github.com/jaegertracing/jaeger/pull/1328), [@black-adder](https://github.com/black-adder))

Migration Path:

1. Run `plugin/storage/cassandra/schema/migration/v001tov002part1.sh` which will copy dependencies into a csv, update the `dependency UDT`, create a new `dependencies_v2` table, and write dependencies from the csv into the `dependencies_v2` table.
2. Run the collector and query services with the cassandra flag `cassandra.enable-dependencies-v2=true` which will instruct jaeger to write and read to and from the new `dependencies_v2` table.
3. Update [spark job](https://github.com/jaegertracing/spark-dependencies) to write to the new `dependencies_v2` table. The feature will be done in [#58](https://github.com/jaegertracing/spark-dependencies/issues/58).
4. Run `plugin/storage/cassandra/schema/migration/v001tov002part2.sh` which will DELETE the old dependency table and the SASI index.

Users who wish to continue to use the v1 table don't have to do anything as the cassandra flag `cassandra.enable-dependencies-v2` will default to false. Users may migrate on their own timeline however new features will be built solely on the `dependencies_v2` table. In the future, we will remove support for v1 completely.

- Remove `ErrorBusy`, it essentially duplicates `SpansDropped` ([#1091](https://github.com/jaegertracing/jaeger/pull/1091), [@cstyan](https://github.com/cstyan))

##### New Features

- Support certificates in elasticsearch scripts ([#1339](https://github.com/jaegertracing/jaeger/pull/1399), [@pavolloffay](https://github.com/pavolloffay))
- Add ES Rollover support to main indices ([#1309](https://github.com/jaegertracing/jaeger/pull/1309), [@pavolloffay](https://github.com/pavolloffay))
- Load ES auth token from file ([#1319](https://github.com/jaegertracing/jaeger/pull/1319), [@pavolloffay](https://github.com/pavolloffay))
- Add username/password authentication to ES index cleaner ([#1318](https://github.com/jaegertracing/jaeger/pull/1318), [@gregoryfranklin](https://github.com/gregoryfranklin))
- Add implementation of FindTraceIDs function for Elasticsearch reader ([#1280](https://github.com/jaegertracing/jaeger/pull/1280), [@vlamug](https://github.com/vlamug))
- Support archive traces for ES storage ([#1197](https://github.com/jaegertracing/jaeger/pull/1197), [@pavolloffay](https://github.com/pavolloffay))

##### Bug fixes, Minor Improvements

- Use Zipkin annotations if the timestamp is zero ([#1341](https://github.com/jaegertracing/jaeger/pull/1341), [@geobeau](https://github.com/geobeau))
- Use GRPC round robin balancing even if only one hostname ([#1329](https://github.com/jaegertracing/jaeger/pull/1329), [@benley](https://github.com/benley))
- Tolerate whitespaces in ES servers and kafka brokers ([#1305](https://github.com/jaegertracing/jaeger/pull/1305), [@verma-varsha](https://github.com/verma-varsha))
- Let cassandra servers contain whitespace in config ([#1301](https://github.com/jaegertracing/jaeger/pull/1301), [@karlpokus](https://github.com/karlpokus))

#### UI Changes


1.9.0 (2019-01-21)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

- Change Elasticsearch index prefix from `:` to `-` ([#1284](https://github.com/jaegertracing/jaeger/pull/1284), [@pavolloffay](https://github.com/pavolloffay))

Changed index prefix separator from `:`  to `-` because Elasticsearch 7 does not allow `:` in index name.
Jaeger query still reads from old indices containing `-` as a separator, therefore no configuration or migration changes are required.



- Add CLI configurable `es.max-num-spans` while retrieving spans from ES ([#1283](https://github.com/jaegertracing/jaeger/pull/1283), [@annanay25](https://github.com/annanay25))

The default value is set to 10000. Before no limit was applied.


- Update to jaeger-lib 2 and latest sha for jaeger-client-go, to pick up refactored metric names ([#1282](https://github.com/jaegertracing/jaeger/pull/1282), [@objectiser](https://github.com/objectiser))

Update to latest version of `jaeger-lib`, which includes a change to the naming of counters exported to
prometheus, to follow the convention of using a `_total` suffix, e.g. `jaeger_query_requests` is now
`jaeger_query_requests_total`.

Jaeger go client metrics, previously under the namespace `jaeger_client_jaeger_` are now under
`jaeger_tracer_`.


- Add gRPC metrics to agent ([#1180](https://github.com/jaegertracing/jaeger/pull/1180), [@pavolloffay](https://github.com/pavolloffay))

The following metrics:
```
jaeger_agent_tchannel_reporter_batch_size{format="jaeger"} 0
jaeger_agent_tchannel_reporter_batch_size{format="zipkin"} 0
jaeger_agent_tchannel_reporter_batches_failures{format="jaeger"} 0
jaeger_agent_tchannel_reporter_batches_failures{format="zipkin"} 0
jaeger_agent_tchannel_reporter_batches_submitted{format="jaeger"} 0
jaeger_agent_tchannel_reporter_batches_submitted{format="zipkin"} 0
jaeger_agent_tchannel_reporter_spans_failures{format="jaeger"} 0
jaeger_agent_tchannel_reporter_spans_failures{format="zipkin"} 0
jaeger_agent_tchannel_reporter_spans_submitted{format="jaeger"} 0
jaeger_agent_tchannel_reporter_spans_submitted{format="zipkin"} 0

jaeger_agent_collector_proxy{endpoint="baggage",result="err"} 0
jaeger_agent_collector_proxy{endpoint="baggage",result="ok"} 0
jaeger_agent_collector_proxy{endpoint="sampling",result="err"} 0
jaeger_agent_collector_proxy{endpoint="sampling",result="ok"} 0
```
have been renamed to:
```
jaeger_agent_reporter_batch_size{format="jaeger",protocol="tchannel"} 0
jaeger_agent_reporter_batch_size{format="zipkin",protocol="tchannel"} 0
jaeger_agent_reporter_batches_failures{format="jaeger",protocol="tchannel"} 0
jaeger_agent_reporter_batches_failures{format="zipkin",protocol="tchannel"} 0
jaeger_agent_reporter_batches_submitted{format="jaeger",protocol="tchannel"} 0
jaeger_agent_reporter_batches_submitted{format="zipkin",protocol="tchannel"} 0
jaeger_agent_reporter_spans_failures{format="jaeger",protocol="tchannel"} 0
jaeger_agent_reporter_spans_failures{format="zipkin",protocol="tchannel"} 0
jaeger_agent_reporter_spans_submitted{format="jaeger",protocol="tchannel"} 0
jaeger_agent_reporter_spans_submitted{format="zipkin",protocol="tchannel"} 0

jaeger_agent_collector_proxy{endpoint="baggage",protocol="tchannel",result="err"} 0
jaeger_agent_collector_proxy{endpoint="baggage",protocol="tchannel",result="ok"} 0
jaeger_agent_collector_proxy{endpoint="sampling",protocol="tchannel",result="err"} 0
jaeger_agent_collector_proxy{endpoint="sampling",protocol="tchannel",result="ok"} 0
```

- Rename tcollector proxy metric in agent ([#1182](https://github.com/jaegertracing/jaeger/pull/1182), [@pavolloffay](https://github.com/pavolloffay))

The following metric:
```
jaeger_http_server_errors{source="tcollector-proxy",status="5xx"}
```
has been renamed to:
```
jaeger_http_server_errors{source="collector-proxy",status="5xx"}
```

##### New Features

- Add tracegen utility for generating traces ([#1245](https://github.com/jaegertracing/jaeger/pull/1245), [@yurishkuro](https://github.com/yurishkuro))
- Use DCAwareRoundRobinPolicy as fallback for TokenAwarePolicy ([#1285](https://github.com/jaegertracing/jaeger/pull/1285), [@vprithvi](https://github.com/vprithvi))
- Add Zipkin Thrift as kafka ingestion format ([#1256](https://github.com/jaegertracing/jaeger/pull/1256), [@geobeau](https://github.com/geobeau))
- Add `FindTraceID` to the spanstore interface ([#1246](https://github.com/jaegertracing/jaeger/pull/1246), [@vprithvi](https://github.com/vprithvi))
- Migrate from glide to dep ([#1240](https://github.com/jaegertracing/jaeger/pull/1240), [@isaachier](https://github.com/isaachier))
- Make tchannel timeout for reporting in agent configurable ([#1034](https://github.com/jaegertracing/jaeger/pull/1034), [@gouthamve](https://github.com/gouthamve))
- Add archive traces to all-in-one ([#1189](https://github.com/jaegertracing/jaeger/pull/1189), [@pavolloffay](https://github.com/pavolloffay))
- Start moving components of adaptive sampling to OSS ([#973](https://github.com/jaegertracing/jaeger/pull/973), [@black-adder](https://github.com/black-adder))
- Add gRPC communication between agent and collector ([#1165](https://github.com/jaegertracing/jaeger/pull/1165), [#1187](https://github.com/jaegertracing/jaeger/pull/1187), [#1181](https://github.com/jaegertracing/jaeger/pull/1181) and [#1180](https://github.com/jaegertracing/jaeger/pull/1180), [@pavolloffay](https://github.com/pavolloffay))

##### Bug fixes, Minor Improvements

- Update exposed ports in ingester dockerfile ([#1289](https://github.com/jaegertracing/jaeger/pull/1289), [@objectiser](https://github.com/objectiser))
- Upgrade Shopify/Sarama for proper handling newest kafka servers 2.x by ingester ([#1248](https://github.com/jaegertracing/jaeger/pull/1248), [@vprithvi](https://github.com/vprithvi))
- Fix sampling strategies overwriting service entry when no sampling type is specified ([#1244](https://github.com/jaegertracing/jaeger/pull/1244), [@objectiser](https://github.com/objectiser))
- Fix dot replacement for int ([#1272](https://github.com/jaegertracing/jaeger/pull/1272), [@pavolloffay](https://github.com/pavolloffay))
- Add C* query to error logs ([#1250](https://github.com/jaegertracing/jaeger/pull/1250), [@vprithvi](https://github.com/vprithvi))
- Add locking around partitionIDToState map accesses ([#1239](https://github.com/jaegertracing/jaeger/pull/1239), [@vprithvi](https://github.com/vprithvi))
- Reorganize config manager packages in agent ([#1198](https://github.com/jaegertracing/jaeger/pull/1198), [@pavolloffay](https://github.com/pavolloffay))

#### UI Changes

* UI pinned to version 1.0.0. The changelog is available here [v1.0.0](https://github.com/jaegertracing/jaeger-ui/blob/main/CHANGELOG.md#v100-january-18-2019)

1.8.2 (2018-11-28)
------------------

#### UI Changes

##### New Features

- Embedded components (SearchTraces and Tracepage) ([#263](https://github.com/jaegertracing/jaeger/pull/263), [@aljesusg](https://github.com/aljesusg))

##### Bug fixes, Minor Improvements

- Fix link in scatter plot when embed mode ([#283](https://github.com/jaegertracing/jaeger-ui/pull/283), [@aljesusg](https://github.com/aljesusg))
- Fix rendering X axis in TraceResultsScatterPlot - pass milliseconds to moment.js ([#274](https://github.com/jaegertracing/jaeger-ui/pull/274), [@istrel](https://github.com/istrel))


1.8.1 (2018-11-23)
------------------

#### Backend Changes

##### Bug fixes, Minor Improvements

- Make agent timeout for reporting configurable and fix flags overriding ([#1034](https://github.com/jaegertracing/jaeger/pull/1034), [@gouthamve](https://github.com/gouthamve))
- Fix metrics handler registration in agent ([#1178](https://github.com/jaegertracing/jaeger/pull/1178), [@pavolloffay](https://github.com/pavolloffay))


1.8.0 (2018-11-12)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

- Refactor agent configuration ([#1092](https://github.com/jaegertracing/jaeger/pull/1092), [@pavolloffay](https://github.com/pavolloffay))

The following agent flags has been deprecated in order to support multiple reporters:
```bash
--collector.host-port
--discovery.conn-check-timeout
--discovery.min-peers
```
New flags:
```bash
--reporter.tchannel.host-port
--reporter.tchannel.discovery.conn-check-timeout
--reporter.tchannel.discovery.min-peers
```

- Various changes around metrics produced by jaeger-query: Names scoped to the query component, generated for all span readers (not just ES), consolidate query metrics and include result tag ([#1074](https://github.com/jaegertracing/jaeger/pull/1074), [#1075](https://github.com/jaegertracing/jaeger/pull/1075) and [#1096](https://github.com/jaegertracing/jaeger/pull/1096), [@objectiser](https://github.com/objectiser))

For example, sample of metrics produced for `find_traces` operation before:

```
jaeger_find_traces_attempts 1
jaeger_find_traces_errLatency_bucket{le="0.005"} 0
jaeger_find_traces_errors 0
jaeger_find_traces_okLatency_bucket{le="0.005"} 0
jaeger_find_traces_responses_bucket{le="0.005"} 1
jaeger_find_traces_successes 1
```

And now:

```
jaeger_query_latency_bucket{operation="find_traces",result="err",le="0.005"} 0
jaeger_query_latency_bucket{operation="find_traces",result="ok",le="0.005"} 2
jaeger_query_requests{operation="find_traces",result="err"} 0
jaeger_query_requests{operation="find_traces",result="ok"} 2
jaeger_query_responses_bucket{operation="find_traces",le="0.005"} 2
```

##### New Features

- Configurable deadlock detector interval for ingester ([#1134](https://github.com/jaegertracing/jaeger/pull/1134), [@marqc](https://github.com/marqc))
- Emit spans for elastic storage backend ([#1128](https://github.com/jaegertracing/jaeger/pull/1128), [@annanay25](https://github.com/annanay25))
- Allow to use TLS certificates for Elasticsearch authentication ([#1139](https://github.com/jaegertracing/jaeger/pull/1139), [@clyang82](https://github.com/clyang82))
- Add ingester metrics, healthcheck and rename Kafka cli flags ([#1094](https://github.com/jaegertracing/jaeger/pull/1094), [@ledor473](https://github.com/ledor473))
- Add a metric for number of partitions held ([#1154](https://github.com/jaegertracing/jaeger/pull/1154), [@vprithvi](https://github.com/vprithvi))
- Log jaeger-collector tchannel port ([#1136](https://github.com/jaegertracing/jaeger/pull/1136), [@mindaugasrukas](https://github.com/mindaugasrukas))
- Support tracer env based initialization in hotrod ([#1115](https://github.com/jaegertracing/jaeger/pull/1115), [@eundoosong](https://github.com/eundoosong))
- Publish ingester as binaries and docker image ([#1086](https://github.com/jaegertracing/jaeger/pull/1086), [@ledor473](https://github.com/ledor473))
- Use Go 1.11 ([#1104](https://github.com/jaegertracing/jaeger/pull/1104), [@isaachier](https://github.com/isaachier))
- Tag images with commit SHA and publish to `-snapshot` repository ([#1082](https://github.com/jaegertracing/jaeger/pull/1082), [@pavolloffay](https://github.com/pavolloffay))

##### Bug fixes, Minor Improvements

- Fix child span context while tracing cassandra queries ([#1131](https://github.com/jaegertracing/jaeger/pull/1131), [@annanay25](https://github.com/annanay25))
- Deadlock detector hack for Kafka driver instability ([#1087](https://github.com/jaegertracing/jaeger/pull/1087), [@vprithvi](https://github.com/vprithvi))
- Fix processor overriding data in a buffer ([#1099](https://github.com/jaegertracing/jaeger/pull/1099), [@pavolloffay](https://github.com/pavolloffay))

#### UI Changes

##### New Features

- Span Search - Highlight search results ([#238](https://github.com/jaegertracing/jaeger-ui/pull/238)), [@davit-y](https://github.com/davit-y)
- Span Search - Improve search logic ([#237](https://github.com/jaegertracing/jaeger-ui/pull/237)),  [@davit-y](https://github.com/davit-y)
- Span Search - Add result count, navigation and clear buttons ([#234](https://github.com/jaegertracing/jaeger-ui/pull/234)), [@davit-y](https://github.com/davit-y)

##### Bug Fixes, Minor Improvements

- Use correct duration format for scatter plot ([#266](https://github.com/jaegertracing/jaeger-ui/pull/266)), [@tiffon](https://github.com/tiffon))
- Fix collapse all issues ([#264](https://github.com/jaegertracing/jaeger-ui/pull/264)), [@tiffon](https://github.com/tiffon))
- Use a moderately sized canvas for the span graph ([#257](https://github.com/jaegertracing/jaeger-ui/pull/257)), [@tiffon](https://github.com/tiffon))


1.7.0 (2018-09-19)
------------------

#### UI Changes

- Compare two traces ([#228](https://github.com/jaegertracing/jaeger-ui/pull/228), [@tiffon](https://github.com/tiffon))
- Make tags clickable ([#223](https://github.com/jaegertracing/jaeger-ui/pull/223), [@divdavem](https://github.com/divdavem))
- Directed graph as React component ([#224](https://github.com/jaegertracing/jaeger-ui/pull/224), [@tiffon](https://github.com/tiffon))
- Timeline Expand and Collapse Features ([#221](https://github.com/jaegertracing/jaeger-ui/issues/221), [@davit-y](https://github.com/davit-y))
- Integrate Google Analytics into Search Page ([#220](https://github.com/jaegertracing/jaeger-ui/issues/220), [@davit-y](https://github.com/davit-y))

#### Backend Changes

##### ⛔ Breaking Changes

- `jaeger-standalone` binary has been renamed to `jaeger-all-in-one`. This change also includes package rename from `standalone` to `all-in-one` ([#1062](https://github.com/jaegertracing/jaeger/pull/1062), [@pavolloffay](https://github.com/pavolloffay))

##### New Features

- (Experimental) Allow storing tags as object fields in Elasticsearch for better Kibana support(([#1018](https://github.com/jaegertracing/jaeger/pull/1018), [@pavolloffay](https://github.com/pavolloffay))
- Enable tracing of Cassandra queries ([#1038](https://github.com/jaegertracing/jaeger/pull/1038), [@yurishkuro](https://github.com/yurishkuro))
- Make Elasticsearch index configurable ([#1009](https://github.com/jaegertracing/jaeger/pull/1009), [@pavolloffay](https://github.com/pavoloffay))
- Add flags to allow changing ports for HotROD services ([#951](https://github.com/jaegertracing/jaeger/pull/951), [@cboornaz17](https://github.com/cboornaz17))
- (Experimental) Kafka ingester ([#952](https://github.com/jaegertracing/jaeger/pull/952), [#942](https://github.com/jaegertracing/jaeger/pull/942), [#944](https://github.com/jaegertracing/jaeger/pull/944), [#940](https://github.com/jaegertracing/jaeger/pull/940), [@davit-y](https://github.com/davit-y) and [@vprithvi](https://github.com/vprithvi)))
- Use tags in agent metrics ([#950](https://github.com/jaegertracing/jaeger/pull/950), [@eundoosong](https://github.com/eundoosong))
- Add support for Cassandra reconnect interval ([#934](https://github.com/jaegertracing/jaeger/pull/934), [@nyanshak](https://github.com/nyanshak))

1.6.0 (2018-07-10)
------------------

#### Backend Changes

##### ⛔ Breaking Changes

- The storage implementations no longer write the parentSpanID field to storage (#856).
  If you are upgrading to this version, **you must upgrade query service first**!

- Update Dockerfiles to reference executable via ENTRYPOINT (#815) by Zachary DiCesare (@zdicesare)

  It is no longer necessary to specify the binary name when passing flags to containers.
  For example, to execute the `help` command of the collector, instead of
  ```
  $ docker run -it --rm jaegertracing/jaeger-collector /go/bin/collector-linux help
  ```
  run
  ```
  $ docker run -it --rm jaegertracing/jaeger-collector help
  ```

- Detect HTTP payload format from Content-Type (#916) by Yuri Shkuro (@yurishkuro)

  When submitting spans in Thrift format to HTTP endpoint `/api/traces`,
  the `format` argument is no longer required, but the Content-Type header
  must be set to "application/vnd.apache.thrift.binary".

- Change metric tag from "service" to "svc" (#883) by Won Jun Jang (@black-adder)

##### New Features

- Add Kafka as a Storage Plugin (#862) by David Yeghshatyan (@davit-y)

  The collectors can be configured to write spans to Kafka for further data mining.

- Package static assets inside the query-service binary (#918) by Yuri Shkuro (@yurishkuro)

  It is no longer necessary (but still possible) to pass the path to UI static assets
  to jaeger-query and jaeger-standalone binaries.

- Replace domain model with Protobuf/gogo-generated model (#856) by Yuri Shkuro (@yurishkuro)

  First step towards switching to Protobuf and gRPC.

- Include HotROD binary in the distributions (#917) by Yuri Shkuro (@yurishkuro)
- Improve HotROD demo (#915) by Yuri Shkuro (@yurishkuro)
- Add DisableAutoDiscovery param to cassandra config (#912) by Bill Westlin (@whistlinwilly)
- Add connCheckTimeout flag to agent (#911) by Henrique Rodrigues (@Henrod)
- Ability to use multiple storage types (#880) by David Yeghshatyan (@davit-y)

##### Minor Improvements

- [ES storage] Log number of total and failed requests (#902) by Tomasz Adamski (@tmszdmsk)
- [ES storage] Do not log requests on error (#901) by Tomasz Adamski (@tmszdmsk)
- [ES storage] Do not exceed ES _id length limit (#905) by Łukasz Harasimowicz (@harnash) and Tomasz Adamski (@tmszdmsk)
- Add cassandra index filter (#876) by Won Jun Jang (@black-adder)
- Close span writer in standalone (#863) (4 weeks ago) by Pavol Loffay (@pavolloffay)
- Log configuration options for memory storage (#852) (6 weeks ago) by Juraci Paixão Kröhling (@jpkrohling)
- Update collector metric counters to have a name (#886) by Won Jun Jang (@black-adder)
- Add CONTRIBUTING_GUIDELINES.md (#864) by (@PikBot)

1.5.0 (2018-05-28)
------------------

#### Backend Changes

- Add bounds to memory storage (#845) by Juraci Paixão Kröhling (@jpkrohling)
- Add metric for debug traces (#796) by Won Jun Jang (@black-adder)
- Change metrics naming scheme (#776) by Juraci Paixão Kröhling (@jpkrohling)
- Remove ParentSpanID from domain model (#831) by Yuri Shkuro (@yurishkuro)
- Add ability to adjust static sampling probabilities per operation (#827) by Won Jun Jang (@black-adder)
- Support log-level flag on agent (#828) by Won Jun Jang (@black-adder)
- Add healthcheck to standalone (#784) by Eundoo Song (@eundoosong)
- Do not use KeyValue fields directly and use KeyValues as decorator only (#810) by Yuri Shkuro (@yurishkuro)
- Upgrade to go 1.10 (#792) by Prithvi Raj (@vprithvi)
- Do not create Cassandra index if it already exists (#782) by Greg Swift (@gregswift)

#### UI Changes

- None

1.4.1 (2018-04-21)
------------------

#### Backend Changes

- Publish binaries for Linux, Darwin, and Windows (#765) - thanks to @grounded042

#### UI Changes

##### New Features

- View Trace JSON buttons return formatted JSON (fixes [#199](https://github.com/jaegertracing/jaeger-ui/issues/199))


1.4.0 (2018-04-20)
------------------

#### Backend Changes

##### New Features

- Support traces with >10k spans in Elasticsearch (#668) - thanks to @sramakr

##### Bug Fixes, Minor Improvements

- Allow slash '/' in service names (#586)
- Log errors from HotROD services (#769)


1.3.0 (2018-03-26)
------------------

#### Backend Changes

##### New Features

- Add sampling handler with file-based configuration for agents to query (#720) (#674) <Won Jun Jang>
- Allow overriding base path for UI/API routes and remove --query.prefix (#748) <Yuri Shkuro>
- Add Dockerfile for hotrod example app (#694) <Guilherme Baufaker Rêgo>
- Publish hotrod image to docker hub (#702) <Pavol Loffay>
- Dockerize es-index-cleaner script (#741) <Pavol Loffay>
- Add a flag to control Cassandra consistency level (#700) <Yuri Shkuro>
- Collect metrics from ES bulk service (#688) <Pavol Loffay>
- Allow zero replicas for Elasticsearch (#754) <bharat-p>

##### Bug Fixes, Minor Improvements

- Apply namespace when creating Prometheus metrics factory (fix for #732) (#733) <Yuri Shkuro>
- Disable double compression on Prom Handler - fixes #697 (#735) <Juraci Paixão Kröhling>
- Use the default metricsFactory if not provided (#739) <Louis-Etienne>
- Avoid duplicate expvar metrics - fixes #716 (#726) <Yuri Shkuro>
- Make sure different tracers in HotROD process use different random generator seeds (#718) <Yuri Shkuro>
- Test that processes with identical tags are deduped (#708) <Yuri Shkuro>
- When converting microseconds to time.Time ensure UTC timezone (#712) <Prithvi Raj>
- Add to WaitGroup before the goroutine creation (#711) <Cruth kvinc>
- Pin testify version to ^1.2.1 (#710) <Pavol Loffay>

#### UI Changes

##### New Features

- Support running Jaeger behind a reverse proxy (fixes [#42](https://github.com/jaegertracing/jaeger-ui/issues/42))
- Track Javascript errors via Google Analytics (fixes [#39](https://github.com/jaegertracing/jaeger-ui/issues/39))
- Add Google Analytics event tracking for actions in trace view ([#191](https://github.com/jaegertracing/jaeger-ui/issues/191))

##### Bug Fixes, Minor Improvements

- Clearly identify traces without a root span (fixes [#190](https://github.com/jaegertracing/jaeger-ui/issues/190))
- Fix [#166](https://github.com/jaegertracing/jaeger-ui/issues/166) JS error on search page after viewing 404 trace

#### Documentation Changes


1.2.0 (2018-02-07)
------------------

#### Backend Changes

##### New Features

- Use elasticsearch bulk API (#656) <Pavol Loffay>
- Support archive storage in the query-service (#604) <Yuri Shkuro>
- Introduce storage factory framework and composable CLI (#625) <Yuri Shkuro>
- Make agent host port configurable in hotrod (#663) <Pavol Loffay>
- Add signal handling to standalone (#657) <Pavol Loffay>

##### Bug Fixes, Minor Improvements

- Remove the override of GOMAXPROCS (#679) <Cruth kvinc>
- Use UTC timezone for ES indices (#646) <Pavol Loffay>
- Fix elasticsearch create index race condition error (#641) <Pavol Loffay>

#### UI Changes

##### New Features

- Use Ant Design instead of Semantic UI (https://github.com/jaegertracing/jaeger-ui/pull/169)
  - Fix [#164](https://github.com/jaegertracing/jaeger-ui/issues/164) - Use Ant Design instead of Semantic UI
  - Fix [#165](https://github.com/jaegertracing/jaeger-ui/issues/165) - Search results are shown without a date
  - Fix [#69](https://github.com/jaegertracing/jaeger-ui/issues/69) - Missing endpoints in jaeger ui dropdown

##### Bug Fixes, Minor Improvements

- Fix 2 digit lookback (12h, 24h) parsing (https://github.com/jaegertracing/jaeger-ui/issues/167)


1.1.0 (2018-01-03)
------------------

#### Backend Changes

##### New Features

- Add support for retrieving unadjusted/raw traces (#615)
- Add CA certificates to collector/query images (#485)
- Parse zipkin v2 high trace id (#596)

##### Bug Fixes, Minor Improvements

- Skip nil and zero length hits in ElasticSearch storage (#601)
- Make Cassandra service_name_index inserts idempotent (#587)
- Align atomic int64 to word boundary to fix SIGSEGV (#592)
- Add adjuster that removes bad span references (#614)
- Set operationNames cache initial capacity to 10000 (#621)

#### UI Changes

##### New Features

- Change tag search input syntax to logfmt (https://github.com/jaegertracing/jaeger-ui/issues/145)
- Make threshold for enabling DAG view configurable (https://github.com/jaegertracing/jaeger-ui/issues/130)
- Show better error messages for failed API calls (https://github.com/jaegertracing/jaeger-ui/issues/127)
- Add View Option for raw/unadjusted trace (https://github.com/jaegertracing/jaeger-ui/issues/153)
- Add timezone tooltip to custom lookback form-field (https://github.com/jaegertracing/jaeger-ui/pull/161)

##### Bug Fixes, Minor Improvements

- Use consistent icons for logs expanded/collapsed (https://github.com/jaegertracing/jaeger-ui/issues/86)
- Encode service name in API calls to allow '/' (https://github.com/jaegertracing/jaeger-ui/issues/138)
- Fix endless trace HTTP requests (https://github.com/jaegertracing/jaeger-ui/issues/128)
- Fix JSON view when running in dev mode (https://github.com/jaegertracing/jaeger-ui/issues/139)
- Fix trace name resolution (https://github.com/jaegertracing/jaeger-ui/pull/134)
- Only JSON.parse JSON strings in tags/logs values (https://github.com/jaegertracing/jaeger-ui/pull/162)


1.0.0 (2017-12-04)
------------------

#### Backend Changes

- Support Prometheus metrics as default for all components (#516)
- Enable TLS client connections to Cassandra (#555)
- Fix issue where Domain to UI model converter double reports references (#579)

#### UI Changes

- Make dependencies tab configurable (#122)


0.10.0 (2017-11-17)
------------------

#### UI Changes

- Verify stored search settings [jaegertracing/jaeger-ui#111](https://github.com/jaegertracing/jaeger-ui/pull/111)
- Fix browser back button not working correctly [jaegertracing/jaeger-ui#110](https://github.com/jaegertracing/jaeger-ui/pull/110)
- Handle FOLLOWS_FROM ref type [jaegertracing/jaeger-ui#118](https://github.com/jaegertracing/jaeger-ui/pull/118)

#### Backend Changes

- Allow embedding custom UI config in index.html [#490](https://github.com/jaegertracing/jaeger/pull/490)
- Add startTimeMillis field to JSON Spans submitted to ElasticSearch [#491](https://github.com/jaegertracing/jaeger/pull/491)
- Introduce version command and handler [#517](https://github.com/jaegertracing/jaeger/pull/517)
- Fix ElasticSearch aggregation errors when index is empty [#535](https://github.com/jaegertracing/jaeger/pull/535)
- Change package from uber to jaegertracing [#528](https://github.com/jaegertracing/jaeger/pull/528)
- Introduce logging level configuration [#514](https://github.com/jaegertracing/jaeger/pull/514)
- Support Zipkin v2 json [#518](https://github.com/jaegertracing/jaeger/pull/518)
- Add HTTP compression handler [#545](https://github.com/jaegertracing/jaeger/pull/545)


0.9.0 (2017-10-25)
------------------

#### UI Changes

- Refactor trace detail [jaegertracing/jaeger-ui#53](https://github.com/jaegertracing/jaeger-ui/pull/53)
- Virtualized scrolling for trace detail view [jaegertracing/jaeger-ui#68](https://github.com/jaegertracing/jaeger-ui/pull/68)
- Mouseover expands truncated text to full length in left column in trace view [jaegertracing/jaeger-ui#71](https://github.com/jaegertracing/jaeger-ui/pull/71)
- Make left column adjustable in trace detail view [jaegertracing/jaeger-ui#74](https://github.com/jaegertracing/jaeger-ui/pull/74)
- Fix trace mini-map blurriness when < 60 spans [jaegertracing/jaeger-ui#77](https://github.com/jaegertracing/jaeger-ui/pull/77)
- Fix Google Analytics tracking [jaegertracing/jaeger-ui#81](https://github.com/jaegertracing/jaeger-ui/pull/81)
- Improve search dropdowns [jaegertracing/jaeger-ui#84](https://github.com/jaegertracing/jaeger-ui/pull/84)
- Add keyboard shortcuts and minimap UX [jaegertracing/jaeger-ui#93](https://github.com/jaegertracing/jaeger-ui/pull/93)

#### Backend Changes

- Add tracing to the query server [#454](https://github.com/uber/jaeger/pull/454)
- Support configuration files [#462](https://github.com/uber/jaeger/pull/462)
- Add cassandra tag filter [#442](https://github.com/uber/jaeger/pull/442)
- Handle ports > 32k in Zipkin JSON [#488](https://github.com/uber/jaeger/pull/488)


0.8.0 (2017-09-24)
------------------

- Convert to Apache 2.0 License


0.7.0 (2017-08-22)
------------------

- Add health check server to collector and query [#280](https://github.com/uber/jaeger/pull/280)
- Add/fix sanitizer for Zipkin span start time and duration [#333](https://github.com/uber/jaeger/pull/333)
- Support Zipkin json encoding for /api/v1/spans HTTP endpoint [#348](https://github.com/uber/jaeger/pull/348)
- Support Zipkin 128bit traceId and ipv6 [#349](https://github.com/uber/jaeger/pull/349)


0.6.0 (2017-08-09)
------------------

- Add viper/cobra configuration support [#245](https://github.com/uber/jaeger/pull/245) [#307](https://github.com/uber/jaeger/pull/307)
- Add Zipkin /api/v1/spans endpoint [#282](https://github.com/uber/jaeger/pull/282)
- Add basic authenticator to configs for cass
