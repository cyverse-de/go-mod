# gotelnats

Adds some utility functions to make adding telemetry info to NATS messages a bit
easier.

Use `otelutils` from `github.com/cyverse-de/go-mod/otelutils` to set up the
default tracer provider based on environment variable configuration.

# Example

Request information from another service using a request type defined as a
protocol buffer in [github.com/cyverse-de/p](https://github.com/cyverse-de/p)
and extract the response type. Ensures that errors are returned if the
responding service sends one back.

Taken from
[github.com/cyverse-de/discoenv-analyses](https://github.com/cyverse-de/discoenv-analyses).

```go
func lookupUsername(ctx context.Context, conn *nats.EncodedConn, subject string, userID string) (string, error) {
	var (
		err      error
		request  *pbuser.UserLookupRequest
		expected *pbuser.User
	)

	request = &pbuser.UserLookupRequest{
		LookupIds: &pbuser.UserLookupRequest_UserId{
			UserId: userID,
		},
	}

	expected = &pbuser.User{}

	if err = gotelnats.Request[*pbuser.UserLookupRequest, *pbuser.User](ctx, conn, subject, request, expected); err != nil {
		return "", err
	}

	return expected.Username, nil
```

Respond to a request with a type defined as a protocol buffer in
[github.com/cyverse-de/p](https://github.com/cyverse-de/p). Handles adding
telemetry information to the response from the context:

```go
	if err = gotelnats.PublishResponse(ctx, conn, reply, responseUser); err != nil {
		log.Error(err)
	}
```

Set error data into an outgoing response and send it:

```go
	responseUser.Error = gotelnats.InitServiceError(ctx, err, &gotelnats.ErrorOptions{
		ErrorCode: svcerror.ErrorCode_INTERNAL,
	})
	if err = gotelnats.PublishResponse(ctx, conn, reply, responseUser); err != nil {
		log.Error(err)
	}
	return
```

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
