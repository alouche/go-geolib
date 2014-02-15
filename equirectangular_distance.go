package geolib

import "math"

// Equirectangular_Distance calculates the distance in KM of 2 points using Pythagoras’ theorem.
// It is a faster alternative to the Haversine_Distance function with less accuracy.
// Scope: small distance calculation

func Equirectangular_Distance(φ1, λ1, φ2, λ2 float64) float64 {
	return math.Sqrt(
		math.Pow((λ2-λ1)*math.Cos((φ1+φ2)/2), 2) + math.Pow(φ2-φ1, 2)*GREAT_CIRCLE_RADIUS_KM)
}
