package equality_test

import (
	"math"
	"testing"

	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_EqualToRadians(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
		eq   bool
	}{
		{"identical angles", 0, 0, 0, true},
		{"small difference", 0, 0.1, 0.1, false},
		{"small difference within tolerance", 0, 0.1e-15, 0.1e-15, true},
		{"negative difference", -1, 1, 2, false},
		{"antipodal angles", 0, math.Pi, math.Pi, false},
		{"wrapping across zero", 0.1, 2*math.Pi - 0.1, 0.2, false},
		{"difference greater than a full turn", 0, 2*math.Pi + 0.5, 0.5, false},
		{
			"negative difference less than a full turn",
			0,
			-2*math.Pi - 0.5,
			0.5,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eq, ineq := equality.EqualToRadians(tt.a, tt.b, 1e-15)

			if eq != tt.eq {
				t.Errorf("got equal %v; want %v", eq, tt.eq)
			}
			if got := ineq.Diff; math.Abs(got-tt.want) > 1e-14 {
				t.Errorf("got difference %v; want %v", got, tt.want)
			}
		})
	}
}
