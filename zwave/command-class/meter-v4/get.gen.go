// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package meterv4

// <no value>

type MeterGet struct {
	Properties1 struct {
		Scale byte

		RateType byte
	}

	Scale2 byte
}

func ParseMeterGet(payload []byte) MeterGet {
	val := MeterGet{}

	i := 2

	val.Properties1.Scale = (payload[i] & 0x38) << 3

	val.Properties1.RateType = (payload[i] & 0xC0) << 6

	i += 1

	val.Scale2 = payload[i]
	i++

	return val
}