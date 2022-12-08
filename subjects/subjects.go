package subjects

import "fmt"

const qmsUser = "cyverse.qms.user"
const qmsPlan = "cyverse.qms.plan"
const cyverseAnalysis = "cyverse.analysis"
const cyverseMonitoring = "cyverse.monitoring"

var (
	QMSGetUserUpdates           = fmt.Sprintf("%s.updates.get", qmsUser)
	SubscriptionsGetUserUpdates = QMSGetUserUpdates

	QMSAddUserUpdate     = fmt.Sprintf("%s.updates.add", qmsUser)
	QMSGetUserOverages   = fmt.Sprintf("%s.overages.get", qmsUser)
	QMSCheckUserOverages = fmt.Sprintf("%s.overages.check", qmsUser)
	QMSGetUserUsages     = fmt.Sprintf("%s.usages.get", qmsUser)
	QMSAddUserUsages     = fmt.Sprintf("%s.usages.add", qmsUser)
	QMSUserSummary       = fmt.Sprintf("%s.summary.get", qmsUser)
	QMSAddUser           = fmt.Sprintf("%s.add", qmsUser)
	QMSGetUserPlan       = fmt.Sprintf("%s.plan.get", qmsUser)
	QMSChangeUserPlan    = fmt.Sprintf("%s.plan.change", qmsUser)
	QMSListPlans         = fmt.Sprintf("%s.list", qmsPlan)
	QMSUpdatePlan        = fmt.Sprintf("%s.update", qmsPlan)
	QMSAddPlan           = fmt.Sprintf("%s.add", qmsPlan)

	AnalysisStatus = fmt.Sprintf("%s.status", cyverseAnalysis)

	MonitoringPing      = fmt.Sprintf("%s.ping", cyverseMonitoring)
	MonitoringDNS       = fmt.Sprintf("%s.dns", cyverseMonitoring)
	MonitoringHeartbeat = fmt.Sprintf("%s.heartbeat", cyverseMonitoring)
)
