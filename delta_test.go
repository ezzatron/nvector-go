package nvector_test

import (
	"context"
	"math"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/options"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"gonum.org/v1/gonum/floats/scalar"
	"gonum.org/v1/gonum/spatial/r3"
	"pgregory.net/rapid"
)

func Test_FromDelta(t *testing.T) {
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
				Opts      []Option
				From      r3.Vec
				FromDepth float64
				Delta     r3.Vec
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				opts := rapidgen.Options().Draw(t, "opts")
				o := options.New(opts)

				return inputs{
					Opts:      opts,
					From:      rapidgen.UnitVector().Draw(t, "from"),
					FromDepth: rapidgen.Depth(o.Ellipsoid).Draw(t, "fromDepth"),
					Delta:     rapidgen.EcefVector(o.Ellipsoid).Draw(t, "delta"),
				}
			}).Filter(func(i inputs) bool {
				o := options.New(i.Opts)
				a := o.Ellipsoid.SemiMajorAxis
				f := o.Ellipsoid.Flattening

				ecefr := o.CoordFrame.MulVec(
					r3.Add(ToECEF(i.From, i.FromDepth, i.Opts...), i.Delta),
				)

				// filter vectors where the x or yz components are zero after rotation
				// this causes a division by zero in the Python implementation
				if ecefr.X == 0 || ecefr.Y+ecefr.Z == 0 {
					return false
				}

				// filter a case that makes the Python implementation try to find the
				// square root of a negative number
				// not sure why this happens, the math is beyond me
				e2 := 2*f - math.Pow(f, 2)
				R2 := math.Pow(ecefr.Y, 2) + math.Pow(ecefr.Z, 2)
				p := R2 / math.Pow(a, 2)
				q := (1 - e2) / math.Pow(a, 2) * math.Pow(ecefr.X, 2)
				r := (p + q - math.Pow(e2, 2)) / 6
				s := math.Pow(e2, 2) * p * q / (4 * math.Pow(r, 3))
				if math.IsNaN(s) || s <= 0 {
					return false
				}

				return true
			}).Draw(t, "inputs")

			from := i.From
			fromDepth := i.FromDepth
			delta := i.Delta
			opts := i.Opts

			wantNv, wantD, err := client.FromDelta(
				ctx,
				from,
				fromDepth,
				delta,
				opts...,
			)
			if err != nil {
				t.Fatal(err)
			}

			gotNv, gotD := FromDelta(from, fromDepth, delta, opts...)

			if !scalar.EqualWithinAbs(gotNv.X, wantNv.X, 1e-12) {
				t.Errorf(
					"FromDelta(%v, %v, %v) = X: %v; want X: %v",
					from,
					fromDepth,
					delta,
					gotNv.X,
					wantNv.X,
				)
			}
			if !scalar.EqualWithinAbs(gotNv.Y, wantNv.Y, 1e-12) {
				t.Errorf(
					"FromDelta(%v, %v, %v) = Y: %v; want Y: %v",
					from,
					fromDepth,
					delta,
					gotNv.Y,
					wantNv.Y,
				)
			}
			if !scalar.EqualWithinAbs(gotNv.Z, wantNv.Z, 1e-12) {
				t.Errorf(
					"FromDelta(%v, %v, %v) = Z: %v; want Z: %v",
					from,
					fromDepth,
					delta,
					gotNv.Z,
					wantNv.Z,
				)
			}
			if !scalar.EqualWithinAbs(gotD, wantD, 1e-8) {
				t.Errorf(
					"FromDelta(%v, %v, %v) = D: %v; want D: %v",
					from,
					fromDepth,
					delta,
					gotD,
					wantD,
				)
			}
		})
	})
}

func Test_ToDelta(t *testing.T) {
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
			opts := rapidgen.Options().Draw(t, "opts")
			o := options.New(opts)

			from := rapidgen.UnitVector().Draw(t, "from")
			fromDepth := rapidgen.Depth(o.Ellipsoid).Draw(t, "fromDepth")
			to := rapidgen.UnitVector().Draw(t, "to")
			toDepth := rapidgen.Depth(o.Ellipsoid).Draw(t, "toDepth")

			want, err := client.ToDelta(
				ctx,
				from,
				fromDepth,
				to,
				toDepth,
				opts...,
			)
			if err != nil {
				t.Fatal(err)
			}

			got := ToDelta(from, fromDepth, to, toDepth, opts...)

			if !scalar.EqualWithinAbs(got.X, want.X, 1e-7) {
				t.Errorf(
					"ToDelta(%v, %v, %v, %v) = X: %v; want X: %v",
					from,
					fromDepth,
					to,
					toDepth,
					got.X,
					want.X,
				)
			}
			if !scalar.EqualWithinAbs(got.Y, want.Y, 1e-7) {
				t.Errorf(
					"ToDelta(%v, %v, %v, %v) = Y: %v; want Y: %v",
					from,
					fromDepth,
					to,
					toDepth,
					got.Y,
					want.Y,
				)
			}
			if !scalar.EqualWithinAbs(got.Z, want.Z, 1e-7) {
				t.Errorf(
					"ToDelta(%v, %v, %v, %v) = Z: %v; want Z: %v",
					from,
					fromDepth,
					to,
					toDepth,
					got.Z,
					want.Z,
				)
			}
		})
	})
}
