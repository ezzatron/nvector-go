package nvector_test

import (
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
)

func Test_ZAxisNorth(t *testing.T) {
	want := Matrix{
		0, 0, 1,
		0, 1, 0,
		-1, 0, 0,
	}

	got := ZAxisNorth

	if eq, ineq := equality.EqualToMatrix(got, want, 1e-64); !eq {
		equality.ReportInequalities(t, ineq)
	}
}

func Test_XAxisNorth(t *testing.T) {
	want := Matrix{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	got := XAxisNorth

	if eq, ineq := equality.EqualToMatrix(got, want, 1e-64); !eq {
		equality.ReportInequalities(t, ineq)
	}
}
