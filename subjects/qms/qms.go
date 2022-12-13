package qms

import "fmt"

const qmsUser = "cyverse.qms.user"
const qmsPlan = "cyverse.qms.plan"

var (
	GetUserUpdates = fmt.Sprintf("%s.updates.get", qmsUser)
	AddUserUpdate  = fmt.Sprintf("%s.updates.add", qmsUser)

	GetUserOverages   = fmt.Sprintf("%s.overages.get", qmsUser)
	CheckUserOverages = fmt.Sprintf("%s.overages.check", qmsUser)

	GetUserUsages = fmt.Sprintf("%s.usages.get", qmsUser)
	AddUserUsages = fmt.Sprintf("%s.usages.add", qmsUser)

	UserSummary = fmt.Sprintf("%s.summary.get", qmsUser)
	AddUser     = fmt.Sprintf("%s.add", qmsUser)

	GetUserPlan    = fmt.Sprintf("%s.plan.get", qmsUser)
	ChangeUserPlan = fmt.Sprintf("%s.plan.change", qmsUser)

	AddQuota = fmt.Sprintf("%s.quota.add", qmsUser)

	ListPlans           = fmt.Sprintf("%s.list", qmsPlan)
	UpdatePlan          = fmt.Sprintf("%s.update", qmsPlan)
	AddPlan             = fmt.Sprintf("%s.add", qmsPlan)
	GetPlan             = fmt.Sprintf("%s.get", qmsPlan)
	UpsertQuotaDefaults = fmt.Sprintf("%s.quota.defaults", qmsPlan)
)
