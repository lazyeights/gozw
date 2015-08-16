// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv3

// <no value>

type NotificationReport struct {
	V1AlarmType byte

	V1AlarmLevel byte

	ZensorNetSourceNodeId byte

	NotificationStatus byte

	NotificationType byte

	Event byte

	Properties1 struct {
		EventParametersLength byte

		Sequence bool
	}

	EventParameter []byte

	SequenceNumber byte
}

func ParseNotificationReport(payload []byte) NotificationReport {
	val := NotificationReport{}

	i := 2

	val.V1AlarmType = payload[i]
	i++

	val.V1AlarmLevel = payload[i]
	i++

	val.ZensorNetSourceNodeId = payload[i]
	i++

	val.NotificationStatus = payload[i]
	i++

	val.NotificationType = payload[i]
	i++

	val.Event = payload[i]
	i++

	val.Properties1.EventParametersLength = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		val.Properties1.Sequence = true
	} else {
		val.Properties1.Sequence = false
	}

	i += 1

	val.EventParameter = payload[i : i+6]
	i += 6

	val.SequenceNumber = payload[i]
	i++

	return val
}