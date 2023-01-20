package qms

import (
	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/qms"
)

func NewAddonRequest() *qms.AddAddonRequest {
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
