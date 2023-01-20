package qms

import (
	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/qms"
)

func NewResourceTypeResponse() *qms.ResourceTypeResponse {
	return &qms.ResourceTypeResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewResourceTypeList() *qms.ResourceTypeList {
	return &qms.ResourceTypeList{
		Header:        gotelnats.NewHeader(),
		ResourceTypes: make([]*qms.ResourceType, 0),
	}
}
