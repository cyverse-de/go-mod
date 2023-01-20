package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewAddUsage(username, resourceName, updateType string, usageValue float64) *qms.AddUsage {
	return &qms.AddUsage{
		Header:       gotelnats.NewHeader(),
		Username:     username,
		ResourceName: resourceName,
		UpdateType:   updateType,
		UsageValue:   usageValue,
	}
}

func NewGetUsages(username string) *qms.GetUsages {
	return &qms.GetUsages{
		Header:   gotelnats.NewHeader(),
		Username: username,
	}
}

func NewUsageResponse() *qms.UsageResponse {
	return &qms.UsageResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewUsageList() *qms.UsageList {
	return &qms.UsageList{
		Header: gotelnats.NewHeader(),
		Usages: make([]*qms.Usage, 0),
	}
}

// Initializers

func InitAddUsage(request *qms.AddUsage, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitGetUsages(request *qms.GetUsages, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
