// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package notificationv5

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandEventSupportedGet cc.CommandID = 0x01

func init() {
	gob.Register(EventSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x01),
		Version:      5,
	}, NewEventSupportedGet)
}

func NewEventSupportedGet() cc.Command {
	return &EventSupportedGet{}
}

// <no value>
type EventSupportedGet struct {
	NotificationType byte
}

func (cmd EventSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd EventSupportedGet) CommandID() cc.CommandID {
	return CommandEventSupportedGet
}

func (cmd EventSupportedGet) CommandIDString() string {
	return "EVENT_SUPPORTED_GET"
}

func (cmd *EventSupportedGet) UnmarshalBinary(data []byte) error {
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
	cmd.NotificationType = payload[i]
	i++
	return nil
}

func (cmd *EventSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	payload = append(payload, cmd.NotificationType)
	return
}
