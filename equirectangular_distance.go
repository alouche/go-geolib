package geolib

import "math"

func Equirectangular_Distance(φ1, λ1, φ2, λ2 float64) float64 {
	return math.Sqrt(
		math.Pow((λ2-λ1)*math.Cos((φ1+φ2)/2), 2) + math.Pow(φ2-φ1, 2)*GREAT_CIRCLE_RADIUS_KM)
}
