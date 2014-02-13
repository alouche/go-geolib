package geolib

import (
	"fmt"
	"testing"
)

func haversine_distance(t *testing.T, expect string, φ1, λ1, φ2, λ2 float64) {
	if result := fmt.Sprintf("%.2f", Haversine_Distance(φ1, λ1, φ2, λ2)); expect != result {
		t.Errorf("Haversine Distance = expect %f; result %f", expect, result)
	}
}

func TestHaversine_Distance(t *testing.T) {
	haversine_distance(t, "421.77", 50.116667, 8.683333, 52.516667, 13.383333)
}
