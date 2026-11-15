package nvector_test

import (
	"math"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_VectorNormalize(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want Vector
	}{
		{"zero vector", Vector{}, Vector{X: 1}},
		{"unit vector", Vector{X: 3, Y: 4}, Vector{X: 0.6, Y: 0.8}},
		{
			"large vector",
			Vector{X: 1e200, Y: 1e200},
			Vector{X: math.Sqrt2 / 2, Y: math.Sqrt2 / 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Normalize()

			if eq, ineq := equality.EqualToVector(got, tt.want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	}
}
