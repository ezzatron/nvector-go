package ellipsoid_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go/ellipsoid"
)

func Test_GRS80(t *testing.T) {
	e := GRS80()

	if e.SemiMajorAxis != 6378137.0 {
		t.Errorf("Expected SemiMajorAxis to be 6378137.0, got %f", e.SemiMajorAxis)
	}
	if e.Flattening != 1/298.257222101 {
		t.Errorf("Expected Flattening to be 1/298.257222101, got %f", e.Flattening)
	}
}

func Test_WGS72(t *testing.T) {
	e := WGS72()

	if e.SemiMajorAxis != 6378135.0 {
		t.Errorf("Expected SemiMajorAxis to be 6378135.0, got %f", e.SemiMajorAxis)
	}
	if e.Flattening != 1/298.26 {
		t.Errorf("Expected Flattening to be 1/298.26, got %f", e.Flattening)
	}
}

func Test_WGS84(t *testing.T) {
	e := WGS84()

	if e.SemiMajorAxis != 6378137.0 {
		t.Errorf("Expected SemiMajorAxis to be 6378137.0, got %f", e.SemiMajorAxis)
	}
	if e.Flattening != 1/298.257223563 {
		t.Errorf("Expected Flattening to be 1/298.257223563, got %f", e.Flattening)
	}
}

func Test_WGS84Sphere(t *testing.T) {
	e := WGS84Sphere()

	if e.SemiMajorAxis != 6378137.0 {
		t.Errorf("Expected SemiMajorAxis to be 6378137.0, got %f", e.SemiMajorAxis)
	}
	if e.Flattening != 0 {
		t.Errorf("Expected Flattening to be 0, got %f", e.Flattening)
	}
}
