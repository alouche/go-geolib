package geolib

import (
	"fmt"
	"testing"
)

func equirectangular_distance(t *testing.T, expect string, φ1, λ1, φ2, λ2 float64) {
	if result := fmt.Sprintf("%.2f", Equirectangular_Distance(φ1, λ1, φ2, λ2)); expect != result {
		t.Errorf("Haversine Distance = expect %f; result %f", expect, result)
	}
}

func TestEquirectangular_Distance(t *testing.T) {
	equirectangular_distance(t, "191.61", 50.116667, 8.683333, 52.516667, 13.383333)
}
