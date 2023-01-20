package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

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

func NewOverageResponse() *qms.OverageResponse {
	return &qms.OverageResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewAllUserOveragesRequest(username string) *qms.AllUserOveragesRequest {
	return &qms.AllUserOveragesRequest{
		Header:   gotelnats.NewHeader(),
		Username: username,
	}
}

func NewUserResourceOveragesRequest(username, resourceName string) *qms.UserResourceOveragesRequest {
	return &qms.UserResourceOveragesRequest{
		Header:       gotelnats.NewHeader(),
		Username:     username,
		ResourceName: resourceName,
	}
}

func NewIsOverageRequest(username, resourceName string) *qms.IsOverageRequest {
	return &qms.IsOverageRequest{
		Header:       gotelnats.NewHeader(),
		Username:     username,
		ResourceName: resourceName,
	}
}

// Initalizers

func InitIsOverageRequest(request *qms.IsOverageRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAllUserOveragesRequest(request *qms.AllUserOveragesRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
