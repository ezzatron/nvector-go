package nvector_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_ToECEFKnownValues(t *testing.T) {
	tests := []struct {
		name string
		n    Vector
		want Vector
	}{
		{"equator", Vector{X: 1}, Vector{X: WGS84.SemiMajorAxis}},
		{"east equator", Vector{Y: 1}, Vector{Y: WGS84.SemiMajorAxis}},
		{"north pole", Vector{Z: 1}, Vector{Z: WGS84.SemiMinorAxis}},
		{"south pole", Vector{Z: -1}, Vector{Z: -WGS84.SemiMinorAxis}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToECEF(Position{Vector: tt.n}, WGS84, ZAxisNorth)

			if eq, ineq := equality.EqualToVector(got, tt.want, 1e-6); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	}
}

func Test_ECEFRoundTripEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		n     Vector
		depth float64
	}{
		{"equator", Vector{X: 1}, 0},
		{"east equator", Vector{Y: 1}, 0},
		{"north pole", Vector{Z: 1}, 0},
		{"south pole", Vector{Z: -1}, 0},
		{"north pole at depth", Vector{Z: 1}, 1000},
		{"north pole above the surface", Vector{Z: 1}, -10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := Position{Vector: tt.n, Depth: tt.depth}
			got := FromECEF(
				ToECEF(want, WGS84, ZAxisNorth),
				WGS84,
				ZAxisNorth,
			)

			if eq, ineq := equality.EqualToVectorWithDepth(got, want, 1e-14, 1e-6); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	}
}

func Test_FromECEFToECEFEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		p    Vector
	}{
		{"x is zero", Vector{Y: 3e6, Z: -3e6}},
		{"y plus z is zero", Vector{X: WGS84.SemiMajorAxis, Y: 3e3, Z: -3e3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToECEF(FromECEF(tt.p, WGS84, ZAxisNorth), WGS84, ZAxisNorth)

			if eq, ineq := equality.EqualToVector(got, tt.p, 1e-6); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	}
}

func Test_ToRotationMatrixNearPole(t *testing.T) {
	identity := Matrix{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	tests := []struct {
		name string
		n    Vector
	}{
		{"x offset", Vector{X: 1e-12, Z: 1}.Normalize()},
		{"negative x offset", Vector{X: -1e-12, Z: 1}.Normalize()},
		{"y offset", Vector{Y: 1e-12, Z: 1}.Normalize()},
		{"negative y offset", Vector{Y: -1e-12, Z: 1}.Normalize()},
		{"both offsets", Vector{X: 1e-12, Y: 1e-12, Z: 1}.Normalize()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ToRotationMatrix(tt.n, ZAxisNorth)

			if eq, ineq := equality.EqualToMatrix(
				r.Transpose().Multiply(r),
				identity,
				1e-14,
			); !eq {
				equality.ReportInequalities(t, ineq)
			}
			if eq, ineq := equality.EqualToVector(
				FromRotationMatrix(r),
				tt.n,
				1e-14,
			); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	}
}
