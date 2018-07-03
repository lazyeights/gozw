// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package timev2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandDateGet cc.CommandID = 0x03

func init() {
	gob.Register(DateGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8A),
		Command:      cc.CommandID(0x03),
		Version:      2,
	}, NewDateGet)
}

func NewDateGet() cc.Command {
	return &DateGet{}
}

// <no value>
type DateGet struct {
}

func (cmd DateGet) CommandClassID() cc.CommandClassID {
	return 0x8A
}

func (cmd DateGet) CommandID() cc.CommandID {
	return CommandDateGet
}

func (cmd DateGet) CommandIDString() string {
	return "DATE_GET"
}

func (cmd *DateGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	return nil
}

func (cmd *DateGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	return
}
