// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensorconfiguration

// <no value>

type SensorTriggerLevelSet struct {
	Properties1 struct {
		Current bool

		Default bool
	}

	SensorType byte

	Properties2 struct {
		Size byte

		Scale byte

		Precision byte
	}

	TriggerValue []byte
}

func ParseSensorTriggerLevelSet(payload []byte) SensorTriggerLevelSet {
	val := SensorTriggerLevelSet{}

	i := 2

	if payload[i]&0x40 == 0x40 {
		val.Properties1.Current = true
	} else {
		val.Properties1.Current = false
	}

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Default = true
	} else {
		val.Properties1.Default = false
	}

	i += 1

	val.SensorType = payload[i]
	i++

	val.Properties2.Size = (payload[i] & 0x07)

	val.Properties2.Scale = (payload[i] & 0x18) << 3

	val.Properties2.Precision = (payload[i] & 0xE0) << 5

	i += 1

	val.TriggerValue = payload[i : i+2]
	i += 2

	return val
}