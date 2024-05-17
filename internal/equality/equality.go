package equality

import (
	"math"

	"github.com/ezzatron/nvector-go"
)

// EqualToEulerAnglesXYZ returns a boolean indicating whether two sets of Euler
// angles in XYZ order are equal within a tolerance, and the inequality for each
// angle.
func EqualToEulerAnglesXYZ(
	a, b nvector.EulerXYZ,
	tol float64,
) (bool, map[string]Float64Inequality) {
	ineq := make(map[string]Float64Inequality, 3)

	if eq, i := EqualToRadians(a.X, b.X, tol); !eq {
		ineq["X"] = i
	}
	if eq, i := EqualToRadians(a.Y, b.Y, tol); !eq {
		ineq["Y"] = i
	}
	if eq, i := EqualToRadians(a.Z, b.Z, tol); !eq {
		ineq["Z"] = i
	}

	return len(ineq) == 0, ineq
}

// EqualToEulerAnglesZYX returns a boolean indicating whether two sets of Euler
// angles in ZYX order are equal within a tolerance, and the inequality for each
// angle.
func EqualToEulerAnglesZYX(
	a, b nvector.EulerZYX,
	tol float64,
) (bool, map[string]Float64Inequality) {
	ineq := make(map[string]Float64Inequality, 3)

	if eq, i := EqualToRadians(a.Z, b.Z, tol); !eq {
		ineq["Z"] = i
	}
	if eq, i := EqualToRadians(a.Y, b.Y, tol); !eq {
		ineq["Y"] = i
	}
	if eq, i := EqualToRadians(a.X, b.X, tol); !eq {
		ineq["X"] = i
	}

	return len(ineq) == 0, ineq
}

// EqualToFloat64 returns a boolean indicating whether two float64s are equal
// within a tolerance, and the inequality.
func EqualToFloat64(a, b, tol float64) (bool, Float64Inequality) {
	diff := math.Abs(b - a)

	return diff <= tol, Float64Inequality{diff, tol, a, b}
}

// EqualToMatrix returns a boolean indicating whether two matrices are equal
// within a tolerance, and the inequality for each component.
func EqualToMatrix(
	a, b nvector.Matrix,
	tol float64,
) (bool, map[string]Float64Inequality) {
	ineq := make(map[string]Float64Inequality, 9)

	if eq, i := EqualToFloat64(a.XX, b.XX, tol); !eq {
		ineq["XX"] = i
	}
	if eq, i := EqualToFloat64(a.XY, b.XY, tol); !eq {
		ineq["XY"] = i
	}
	if eq, i := EqualToFloat64(a.XZ, b.XZ, tol); !eq {
		ineq["XZ"] = i
	}
	if eq, i := EqualToFloat64(a.YX, b.YX, tol); !eq {
		ineq["YX"] = i
	}
	if eq, i := EqualToFloat64(a.YY, b.YY, tol); !eq {
		ineq["YY"] = i
	}
	if eq, i := EqualToFloat64(a.YZ, b.YZ, tol); !eq {
		ineq["YZ"] = i
	}
	if eq, i := EqualToFloat64(a.ZX, b.ZX, tol); !eq {
		ineq["ZX"] = i
	}
	if eq, i := EqualToFloat64(a.ZY, b.ZY, tol); !eq {
		ineq["ZY"] = i
	}
	if eq, i := EqualToFloat64(a.ZZ, b.ZZ, tol); !eq {
		ineq["ZZ"] = i
	}

	return len(ineq) == 0, ineq
}

// EqualToRadians returns a boolean indicating whether two angles are equal
// within a tolerance, and the inequality.
func EqualToRadians(a, b, tol float64) (bool, Float64Inequality) {
	diff := math.Pi - math.Abs(math.Abs(a-b)-math.Pi)

	return diff <= tol, Float64Inequality{diff, tol, a, b}
}

// EqualToVector returns a boolean indicating whether two vectors are equal
// within a tolerance, and the inequality for each component.
func EqualToVector(
	a, b nvector.Vector,
	tol float64,
) (bool, map[string]Float64Inequality) {
	ineq := make(map[string]Float64Inequality, 3)

	if eq, i := EqualToFloat64(a.X, b.X, tol); !eq {
		ineq["X"] = i
	}
	if eq, i := EqualToFloat64(a.Y, b.Y, tol); !eq {
		ineq["Y"] = i
	}
	if eq, i := EqualToFloat64(a.Z, b.Z, tol); !eq {
		ineq["Z"] = i
	}

	return len(ineq) == 0, ineq
}

// EqualToVectorWithDepth returns a boolean indicating whether two vectors with
// depth are equal within a tolerance, and the inequality for each component.
func EqualToVectorWithDepth(
	a, b nvector.Position,
	vTol, dTol float64,
) (bool, map[string]Float64Inequality) {
	ineq := make(map[string]Float64Inequality, 4)

	if eq, i := EqualToFloat64(a.Vector.X, b.Vector.X, vTol); !eq {
		ineq["Vector.X"] = i
	}
	if eq, i := EqualToFloat64(a.Vector.Y, b.Vector.Y, vTol); !eq {
		ineq["Vector.Y"] = i
	}
	if eq, i := EqualToFloat64(a.Vector.Z, b.Vector.Z, vTol); !eq {
		ineq["Vector.Z"] = i
	}
	if eq, i := EqualToFloat64(a.Depth, b.Depth, dTol); !eq {
		ineq["Depth"] = i
	}

	return len(ineq) == 0, ineq
}

// Float64Inequality is an inequality between two float64s.
type Float64Inequality struct {
	// Diff is the absolute difference between the two components.
	Diff float64
	// Tolerance is the tolerance for the difference.
	Tolerance float64
	// Got is the actual component's value.
	Got float64
	// Want is the expected component's value.
	Want float64
}

// ReportInequality fails the test with a formatted message of the inequality.
//
// p is a prefix for the message. c is the component being compared.
func ReportInequality(t TestErrorReporter, c string, i Float64Inequality) {
	t.Errorf(
		"got %v component %v; want %v (got difference %v; want <=%v)",
		c,
		i.Got,
		i.Want,
		i.Diff,
		i.Tolerance,
	)
}

// ReportInequalities fails the test with a formatted message of the
// inequalities.
//
// p is a prefix for the message. ineq is a map of components and their
// inequalities.
func ReportInequalities(
	t TestErrorReporter,
	ineq map[string]Float64Inequality,
) {
	for c, i := range ineq {
		ReportInequality(t, c, i)
	}
}

// TestErrorReporter is an interface for reporting testing errors.
type TestErrorReporter interface {
	Errorf(format string, args ...interface{})
}
