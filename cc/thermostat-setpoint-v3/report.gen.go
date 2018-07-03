// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatsetpointv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x03

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x43),
		Command:      cc.CommandID(0x03),
		Version:      3,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	Level struct {
		SetpointType byte
	}
	Level2 struct {
		Size      byte
		Scale     byte
		Precision byte
	}
	Value []byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x43
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "THERMOSTAT_SETPOINT_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}
	i := 2
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.Level.SetpointType = (payload[i] & 0x0F)
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.Level2.Size = (payload[i] & 0x07)
	cmd.Level2.Scale = (payload[i] & 0x18) >> 3
	cmd.Level2.Precision = (payload[i] & 0xE0) >> 5
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	{
		length := (payload[1+2] >> 0) & 0x07
		cmd.Value = payload[i : i+int(length)]
		i += int(length)
	}
	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		var val byte
		val |= (cmd.Level.SetpointType) & byte(0x0F)
		payload = append(payload, val)
	}
	{
		var val byte
		val |= (cmd.Level2.Size) & byte(0x07)
		val |= (cmd.Level2.Scale << byte(3)) & byte(0x18)
		val |= (cmd.Level2.Precision << byte(5)) & byte(0xE0)
		payload = append(payload, val)
	}
	if cmd.Value != nil && len(cmd.Value) > 0 {
		payload = append(payload, cmd.Value...)
	}
	return
}
