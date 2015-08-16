// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

// <no value>

type ChimneyFanAlarmTempSet struct {
	Properties1 struct {
		Size byte

		Scale byte

		Precision byte
	}

	Value []byte
}

func ParseChimneyFanAlarmTempSet(payload []byte) ChimneyFanAlarmTempSet {
	val := ChimneyFanAlarmTempSet{}

	i := 2

	val.Properties1.Size = (payload[i] & 0x07)

	val.Properties1.Scale = (payload[i] & 0x18) << 3

	val.Properties1.Precision = (payload[i] & 0xE0) << 5

	i += 1

	val.Value = payload[i : i+0]
	i += 0

	return val
}