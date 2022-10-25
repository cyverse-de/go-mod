package subjects

import "fmt"

const qmsUser = "cyverse.qms.user"

var (
	QMSGetUserUpdates = fmt.Sprintf("%s.updates.get", qmsUser)
	QMSAddUserUpdate  = fmt.Sprintf("%s.updates.add", qmsUser)

	QMSGetUserOverages   = fmt.Sprintf("%s.overages.get", qmsUser)
	QMSCheckUserOverages = fmt.Sprintf("%s.overages.check", qmsUser)

	QMSGetUserUsages = fmt.Sprintf("%s.usages.get", qmsUser)
	QMSAddUserUsages = fmt.Sprintf("%s.usages.add", qmsUser)
)
