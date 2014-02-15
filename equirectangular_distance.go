package geolib

import "math"

// EquirectangularDistance calculates the distance in KM of 2 points using Pythagoras’ theorem.
// It is a faster alternative to the Haversine_Distance function with less accuracy.
// Scope: small distance calculation
// Result:
//  - Type: float64
//  - Result: Distance (km)
func EquirectangularDistance(φ1, λ1, φ2, λ2 float64) float64 {
	return math.Sqrt(
		math.Pow((λ2-λ1)*math.Cos((φ1+φ2)/2), 2) + math.Pow(φ2-φ1, 2)*GreatEarthCircleRadiusKM)
}
