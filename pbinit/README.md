# pbinit

This package contains constructors for the types defined in
github.com/cyverse-de/p/go. This prevents users of the package from accidentally
forgetting to initialize header types and allows them to initialize requests
with telemetry support in a bit less code.

## Examples

Sets up a response for a request that checks if a user is in overage for a
resource, along with initializing a the request to get telemetry objects.

```go
	response := pbinit.NewIsOverage()
	ctx, span := pbinit.InitIsOverageRequest(request, subject)
	defer span.End()
```
