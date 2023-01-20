package pbinit

import (
	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/p/go/monitoring"
)

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
