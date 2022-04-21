# otelutils

Makes for a central spot to stick creating a TracerProvider for OpenTelemetry.

Do:

```go
var serviceName = "some-service-name"
var tracerCtx, cancel = context.WithCancel(context.Background())
defer cancel()
shutdown := otelutils.TracerProviderFromEnv(tracerCtx, serviceName, func(e error) { log.Fatal(e) })
defer shutdown()
```

`shutdown` will try to shut down the created TracerProvider with a short timeout, so as to not hang shutting down the service.
