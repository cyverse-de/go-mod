package pbinit

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/header"
	"github.com/cyverse-de/p/go/monitoring"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

// The initializing code has an annoying amount of duplication because of an
// intersection of missing features in protocol buffers for Go and the Go
// generics implementation. Go generics don't currently support structural types
// (using the presence and type of a field in a struct as part of a type
// constraint) while the protocol buffers implementation for Go doesn't generate
// setter functions for fields, which could be used in a type constraint
// interface. That combo means that we can't have a generic function that
// initializes types generated from protocol bufffers since we can't define a
// constraining interface requiring a field (not supported by Go) or a setter
// function (not supported by protocol buffers).
//
// Hopefully one or the other will get updated at some point and this code can
// get collapsed down to something less annoying to maintain.

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

func NewUserPlanResponse() *qms.UserPlanResponse {
	return &qms.UserPlanResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewUserPlanList() *qms.UserPlanList {
	return &qms.UserPlanList{
		Header:    gotelnats.NewHeader(),
		UserPlans: make([]*qms.UserPlan, 0),
	}
}

func NewQMSUserResponse() *qms.QMSUserResponse {
	return &qms.QMSUserResponse{
		Header: gotelnats.NewHeader(),
	}
}

func NewQMSUserList() *qms.QMSUserList {
	return &qms.QMSUserList{
		Header: gotelnats.NewHeader(),
		Users:  make([]*qms.QMSUser, 0),
	}
}

func NewMonitoringHeartbeat() *monitoring.Heartbeat {
	return &monitoring.Heartbeat{
		Header: gotelnats.NewHeader(),
	}
}

func NewDNSCheckResult() *monitoring.DNSCheckResult {
	return &monitoring.DNSCheckResult{
		Header:  gotelnats.NewHeader(),
		Lookups: make([]*monitoring.DNSLookup, 0),
	}
}

// Initialize requests from here on down.
func commonInit(h *header.Header, subject string) (context.Context, trace.Span) {
	carrier := gotelnats.PBTextMapCarrier{
		Header: h,
	}

	return gotelnats.StartSpan(&carrier, subject, gotelnats.Process)
}

func InitAllUserOveragesRequest(request *qms.AllUserOveragesRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}

func InitIsOverageRequest(request *qms.IsOverageRequest, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}

func InitAddUsage(request *qms.AddUsage, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}

func InitGetUsages(request *qms.GetUsages, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return commonInit(request.Header, subject)
}
