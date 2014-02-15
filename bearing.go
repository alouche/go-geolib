package geolib

import (
	"errors"
	"math"
)

// Bearing {from, to} a point
// flag: 0    => Initial Bearing
// flag: 1    => Initial Bearing
// flag: nil  => Bearing
// Result:
//  - Type: float64
//  - Metric: Degress from North
func Bearing(flag int, φ1, λ1, φ2, λ2 float64) (float64, error) {
	φ1 = Deg2Rad(φ1)
	φ2 = Deg2Rad(φ2)

	Δλ := Deg2Rad(λ2 - λ1)

	bearing := Rad2Deg(
		math.Atan2(
			math.Sin(
				Δλ)*math.Cos(φ2), math.Cos(φ1)*math.Sin(φ2)-math.Sin(φ1)*math.Cos(φ2)*math.Cos(Δλ)))

	var incDeg float64

	switch flag {
	case 0:
		incDeg = 360
	case 1:
		incDeg = 180
	default:
		return -1, errors.New("flag not supported")
	}

	return math.Mod(bearing+incDeg, 360), nil
}
