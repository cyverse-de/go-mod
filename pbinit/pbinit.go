package pbinit

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/analysis"
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

func NewAnalysisStatus() *analysis.AnalysisStatus {
	return &analysis.AnalysisStatus{
		Header: gotelnats.NewHeader(),
	}
}

func InitAnalysisStatus(request *analysis.AnalysisStatus, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
