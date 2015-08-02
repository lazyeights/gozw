package commandclass

const (
	CommandAlarmVersion byte = 0x01
	CommandAlarmGet          = 0x04
	CommandAlarmReport       = 0x05
)

type AlarmReport struct {
	Type  byte
	Level byte
}

// @todo handle multiple versions, since this one has a lot
func ParseAlarmReport(payload []byte) AlarmReport {
	return AlarmReport{
		Type:  payload[2],
		Level: payload[3],
	}
}