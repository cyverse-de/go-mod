package qms

import (
	"fmt"
)

const qmsUser = "cyverse.qms.user"
const qmsPlan = "cyverse.qms.plan"
const qmsAddon = "cyverse.qms.addon"

var qmsSubAddon = fmt.Sprintf("%s.plan.addons", qmsUser)

var (
	GetUserUpdates = fmt.Sprintf("%s.updates.get", qmsUser)
	AddUserUpdate  = fmt.Sprintf("%s.updates.add", qmsUser)

	GetUserOverages   = fmt.Sprintf("%s.overages.get", qmsUser)
	CheckUserOverages = fmt.Sprintf("%s.overages.check", qmsUser)

	GetUserUsages = fmt.Sprintf("%s.usages.get", qmsUser)
	AddUserUsages = fmt.Sprintf("%s.usages.add", qmsUser)

	UserSummary = fmt.Sprintf("%s.summary.get", qmsUser)
	AddUser     = fmt.Sprintf("%s.add", qmsUser)

	GetSubscription         = fmt.Sprintf("%s.plan.get", qmsUser)
	ChangeSubscription      = fmt.Sprintf("%s.plan.change", qmsUser)
	ListSubscriptionAddons  = fmt.Sprintf("%s.list", qmsSubAddon)
	AddSubscriptionAddon    = fmt.Sprintf("%s.add", qmsSubAddon)
	UpdateSubscriptionAddon = fmt.Sprintf("%s.update", qmsSubAddon)
	DeleteSubscriptionAddon = fmt.Sprintf("%s.delete", qmsSubAddon)
	GetSubscriptionAddon    = fmt.Sprintf("%s.get", qmsSubAddon)

	AddQuota = fmt.Sprintf("%s.quota.add", qmsUser)

	ListPlans           = fmt.Sprintf("%s.list", qmsPlan)
	UpdatePlan          = fmt.Sprintf("%s.update", qmsPlan)
	AddPlan             = fmt.Sprintf("%s.add", qmsPlan)
	GetPlan             = fmt.Sprintf("%s.get", qmsPlan)
	UpsertQuotaDefaults = fmt.Sprintf("%s.quota.defaults", qmsPlan)

	AddAddon        = fmt.Sprintf("%s.add", qmsAddon)
	ListAddons      = fmt.Sprintf("%s.list", qmsAddon)
	UpdateAddon     = fmt.Sprintf("%s.update", qmsAddon)
	DeleteAddon     = fmt.Sprintf("%s.delete", qmsAddon)
	ToggleAddonPaid = fmt.Sprintf("%s.togglepaid", qmsAddon)
)
