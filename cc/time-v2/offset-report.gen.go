// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package timev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandOffsetReport cc.CommandID = 0x07

func init() {
	gob.Register(OffsetReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8A),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewOffsetReport)
}

func NewOffsetReport() cc.Command {
	return &OffsetReport{}
}

// <no value>
type OffsetReport struct {
	Level struct {
		HourTzo byte
		SignTzo bool
	}
	MinuteTzo byte
	Level2    struct {
		MinuteOffsetDst byte
		SignOffsetDst   bool
	}
	MonthStartDst byte
	DayStartDst   byte
	HourStartDst  byte
	MonthEndDst   byte
	DayEndDst     byte
	HourEndDst    byte
}

func (cmd OffsetReport) CommandClassID() cc.CommandClassID {
	return 0x8A
}

func (cmd OffsetReport) CommandID() cc.CommandID {
	return CommandOffsetReport
}

func (cmd OffsetReport) CommandIDString() string {
	return "TIME_OFFSET_REPORT"
}

func (cmd *OffsetReport) UnmarshalBinary(data []byte) error {
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
	cmd.Level.HourTzo = (payload[i] & 0x7F)
	cmd.Level.SignTzo = payload[i]&0x80 == 0x80
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.MinuteTzo = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.Level2.MinuteOffsetDst = (payload[i] & 0x7F)
	cmd.Level2.SignOffsetDst = payload[i]&0x80 == 0x80
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.MonthStartDst = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.DayStartDst = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.HourStartDst = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.MonthEndDst = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.DayEndDst = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.HourEndDst = payload[i]
	i++
	return nil
}

func (cmd *OffsetReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		var val byte
		val |= (cmd.Level.HourTzo) & byte(0x7F)
		if cmd.Level.SignTzo {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}
		payload = append(payload, val)
	}
	payload = append(payload, cmd.MinuteTzo)
	{
		var val byte
		val |= (cmd.Level2.MinuteOffsetDst) & byte(0x7F)
		if cmd.Level2.SignOffsetDst {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}
		payload = append(payload, val)
	}
	payload = append(payload, cmd.MonthStartDst)
	payload = append(payload, cmd.DayStartDst)
	payload = append(payload, cmd.HourStartDst)
	payload = append(payload, cmd.MonthEndDst)
	payload = append(payload, cmd.DayEndDst)
	payload = append(payload, cmd.HourEndDst)
	return
}
