package geolib

import (
	"fmt"
	"testing"
)

func deg2rad(t *testing.T, expect string, num float64) {
	if result := fmt.Sprintf("%.2f", Deg2Rad(num)); result != expect {
		t.Errorf("Deg2Rad = expect %f; result %f", expect, result)
	}
}

func TestDeg2Rad(t *testing.T) {
	deg2rad(t, "0.87", 50.116667)
	deg2rad(t, "0.15", 8.683333)
	deg2rad(t, "0.92", 52.516667)
	deg2rad(t, "0.23", 13.383333)
}
