// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package wakeup

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandIntervalReport cc.CommandID = 0x06

func init() {
	gob.Register(IntervalReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x84),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewIntervalReport)
}

func NewIntervalReport() cc.Command {
	return &IntervalReport{}
}

// <no value>
type IntervalReport struct {
	Seconds uint32
	Nodeid  byte
}

func (cmd IntervalReport) CommandClassID() cc.CommandClassID {
	return 0x84
}

func (cmd IntervalReport) CommandID() cc.CommandID {
	return CommandIntervalReport
}

func (cmd IntervalReport) CommandIDString() string {
	return "WAKE_UP_INTERVAL_REPORT"
}

func (cmd *IntervalReport) UnmarshalBinary(data []byte) error {
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
	cmd.Seconds = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.Nodeid = payload[i]
	i++
	return nil
}

func (cmd *IntervalReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Seconds)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}
	payload = append(payload, cmd.Nodeid)
	return
}
