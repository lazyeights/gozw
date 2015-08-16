// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev3

// <no value>

type ThermostatModeReport struct {
	Level struct {
		NoOfManufacturerDataFields byte

		Mode byte
	}

	ManufacturerData []byte
}

func ParseThermostatModeReport(payload []byte) ThermostatModeReport {
	val := ThermostatModeReport{}

	i := 2

	val.Level.NoOfManufacturerDataFields = (payload[i] & 0xE0) << 5

	val.Level.Mode = (payload[i] & 0x1F)

	i += 1

	val.ManufacturerData = payload[i : i+0]
	i += 0

	return val
}