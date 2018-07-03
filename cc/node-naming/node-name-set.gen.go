// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package nodenaming

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeNameSet cc.CommandID = 0x01

func init() {
	gob.Register(NodeNameSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x77),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewNodeNameSet)
}

func NewNodeNameSet() cc.Command {
	return &NodeNameSet{}
}

// <no value>
type NodeNameSet struct {
	Level struct {
		CharPresentation byte
	}
	NodeNameChar string
}

func (cmd NodeNameSet) CommandClassID() cc.CommandClassID {
	return 0x77
}

func (cmd NodeNameSet) CommandID() cc.CommandID {
	return CommandNodeNameSet
}

func (cmd NodeNameSet) CommandIDString() string {
	return "NODE_NAMING_NODE_NAME_SET"
}

func (cmd *NodeNameSet) UnmarshalBinary(data []byte) error {
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
	cmd.Level.CharPresentation = (payload[i] & 0x07)
	i += 1
	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}
	cmd.NodeNameChar = string(payload[i : i+16])
	i += 16
	return nil
}

func (cmd *NodeNameSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())
	{
		var val byte
		val |= (cmd.Level.CharPresentation) & byte(0x07)
		payload = append(payload, val)
	}
	if paramLen := len(cmd.NodeNameChar); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter NodeNameChar")
	}
	payload = append(payload, []byte(cmd.NodeNameChar)...)
	return
}
