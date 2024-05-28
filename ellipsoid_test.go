package nvector_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_GRS80(t *testing.T) {
	got := GRS80
	wantA := 6378137.0
	wantF := 1 / 298.257222101
	wantB := wantA * (1 - wantF)

	if eq, ineq := equality.EqualToFloat64(got.SemiMajorAxis, wantA, 0); !eq {
		equality.ReportInequality(t, "SemiMajorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.SemiMinorAxis, wantB, 0); !eq {
		equality.ReportInequality(t, "SemiMinorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.Flattening, wantF, 0); !eq {
		equality.ReportInequality(t, "Flattening", ineq)
	}
}

func Test_WGS72(t *testing.T) {
	got := WGS72
	wantA := 6378135.0
	wantF := 1 / 298.26
	wantB := wantA * (1 - wantF)

	if eq, ineq := equality.EqualToFloat64(got.SemiMajorAxis, wantA, 0); !eq {
		equality.ReportInequality(t, "SemiMajorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.SemiMinorAxis, wantB, 0); !eq {
		equality.ReportInequality(t, "SemiMinorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.Flattening, wantF, 0); !eq {
		equality.ReportInequality(t, "Flattening", ineq)
	}
}

func Test_WGS84(t *testing.T) {
	got := WGS84
	wantA := 6378137.0
	wantF := 1 / 298.257223563
	wantB := wantA * (1 - wantF)

	if eq, ineq := equality.EqualToFloat64(got.SemiMajorAxis, wantA, 0); !eq {
		equality.ReportInequality(t, "SemiMajorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.SemiMinorAxis, wantB, 0); !eq {
		equality.ReportInequality(t, "SemiMinorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.Flattening, wantF, 0); !eq {
		equality.ReportInequality(t, "Flattening", ineq)
	}
}

func Test_Sphere(t *testing.T) {
	got := Sphere(6371e3)
	wantA := 6371e3
	wantF := 0.0
	wantB := 6371e3

	if eq, ineq := equality.EqualToFloat64(got.SemiMajorAxis, wantA, 0); !eq {
		equality.ReportInequality(t, "SemiMajorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.SemiMinorAxis, wantB, 0); !eq {
		equality.ReportInequality(t, "SemiMinorAxis", ineq)
	}
	if eq, ineq := equality.EqualToFloat64(got.Flattening, wantF, 0); !eq {
		equality.ReportInequality(t, "Flattening", ineq)
	}
}
