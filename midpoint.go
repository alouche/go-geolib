package geolib

import "math"

type Point struct {
	lat, lon float64
}

// Midpoint calculates the half-way point along a great circle path between the two points
// Result:
//  - Type: struct{float64, float64}
//  - Metric: Radian
func Midpoint(φ1, λ1, φ2, λ2 float64) Point {
	var point = Point{}

	φ1 = Deg2Rad(φ1)
	λ1 = Deg2Rad(λ1)
	φ2 = Deg2Rad(φ2)

	Δλ := Deg2Rad(λ2) - λ1

	Bx := math.Cos(φ2) * math.Cos(Δλ)
	By := math.Cos(φ2) * math.Sin(Δλ)

	point.lat = Rad2Deg(math.Atan2(math.Sin(φ1)+math.Sin(φ2), math.Sqrt(math.Pow(math.Cos(φ1)+Bx, 2)+math.Pow(By, 2))))
	point.lon = Rad2Deg(λ1 + math.Atan2(By, math.Cos(φ1)+Bx))

	return point
}
