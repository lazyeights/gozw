// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package basicwindowcovering

// <no value>

type BasicWindowCoveringStartLevelChange struct {
	Level struct {
		OpenClose bool
	}
}

func ParseBasicWindowCoveringStartLevelChange(payload []byte) BasicWindowCoveringStartLevelChange {
	val := BasicWindowCoveringStartLevelChange{}

	i := 2

	if payload[i]&0x40 == 0x40 {
		val.Level.OpenClose = true
	} else {
		val.Level.OpenClose = false
	}

	i += 1

	return val
}