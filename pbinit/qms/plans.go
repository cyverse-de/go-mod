package qms

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewPlanResponse() *qms.PlanResponse {
	return &qms.PlanResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewPlanList() *qms.PlanList {
	return &qms.PlanList{
		Header: gotelnats.NewHeader(),
		Plans:  make([]*qms.Plan, 0),
	}
}

func NewPlanRequest() *qms.PlanRequest {
	return &qms.PlanRequest{
		Header: gotelnats.NewHeader(),
	}
}

func NewAddPlanRequest() *qms.AddPlanRequest {
	return &qms.AddPlanRequest{
		Header: gotelnats.NewHeader(),
	}
}

// Initializers

func InitPlanRequest(request *qms.PlanRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAddPlanRequest(request *qms.AddPlanRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
