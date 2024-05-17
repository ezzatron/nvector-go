package nvector_test

import (
	"context"
	"math"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"pgregory.net/rapid"
)

func Test_RotationMatrixToXYZ(t *testing.T) {
	client, err := testapi.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		client.Close()
	})

	ctx := context.Background()

	t.Run("it matches the reference implementation", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			r := rapidgen.RotationMatrix().Filter(func(r Matrix) bool {
				cy := math.Sqrt((math.Pow(r.XX, 2) +
					math.Pow(r.XY, 2) +
					math.Pow(r.YZ, 2) +
					math.Pow(r.ZZ, 2)) / 2,
				)

				return cy > 1e-14
			}).Draw(t, "r")

			want, err := client.RotationMatrixToXYZ(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			got := RotationMatrixToEulerXYZ(r)

			if eq, ineq := equality.EqualToEulerAnglesXYZ(got, want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})

	t.Run("handles Euler angle singularity", func(t *testing.T) {
		r := Matrix{
			0, 0, 1,
			0, 1, 0,
			1, 0, 0,
		}

		want := EulerXYZ{0.0, math.Pi / 2, 0.0}
		got := RotationMatrixToEulerXYZ(r)

		if eq, ineq := equality.EqualToEulerAnglesXYZ(got, want, 1e-15); !eq {
			equality.ReportInequalities(t, ineq)
		}
	})
}

func Test_RotationMatrixToZYX(t *testing.T) {
	client, err := testapi.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		client.Close()
	})

	ctx := context.Background()

	t.Run("it matches the reference implementation", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			r := rapidgen.RotationMatrix().Filter(func(r Matrix) bool {
				cy := math.Sqrt((math.Pow(r.XX, 2) +
					math.Pow(r.YX, 2) +
					math.Pow(r.ZY, 2) +
					math.Pow(r.ZZ, 2)) / 2,
				)

				return cy > 1e-14
			}).Draw(t, "r")

			want, err := client.RotationMatrixToZYX(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			got := RotationMatrixToEulerZYX(r)

			if eq, ineq := equality.EqualToEulerAnglesZYX(got, want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})

	t.Run("handles Euler angle singularity", func(t *testing.T) {
		r := Matrix{
			0, 0, 1,
			0, 1, 0,
			1, 0, 0,
		}

		want := EulerZYX{0.0, math.Pi / 2, -0.0}
		got := RotationMatrixToEulerZYX(r)

		if eq, ineq := equality.EqualToEulerAnglesZYX(got, want, 1e-15); !eq {
			equality.ReportInequalities(t, ineq)
		}
	})
}

func Test_XYZToRotationMatrix(t *testing.T) {
	client, err := testapi.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		client.Close()
	})

	ctx := context.Background()

	t.Run("it matches the reference implementation", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			a := EulerXYZ{
				rapidgen.Radians().Draw(t, "x"),
				rapidgen.Radians().Draw(t, "y"),
				rapidgen.Radians().Draw(t, "z"),
			}

			want, err := client.XYZToRotationMatrix(ctx, a)
			if err != nil {
				t.Fatal(err)
			}

			got := EulerXYZToRotationMatrix(a)

			if eq, ineq := equality.EqualToMatrix(got, want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}

func Test_ZYXToRotationMatrix(t *testing.T) {
	client, err := testapi.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		client.Close()
	})

	ctx := context.Background()

	t.Run("it matches the reference implementation", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			a := EulerZYX{
				rapidgen.Radians().Draw(t, "x"),
				rapidgen.Radians().Draw(t, "y"),
				rapidgen.Radians().Draw(t, "z"),
			}

			want, err := client.ZYXToRotationMatrix(ctx, a)
			if err != nil {
				t.Fatal(err)
			}

			got := EulerZYXToRotationMatrix(a)

			if eq, ineq := equality.EqualToMatrix(got, want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}
