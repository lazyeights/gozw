// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package securitypanelzone

// <no value>

type SecurityPanelZoneTypeGet struct {
	ZoneNumber byte
}

func ParseSecurityPanelZoneTypeGet(payload []byte) SecurityPanelZoneTypeGet {
	val := SecurityPanelZoneTypeGet{}

	i := 2

	val.ZoneNumber = payload[i]
	i++

	return val
}