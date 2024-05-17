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

func Test_Delta(t *testing.T) {
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
			e := rapidgen.Ellipsoid().Draw(t, "ellipsoid")

			from := Position{
				Vector: rapidgen.UnitVector().Draw(t, "fromNVector"),
				Depth:  rapidgen.Depth(e).Draw(t, "fromDepth"),
			}
			to := Position{
				Vector: rapidgen.UnitVector().Draw(t, "toNVector"),
				Depth:  rapidgen.Depth(e).Draw(t, "toDepth"),
			}
			f := rapidgen.RotationMatrix().Draw(t, "coordFrame")

			want, err := client.Delta(ctx, from, to, e, f)
			if err != nil {
				t.Fatal(err)
			}

			got := Delta(from, to, e, f)

			if eq, ineq := equality.EqualToVector(got, want, 1e-7); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}

func Test_Destination(t *testing.T) {
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
				From  Position
				Delta Vector
				E     Ellipsoid
				F     Matrix
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				e := rapidgen.Ellipsoid().Draw(t, "ellipsoid")

				return inputs{
					From: Position{
						Vector: rapidgen.UnitVector().Draw(t, "fromNVector"),
						Depth:  rapidgen.Depth(e).Draw(t, "fromDepth"),
					},
					Delta: rapidgen.EcefVector(e).Draw(t, "delta"),
					E:     e,
					F:     rapidgen.RotationMatrix().Draw(t, "coordFrame"),
				}
			}).Filter(func(i inputs) bool {
				a := i.E.SemiMajorAxis
				f := i.E.Flattening

				v := ToECEF(i.From, i.E, i.F).Add(i.Delta).Transform(i.F)

				// filter vectors where the x or yz components are zero after rotation
				// this causes a division by zero in the Python implementation
				if v.X == 0 || v.Y+v.Z == 0 {
					return false
				}

				// filter a case that makes the Python implementation try to find the
				// square root of a negative number
				// not sure why this happens, the math is beyond me
				e2 := 2*f - math.Pow(f, 2)
				R2 := math.Pow(v.Y, 2) + math.Pow(v.Z, 2)
				p := R2 / math.Pow(a, 2)
				q := (1 - e2) / math.Pow(a, 2) * math.Pow(v.X, 2)
				r := (p + q - math.Pow(e2, 2)) / 6
				s := math.Pow(e2, 2) * p * q / (4 * math.Pow(r, 3))
				if math.IsNaN(s) || s <= 0 {
					return false
				}

				return true
			}).Draw(t, "inputs")

			from := i.From
			delta := i.Delta
			e := i.E
			f := i.F

			want, err := client.Destination(ctx, from, delta, e, f)
			if err != nil {
				t.Fatal(err)
			}

			got := Destination(from, delta, e, f)

			eq, ineq := equality.EqualToVectorWithDepth(got, want, 1e-12, 1e-7)

			if !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}
