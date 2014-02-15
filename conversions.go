package geolib

import "math"

// GreatEarthCircleRadiusKM varies between 6,378km and 6,357km (equatorial and polar).
// The local radius of curvature varies between 6,336km and 6,399km (equatorial meridian and polar)
var GreatEarthCircleRadiusKM = 6372.8

// Deg2Rad converts a degree to randian
// Result:
//  - Type: float64
//  - Metric: Radian
func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}

// Rad2Deg converts a randian to degree
// Result:
//  - Type: float64
//  - Metric: Degree
func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}
