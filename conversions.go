package geolib

import "math"

const GREAT_CIRCLE_RADIUS_KM = 6372.8

// Deg2Rad converts a degree to randian
func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}

// Rad2Deg converts a randian to degree
func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}
