package geolib

import (
	"fmt"
	"testing"
)

func midpoint(t *testing.T, lat_e string, lon_e string, φ1, λ1, φ2, λ2 float64) {
	point := Midpoint(φ1, λ1, φ2, λ2)
	lat_r := fmt.Sprintf("%.2f", point.lat)
	lon_r := fmt.Sprintf("%.2f", point.lon)

	if lat_r != lat_e || lon_r != lon_e {
		t.Errorf("Midpoint = expect %f %f; result %f %f", lat_e, lon_e, lat_r, lon_e)
	}
}

func TestMidpoint(t *testing.T) {
	midpoint(t, "51.34", "10.97", 50.116667, 8.683333, 52.516667, 13.383333)
}
