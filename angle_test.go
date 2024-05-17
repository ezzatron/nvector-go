package nvector_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_Degrees(t *testing.T) {
	t.Run("it converts an angle in radians to degrees", func(t *testing.T) {
		got := Degrees(1)
		want := 57.29577951308232

		if eq, ineq := equality.EqualToFloat64(got, want, 1e-15); !eq {
			equality.ReportInequality(t, "degrees", ineq)
		}
	})
}

func Test_Radians(t *testing.T) {
	t.Run("it converts an angle in degrees to radians", func(t *testing.T) {
		got := Radians(57.29577951308232)
		want := 1.0

		if eq, ineq := equality.EqualToFloat64(got, want, 1e-15); !eq {
			equality.ReportInequality(t, "radians", ineq)
		}
	})
}
