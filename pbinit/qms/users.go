package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewUserResponse() *qms.QMSUserResponse {
	return &qms.QMSUserResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewUserList() *qms.QMSUserList {
	return &qms.QMSUserList{
		Header: gotelnats.NewHeader(),
		Users:  make([]*qms.QMSUser, 0),
	}
}

func NewAddUserRequest() *qms.AddUserRequest {
	return &qms.AddUserRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewAddUserResponse() *qms.AddUserResponse {
	return &qms.AddUserResponse{
		Header: gotelnats.NewHeader(),
	}
}

// Initializers

func InitQMSAddUserRequest(request *qms.AddUserRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitQMSAddUserResponse(request *qms.AddUserResponse, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
