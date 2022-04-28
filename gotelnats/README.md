# gotelnats

Adds some utility functions to make adding telemetry info to NATS messages a bit
easier.

Use `otelutils` from `github.com/cyverse-de/go-mod/otelutils` to set up the
default tracer provider based on environment variable configuration.

# Example

Logic to extract span information from a protocol buffer header field.

```go
...
	return func(subject, reply string, request *user.UserLookupRequest) {
		var err error

		carrier := gotelnats.PBTextMapCarrier{
			Header: request.Header,
		}

		ctx, span := gotelnats.StartSpan(&carrier, subject, gotelnats.Process)
		defer span.End()
...
```

Functions that add span info to the protocol buffer header field for outgoing
messages.

```go
func publishResponse(ctx context.Context, conn *nats.EncodedConn, reply string, responseUser *user.User) error {
	carrier := gotelnats.PBTextMapCarrier{
		Header: responseUser.Header,
	}

	_, span := gotelnats.InjectSpan(ctx, &carrier, reply, gotelnats.Send)
	defer span.End()

	return conn.Publish(reply, &responseUser)
}

func handleError(ctx context.Context, err error, code svcerror.Code, reply string, conn *nats.EncodedConn) {
	svcerr := svcerror.Error{
		ErrorCode: code,
		Message:   err.Error(),
	}

	log.Error(&svcerr)

	carrier := gotelnats.PBTextMapCarrier{
		Header: svcerr.Header,
	}

	_, span := gotelnats.InjectSpan(ctx, &carrier, reply, gotelnats.Send)
	defer span.End()

	if err = conn.Publish(reply, &svcerr); err != nil {
		log.Error(err)
	}
}
```
