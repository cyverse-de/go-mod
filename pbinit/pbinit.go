package pbinit

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/header"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

// The initializing code has an annoying amount of duplication because of an
// intersection of missing features in protocol buffers for Go and the Go
// generics implementation. Go generics don't currently support structural types
// (using the presence and type of a field in a struct as part of a type
// constraint) while the protocol buffers implementation for Go doesn't generate
// setter functions for fields, which could be used in a type constraint
// interface. That combo means that we can't have a generic function that
// initializes types generated from protocol bufffers since we can't define a
// constraining interface requiring a field (not supported by Go) or a setter
// function (not supported by protocol buffers).
//
// Hopefully one or the other will get updated at some point and this code can
// get collapsed down to something less annoying to maintain.

func NewOverageList() *qms.OverageList {
	return &qms.OverageList{
		Header:   gotelnats.NewHeader(),
		Overages: make([]*qms.Overage, 0),
	}
}

func NewIsOverage() *qms.IsOverage {
	return &qms.IsOverage{
		Header:    gotelnats.NewHeader(),
		IsOverage: false,
	}
}

func commonInit(h *header.Header, subject string) (context.Context, trace.Span) {
	carrier := gotelnats.PBTextMapCarrier{
		Header: h,
	}

	return gotelnats.StartSpan(&carrier, subject, gotelnats.Process)
}

func InitAllUserOveragesRequest(request *qms.AllUserOveragesRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}

func InitIsOverageRequest(request *qms.IsOverageRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}
