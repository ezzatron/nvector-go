package coordframe_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go/coordframe"
)

func Test_ZAxisNorth(t *testing.T) {
	want := []float64{
		0, 0, 1,
		0, 1, 0,
		-1, 0, 0,
	}

	got := ZAxisNorth()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			wantN := want[i*3+j]
			gotN := got.At(i, j)

			if gotN != wantN {
				t.Errorf(
					"ZAxisNorth() = %v,%v: %v; want %v,%v: %v",
					i,
					j,
					gotN,
					i,
					j,
					wantN,
				)
			}
		}
	}
}

func Test_XAxisNorth(t *testing.T) {
	want := []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	got := XAxisNorth()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			wantN := want[i*3+j]
			gotN := got.At(i, j)

			if gotN != wantN {
				t.Errorf(
					"XAxisNorth() = %v,%v: %v; want %v,%v: %v",
					i,
					j,
					gotN,
					i,
					j,
					wantN,
				)
			}
		}
	}
}
