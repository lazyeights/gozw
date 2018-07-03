// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatmodev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x40),
		Command:      cc.CommandID(0x01),
		Version:      2,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	Level struct {
		Mode byte
	}
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x40
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "THERMOSTAT_MODE_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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
	cmd.Level.Mode = (payload[i] & 0x1F)
	i += 1
	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		var val byte
		val |= (cmd.Level.Mode) & byte(0x1F)
		payload = append(payload, val)
	}
	return
}
