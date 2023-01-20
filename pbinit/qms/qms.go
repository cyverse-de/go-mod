package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewRequestByUsername() *qms.RequestByUsername {
	return &qms.RequestByUsername{
		Header: gotelnats.NewHeader(),
	}
}

func NewRequestByUserID() *qms.RequestByUserID {
	return &qms.RequestByUserID{
		Header: gotelnats.NewHeader(),
	}
}

func NewNoParamsRequest() *qms.NoParamsRequest {
	return &qms.NoParamsRequest{
		Header: gotelnats.NewHeader(),
	}
}

// Initializers

func InitRequestByUsername(request *qms.RequestByUsername, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitRequestByUserID(request *qms.RequestByUserID, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitNoParamsRequest(request *qms.NoParamsRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
