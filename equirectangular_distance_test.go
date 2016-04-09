package geolib

import (
	"fmt"
	"testing"
)

func equirectangularDistance(t *testing.T, expect string, φ1, λ1, φ2, λ2 float64) {
	if result := fmt.Sprintf("%.2f", EquirectangularDistance(φ1, λ1, φ2, λ2)); expect != result {
		t.Errorf("Equirectangular Distance = expect %s; result %s", expect, result)
	}
}

func TestEquirectangularDistance(t *testing.T) {
	equirectangularDistance(t, "191.61", 50.116667, 8.683333, 52.516667, 13.383333)
}
