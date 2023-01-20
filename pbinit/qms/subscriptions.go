package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewSubscriptionResponse() *qms.SubscriptionResponse {
	return &qms.SubscriptionResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewSubscriptionList() *qms.SubscriptionList {
	return &qms.SubscriptionList{
		Header:        gotelnats.NewHeader(),
		Subscriptions: make([]*qms.Subscription, 0),
	}
}

func NewChangeSubscriptionRequest() *qms.ChangeSubscriptionRequest {
	return &qms.ChangeSubscriptionRequest{
		Header: gotelnats.NewHeader(),
	}
}

// Initializers

func InitChangeSubscriptionRequest(request *qms.ChangeSubscriptionRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
