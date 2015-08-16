// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sceneactuatorconf

// <no value>

type SceneActuatorConfSet struct {
	SceneId byte

	DimmingDuration byte

	Level2 struct {
		Override bool
	}

	Level byte
}

func ParseSceneActuatorConfSet(payload []byte) SceneActuatorConfSet {
	val := SceneActuatorConfSet{}

	i := 2

	val.SceneId = payload[i]
	i++

	val.DimmingDuration = payload[i]
	i++

	if payload[i]&0x80 == 0x80 {
		val.Level2.Override = true
	} else {
		val.Level2.Override = false
	}

	i += 1

	val.Level = payload[i]
	i++

	return val
}