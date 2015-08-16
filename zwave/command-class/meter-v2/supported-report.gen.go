// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package meterv2

// <no value>

type MeterSupportedReport struct {
	Properties1 struct {
		MeterType byte

		MeterReset bool
	}

	Properties2 struct {
		ScaleSupported byte
	}
}

func ParseMeterSupportedReport(payload []byte) MeterSupportedReport {
	val := MeterSupportedReport{}

	i := 2

	val.Properties1.MeterType = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.MeterReset = true
	} else {
		val.Properties1.MeterReset = false
	}

	i += 1

	val.Properties2.ScaleSupported = (payload[i] & 0x0F)

	i += 1

	return val
}