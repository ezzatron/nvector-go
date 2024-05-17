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

func Test_FromRotationMatrix(t *testing.T) {
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
			r := rapidgen.RotationMatrix().Draw(t, "r")

			want, err := client.FromRotationMatrix(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			got := FromRotationMatrix(r)

			if eq, ineq := equality.EqualToVector(got, want, 1e-15); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}

func Test_ToRotationMatrix(t *testing.T) {
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
			type inputs struct {
				V Vector
				F Matrix
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				return inputs{
					V: rapidgen.UnitVector().Draw(t, "nVector"),
					F: rapidgen.RotationMatrix().Draw(t, "coordFrame"),
				}
			}).Filter(func(i inputs) bool {
				// Avoid situations where very close to poles
				// Python implementation rounds to zero in these cases, which causes
				// the Y axis to be [0, 1, 0] instead of the calculated value,
				// producing very different results.
				v := i.V.Transform(i.F)
				yDirNorm := math.Hypot(-v.Z, v.Y)
				if yDirNorm > 0 && yDirNorm <= 1e-100 {
					return false
				}

				return true
			}).Draw(t, "inputs")

			v := i.V
			f := i.F

			want, err := client.ToRotationMatrix(ctx, v, f)
			if err != nil {
				t.Fatal(err)
			}

			got := ToRotationMatrix(v, f)

			if eq, ineq := equality.EqualToMatrix(got, want, 1e-14); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})

	t.Run("it handles the poles", func(t *testing.T) {
		v := Vector{X: 0, Y: 0, Z: 1}
		f := ZAxisNorth
		want := Matrix{
			-1, 0, 0,
			0, 1, -0,
			0, 0, -1,
		}
		got := ToRotationMatrix(v, f)

		if eq, ineq := equality.EqualToMatrix(got, want, 1e-14); !eq {
			equality.ReportInequalities(t, ineq)
		}
	})
}

func Test_RoundTrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		v := rapidgen.UnitVector().Draw(t, "nVector")
		f := rapidgen.RotationMatrix().Draw(t, "coordFrame")

		got := FromRotationMatrix(ToRotationMatrix(v, f))

		if eq, ineq := equality.EqualToVector(got, v, 1e-14); !eq {
			equality.ReportInequalities(t, ineq)
		}
	})
}

func Test_ToRotationMatrixUsingWanderAzimuth(t *testing.T) {
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
			type inputs struct {
				V Vector
				W float64
				F Matrix
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				return inputs{
					V: rapidgen.UnitVector().Draw(t, "nVector"),
					W: rapidgen.Radians().Draw(t, "wanderAzimuth"),
					F: rapidgen.RotationMatrix().Draw(t, "coordFrame"),
				}
			}).Filter(func(i inputs) bool {
				// Avoid situations where components of the xyz2R matrix are close
				// to zero. The Python implementation rounds to zero in these cases,
				// which produces very different results.
				l := ToGeodeticCoordinates(i.V, i.F)
				r := EulerXYZToRotationMatrix(
					EulerXYZ{l.Longitude, -l.Latitude, i.W},
				)
				for _, n := range []float64{
					r.XX, r.XY, r.XZ,
					r.YX, r.YY, r.YZ,
					r.ZX, r.ZY, r.ZZ,
				} {
					if n != 0 && n < 1e-15 && n > -1e-15 {
						return false
					}
				}

				return true
			}).Draw(t, "inputs")

			v := i.V
			w := i.W
			f := i.F

			want, err := client.ToRotationMatrixUsingWanderAzimuth(ctx, v, w, f)
			if err != nil {
				t.Fatal(err)
			}

			got := ToRotationMatrixUsingWanderAzimuth(v, w, f)

			if eq, ineq := equality.EqualToMatrix(got, want, 1e-14); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}
