// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package climatecontrolschedule

// <no value>

type ScheduleGet struct {
	Properties1 struct {
		Weekday byte
	}
}

func ParseScheduleGet(payload []byte) ScheduleGet {
	val := ScheduleGet{}

	i := 2

	val.Properties1.Weekday = (payload[i] & 0x07)

	i += 1

	return val
}