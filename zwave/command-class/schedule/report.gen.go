// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package schedule

import "encoding/binary"

// <no value>

type CommandScheduleReport struct {
	ScheduleId byte

	UserIdentifier byte

	StartYear byte

	Properties1 struct {
		StartMonth byte

		ActiveId byte
	}

	Properties2 struct {
		StartDayOfMonth byte
	}

	Properties3 struct {
		StartWeekday byte

		Res bool
	}

	Properties4 struct {
		StartHour byte

		DurationType byte
	}

	Properties5 struct {
		StartMinute byte
	}

	DurationByte uint16

	ReportsToFollow byte

	NumberOfCmdToFollow byte
}

func ParseCommandScheduleReport(payload []byte) CommandScheduleReport {
	val := CommandScheduleReport{}

	i := 2

	val.ScheduleId = payload[i]
	i++

	val.UserIdentifier = payload[i]
	i++

	val.StartYear = payload[i]
	i++

	val.Properties1.StartMonth = (payload[i] & 0x0F)

	val.Properties1.ActiveId = (payload[i] & 0xF0) << 4

	i += 1

	val.Properties2.StartDayOfMonth = (payload[i] & 0x1F)

	i += 1

	val.Properties3.StartWeekday = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Properties3.Res = true
	} else {
		val.Properties3.Res = false
	}

	i += 1

	val.Properties4.StartHour = (payload[i] & 0x1F)

	val.Properties4.DurationType = (payload[i] & 0xE0) << 5

	i += 1

	val.Properties5.StartMinute = (payload[i] & 0x3F)

	i += 1

	val.DurationByte = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.ReportsToFollow = payload[i]
	i++

	val.NumberOfCmdToFollow = payload[i]
	i++

	return val
}