/**
 * Do not add new functions to this file unless you are fixing backwards
 * compatibility issues related to this module. New functionality should go into
 * the pbinit/qms module and your application should be updated to use it, instead.
 */

package pbinit

import (
	"context"

	qmsinit "github.com/cyverse-de/go-mod/pbinit/qms"
	"github.com/cyverse-de/p/go/qms"
	"go.opentelemetry.io/otel/trace"
)

func NewOverageList() *qms.OverageList                   { return qmsinit.NewOverageList() }
func NewIsOverage() *qms.IsOverage                       { return qmsinit.NewIsOverage() }
func NewOverageResponse() *qms.OverageResponse           { return qmsinit.NewOverageResponse() }
func NewQuotaDefaultResponse() *qms.QuotaDefaultResponse { return qmsinit.NewQuotaDefaultResponse() }
func NewQuotaDefaultList() *qms.QuotaDefaultList         { return qmsinit.NewQuotaDefaultList() }
func NewPlanResponse() *qms.PlanResponse                 { return qmsinit.NewPlanResponse() }
func NewPlanList() *qms.PlanList                         { return qmsinit.NewPlanList() }
func NewQuotaResponse() *qms.QuotaResponse               { return qmsinit.NewQuotaResponse() }
func NewQuotaList() *qms.QuotaList                       { return qmsinit.NewQuotaList() }
func NewResourceTypeResponse() *qms.ResourceTypeResponse { return qmsinit.NewResourceTypeResponse() }
func NewResourceTypeList() *qms.ResourceTypeList         { return qmsinit.NewResourceTypeList() }
func NewGetUsages(username string) *qms.GetUsages        { return qmsinit.NewGetUsages(username) }
func NewUsageResponse() *qms.UsageResponse               { return qmsinit.NewUsageResponse() }
func NewUsageList() *qms.UsageList                       { return qmsinit.NewUsageList() }
func NewSubscriptionResponse() *qms.SubscriptionResponse { return qmsinit.NewSubscriptionResponse() }
func NewSubscriptionList() *qms.SubscriptionList         { return qmsinit.NewSubscriptionList() }
func NewQMSUserResponse() *qms.QMSUserResponse           { return qmsinit.NewUserResponse() }
func NewQMSUserList() *qms.QMSUserList                   { return qmsinit.NewUserList() }
func NewQMSUpdateListRequest() *qms.UpdateListRequest    { return qmsinit.NewUpdateListRequest() }
func NewQMSUpdateListResponse() *qms.UpdateListResponse  { return qmsinit.NewUpdateListResponse() }
func NewQMSAddUpdateResponse() *qms.AddUpdateResponse    { return qmsinit.NewAddUpdateResponse() }
func NewQMSRequestByUsername() *qms.RequestByUsername    { return qmsinit.NewRequestByUsername() }
func NewQMSRequestByUserID() *qms.RequestByUserID        { return qmsinit.NewRequestByUserID() }
func NewQMSAddUserRequest() *qms.AddUserRequest          { return qmsinit.NewAddUserRequest() }
func NewQMSAddUserResponse() *qms.AddUserResponse        { return qmsinit.NewAddUserResponse() }
func NewQMSNoParamsRequest() *qms.NoParamsRequest        { return qmsinit.NewNoParamsRequest() }
func NewQMSPlanRequest() *qms.PlanRequest                { return qmsinit.NewPlanRequest() }
func NewQMSAddPlanRequest() *qms.AddPlanRequest          { return qmsinit.NewAddPlanRequest() }

func NewAllUserOveragesRequest(username string) *qms.AllUserOveragesRequest {
	return qmsinit.NewAllUserOveragesRequest(username)
}

func NewUserResourceOveragesRequest(username, resourceName string) *qms.UserResourceOveragesRequest {
	return qmsinit.NewUserResourceOveragesRequest(username, resourceName)
}

func NewIsOverageRequest(username, resourceName string) *qms.IsOverageRequest {
	return qmsinit.NewIsOverageRequest(username, resourceName)
}

func NewAddUsage(username, resourceName, updateType string, usageValue float64) *qms.AddUsage {
	return qmsinit.NewAddUsage(username, resourceName, updateType, usageValue)
}

func NewAddUpdateRequest(update *qms.Update) *qms.AddUpdateRequest {
	return qmsinit.NewAddUpdateRequest(update)
}

func NewChangeSubscriptionRequest() *qms.ChangeSubscriptionRequest {
	return qmsinit.NewChangeSubscriptionRequest()
}

func NewQMSAddPlanQuotaDefaultRequest() *qms.AddPlanQuotaDefaultRequest {
	return qmsinit.NewQMSAddPlanQuotaDefaultRequest()
}

// Initializers

func InitAllUserOveragesRequest(request *qms.AllUserOveragesRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAllUserOveragesRequest(request, subject)
}

func InitIsOverageRequest(request *qms.IsOverageRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitIsOverageRequest(request, subject)
}

func InitAddUsage(request *qms.AddUsage, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAddUsage(request, subject)
}

func InitGetUsages(request *qms.GetUsages, subject string) (context.Context, trace.Span) {
	return qmsinit.InitGetUsages(request, subject)
}

func InitQMSUpdateListRequest(request *qms.UpdateListRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitUpdateListRequest(request, subject)
}

func InitQMSRequestByUsername(request *qms.RequestByUsername, subject string) (context.Context, trace.Span) {
	return qmsinit.InitRequestByUsername(request, subject)
}

func InitQMSRequestByUserID(request *qms.RequestByUserID, subject string) (context.Context, trace.Span) {
	return qmsinit.InitRequestByUserID(request, subject)
}

func InitQMSNoParamsRequest(request *qms.NoParamsRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitNoParamsRequest(request, subject)
}

func InitQMSAddUpdateRequest(request *qms.AddUpdateRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAddUpdateRequest(request, subject)
}

func InitChangeSubscriptionRequest(request *qms.ChangeSubscriptionRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitChangeSubscriptionRequest(request, subject)
}

func InitQMSAddUserRequest(request *qms.AddUserRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitQMSAddUserRequest(request, subject)
}

func InitQMSAddUserResponse(request *qms.AddUserResponse, subject string) (context.Context, trace.Span) {
	return qmsinit.InitQMSAddUserResponse(request, subject)
}

func InitQMSPlanRequest(request *qms.PlanRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitPlanRequest(request, subject)
}

func InitQMSAddPlanRequest(request *qms.AddPlanRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAddPlanRequest(request, subject)
}

func InitQMSAddQuotaRequest(request *qms.AddQuotaRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAddQuotaRequest(request, subject)
}

func InitQMSAddPlanQuotaDefaultRequest(request *qms.AddPlanQuotaDefaultRequest, subject string) (context.Context, trace.Span) {
	return qmsinit.InitAddPlanQuotaDefaultRequest(request, subject)
}
