package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewAddAddonRequest() *qms.AddAddonRequest {
	return &qms.AddAddonRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewAddonResponse() *qms.AddonResponse {
	return &qms.AddonResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewAddonListResponse() *qms.AddonListResponse {
	return &qms.AddonListResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewUpdateAddonRequest() *qms.UpdateAddonRequest {
	return &qms.UpdateAddonRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewUpdateSubscriptionAddonRequest() *qms.UpdateSubscriptionAddonRequest {
	return &qms.UpdateSubscriptionAddonRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewSubscriptionAddonResponse() *qms.SubscriptionAddonResponse {
	return &qms.SubscriptionAddonResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewSubscriptionAddonListResponse() *qms.SubscriptionAddonListResponse {
	return &qms.SubscriptionAddonListResponse{
		Header: gotelnats.NewHeader(),
	}
}

type addonLookupRequestParams struct {
	usesID   bool
	id       string
	usesName bool
	name     string
}

type AddonLookupRequestOption func(*addonLookupRequestParams)

func WithALRByID(id string) AddonLookupRequestOption {
	return func(a *addonLookupRequestParams) {
		a.usesID = true
		a.id = id
	}
}

func WithALRByName(name string) AddonLookupRequestOption {
	return func(a *addonLookupRequestParams) {
		a.usesName = true
		a.name = name
	}
}

func NewAddonLookupRequest(option AddonLookupRequestOption) *qms.AddonLookupRequest {
	settings := &addonLookupRequestParams{}
	option(settings)

	retval := &qms.AddonLookupRequest{
		Header: gotelnats.NewHeader(),
	}

	if settings.usesName {
		retval.Addon = &qms.AddonLookupRequest_Name{
			Name: settings.name,
		}
	} else {
		retval.Addon = &qms.AddonLookupRequest_Uuid{
			Uuid: settings.id,
		}
	}

	return retval
}

// Initializers

func InitAddonLookupRequest(request *qms.AddonLookupRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAddAddonRequest(request *qms.AddAddonRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitUpdateAddonRequest(request *qms.UpdateAddonRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitUpdateSubscriptionAddonRequest(request *qms.UpdateSubscriptionAddonRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
