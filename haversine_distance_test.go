package geolib

import (
	"fmt"
	"testing"
)

func haversineDistance(t *testing.T, expect string, φ1, λ1, φ2, λ2 float64) {
	if result := fmt.Sprintf("%.2f", HaversineDistance(φ1, λ1, φ2, λ2)); expect != result {
		t.Errorf("Haversine Distance = expect %f; result %f", expect, result)
	}
}

func TestHaversineDistance(t *testing.T) {
	haversineDistance(t, "421.77", 50.116667, 8.683333, 52.516667, 13.383333)
}
