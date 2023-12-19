ctl.addr=localhost:50052


## Obs configs
obs.enabled=true
obs.exporter.trace.grpc.host=localhost:4317
obs.exporter.metrics.http.host=localhost:9090
obs.exporter.metrics.host.path=/api/v1/otlp/v1/metrics
obs.metrics.host=true
obs.metrics.runtime=true

log.level=0
log.console.enabled=false
log.file.enabled=true
log.file.name=ctl.log