// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package configurationv2

// <no value>

type ConfigurationSet struct {
	ParameterNumber byte

	Level struct {
		Size byte

		Default bool
	}

	ConfigurationValue []byte
}

func ParseConfigurationSet(payload []byte) ConfigurationSet {
	val := ConfigurationSet{}

	i := 2

	val.ParameterNumber = payload[i]
	i++

	val.Level.Size = (payload[i] & 0x07)

	if payload[i]&0x80 == 0x80 {
		val.Level.Default = true
	} else {
		val.Level.Default = false
	}

	i += 1

	val.ConfigurationValue = payload[i : i+1]
	i += 1

	return val
}