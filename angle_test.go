package nvector_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/floats/scalar"
)

func Test_Deg(t *testing.T) {
	t.Run("it converts an angle in radians to degrees", func(t *testing.T) {
		got := Deg(1)
		want := 57.29577951308232

		if !scalar.EqualWithinAbs(got, want, 1e-15) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func Test_Rad(t *testing.T) {
	t.Run("it converts an angle in degrees to radians", func(t *testing.T) {
		got := Rad(57.29577951308232)
		want := 1.0

		if !scalar.EqualWithinAbs(got, want, 1e-15) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
