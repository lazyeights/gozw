// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package applicationstatus

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandApplicationBusy cc.CommandID = 0x01

func init() {
	gob.Register(ApplicationBusy{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x22),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewApplicationBusy)
}

func NewApplicationBusy() cc.Command {
	return &ApplicationBusy{}
}

// <no value>
type ApplicationBusy struct {
	Status   byte
	WaitTime byte
}

func (cmd ApplicationBusy) CommandClassID() cc.CommandClassID {
	return 0x22
}

func (cmd ApplicationBusy) CommandID() cc.CommandID {
	return CommandApplicationBusy
}

func (cmd ApplicationBusy) CommandIDString() string {
	return "APPLICATION_BUSY"
}

func (cmd *ApplicationBusy) UnmarshalBinary(data []byte) error {
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
	cmd.Status = payload[i]
	i++
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.WaitTime = payload[i]
	i++
	return nil
}

func (cmd *ApplicationBusy) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	payload = append(payload, cmd.Status)
	payload = append(payload, cmd.WaitTime)
	return
}
