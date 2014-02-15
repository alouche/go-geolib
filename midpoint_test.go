package geolib

import (
	"fmt"
	"testing"
)

func midpoint(t *testing.T, latE string, lonE string, φ1, λ1, φ2, λ2 float64) {
	point := Midpoint(φ1, λ1, φ2, λ2)
	latR := fmt.Sprintf("%.2f", point.lat)
	lonR := fmt.Sprintf("%.2f", point.lon)

	if latR != latE || lonR != lonE {
		t.Errorf("Midpoint = expect %f %f; result %f %f", latE, lonE, latR, lonR)
	}
}

func TestMidpoint(t *testing.T) {
	midpoint(t, "51.34", "10.97", 50.116667, 8.683333, 52.516667, 13.383333)
}
