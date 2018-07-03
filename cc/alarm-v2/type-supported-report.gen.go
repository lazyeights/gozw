// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package alarmv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandTypeSupportedReport cc.CommandID = 0x08

func init() {
	gob.Register(TypeSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x08),
		Version:      2,
	}, NewTypeSupportedReport)
}

func NewTypeSupportedReport() cc.Command {
	return &TypeSupportedReport{}
}

// <no value>
type TypeSupportedReport struct {
	Properties1 struct {
		NumberOfBitMasks byte
		V1Alarm          bool
	}
	BitMask []byte
}

func (cmd TypeSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd TypeSupportedReport) CommandID() cc.CommandID {
	return CommandTypeSupportedReport
}

func (cmd TypeSupportedReport) CommandIDString() string {
	return "ALARM_TYPE_SUPPORTED_REPORT"
}

func (cmd *TypeSupportedReport) UnmarshalBinary(data []byte) error {
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
	cmd.Properties1.NumberOfBitMasks = (payload[i] & 0x1F)
	cmd.Properties1.V1Alarm = payload[i]&0x80 == 0x80
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.BitMask = payload[i:]
	return nil
}

func (cmd *TypeSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		var val byte
		val |= (cmd.Properties1.NumberOfBitMasks) & byte(0x1F)
		if cmd.Properties1.V1Alarm {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}
		payload = append(payload, val)
	}
	payload = append(payload, cmd.BitMask...)
	return
}
