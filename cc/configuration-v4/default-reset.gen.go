// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv4

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandDefaultReset cc.CommandID = 0x01

func init() {
	gob.Register(DefaultReset{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x01),
		Version:      4,
	}, NewDefaultReset)
}

func NewDefaultReset() cc.Command {
	return &DefaultReset{}
}

// <no value>
type DefaultReset struct {
}

func (cmd DefaultReset) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd DefaultReset) CommandID() cc.CommandID {
	return CommandDefaultReset
}

func (cmd DefaultReset) CommandIDString() string {
	return "CONFIGURATION_DEFAULT_RESET"
}

func (cmd *DefaultReset) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	return nil
}

func (cmd *DefaultReset) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	return
}
