package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewUpdateListRequest() *qms.UpdateListRequest {
	return &qms.UpdateListRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewUpdateListResponse() *qms.UpdateListResponse {
	return &qms.UpdateListResponse{
		Header:  gotelnats.NewHeader(),
		Updates: make([]*qms.Update, 0),
	}
}

func NewAddUpdateRequest(update *qms.Update) *qms.AddUpdateRequest {
	return &qms.AddUpdateRequest{
		Header: gotelnats.NewHeader(),
		Update: update,
	}
}

func NewAddUpdateResponse() *qms.AddUpdateResponse {
	return &qms.AddUpdateResponse{
		Header: gotelnats.NewHeader(),
		Update: &qms.Update{},
	}
}

// Initializers

func InitUpdateListRequest(request *qms.UpdateListRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAddUpdateRequest(request *qms.AddUpdateRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
