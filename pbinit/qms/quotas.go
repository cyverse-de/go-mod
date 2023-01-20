package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewQMSAddPlanQuotaDefaultRequest() *qms.AddPlanQuotaDefaultRequest {
	return &qms.AddPlanQuotaDefaultRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewQuotaDefaultResponse() *qms.QuotaDefaultResponse {
	return &qms.QuotaDefaultResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewQuotaDefaultList() *qms.QuotaDefaultList {
	return &qms.QuotaDefaultList{
		Header:        gotelnats.NewHeader(),
		QuotaDefaults: make([]*qms.QuotaDefault, 0),
	}
}

func NewQuotaResponse() *qms.QuotaResponse {
	return &qms.QuotaResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewQuotaList() *qms.QuotaList {
	return &qms.QuotaList{
		Header: gotelnats.NewHeader(),
		Quotas: make([]*qms.Quota, 0),
	}
}

// Initializers

func InitAddQuotaRequest(request *qms.AddQuotaRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAddPlanQuotaDefaultRequest(request *qms.AddPlanQuotaDefaultRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
