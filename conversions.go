package geolib

import "math"

// GREAT_CIRCLE_RADIUS_KM varies between 6,378km and 6,357km (equatorial and polar).
// The local radius of curvature varies between 6,336km and 6,399km (equatorial meridian and polar)
var GREAT_CIRCLE_RADIUS_KM float64 = 6372.8

// Deg2Rad converts a degree to randian
func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}

// Rad2Deg converts a randian to degree
func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}
