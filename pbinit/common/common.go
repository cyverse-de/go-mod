package common

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/header"
	"go.opentelemetry.io/otel/trace"
)

// Init initialize messages that go over NATS.
func Init(h *header.Header, subject string) (context.Context, trace.Span) {
	carrier := gotelnats.PBTextMapCarrier{
		Header: h,
	}

	return gotelnats.StartSpan(&carrier, subject, gotelnats.Process)
}
