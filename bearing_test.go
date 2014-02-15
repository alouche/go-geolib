package geolib

import (
	"fmt"
	"testing"
)

func bearing(t *testing.T, flag int, expect string, φ1, λ1, φ2, λ2 float64) {
	bearing, _ := Bearing(flag, φ1, λ1, φ2, λ2)
	if result := fmt.Sprintf("%.2f", bearing); expect != result {
		t.Errorf("Bearing(%d) = expect %f; result %f", flag, expect, result)
	}
}

func TestInitial_Bearing(t *testing.T) {
	bearing(t, 0, "48.93", 50.116667, 8.683333, 52.516667, 13.383333)
}

func TestFinal_Bearing(t *testing.T) {
	bearing(t, 1, "52.60", 52.516667, 13.383333, 50.116667, 8.683333)
}
