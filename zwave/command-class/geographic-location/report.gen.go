// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package geographiclocation

// <no value>

type GeographicLocationReport struct {
	LongitudeDegrees byte

	Level struct {
		LongitudeMinutes byte

		LongSign bool
	}

	LatitudeDegrees byte

	Level2 struct {
		LatitudeMinutes byte

		LatSign bool
	}
}

func ParseGeographicLocationReport(payload []byte) GeographicLocationReport {
	val := GeographicLocationReport{}

	i := 2

	val.LongitudeDegrees = payload[i]
	i++

	val.Level.LongitudeMinutes = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Level.LongSign = true
	} else {
		val.Level.LongSign = false
	}

	i += 1

	val.LatitudeDegrees = payload[i]
	i++

	val.Level2.LatitudeMinutes = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Level2.LatSign = true
	} else {
		val.Level2.LatSign = false
	}

	i += 1

	return val
}