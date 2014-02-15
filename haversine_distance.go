package geolib

import "math"

func haversin(θ float64) float64 {
	return (1 - math.Cos(θ)) / 2
}

// HaversineDistance calculates the great-circle distance between 2 points.
// The function uses a spherical model, since the earth is more is oblate spheroidal, there will be an accuracy error margin.
// Scope: Distance computation below 0.3% accuracy error margin
// Result:
//  - Type: float64
//  - Metric: Distance (km)
func HaversineDistance(φ1, λ1, φ2, λ2 float64) float64 {
	φ1 = Deg2Rad(φ1)
	φ2 = Deg2Rad(φ2)
	λ1 = Deg2Rad(λ1)
	λ2 = Deg2Rad(λ2)

	return 2 * GreatEarthCircleRadiusKM *
		math.Asin(
			math.Sqrt(
				haversin(φ2-φ1)+
					math.Cos(φ1)*math.Cos(φ2)*haversin(λ2-λ1)))
}
