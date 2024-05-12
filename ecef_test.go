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

func Test_FromECEF(t *testing.T) {
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
				Opts []Option
				ECEF r3.Vec
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				opts := rapidgen.Options().Draw(t, "opts")
				o := options.New(opts)

				return inputs{
					Opts: opts,
					ECEF: rapidgen.EcefVector(o.Ellipsoid).Draw(t, "ecef"),
				}
			}).Filter(func(i inputs) bool {
				o := options.New(i.Opts)
				a := o.Ellipsoid.SemiMajorAxis
				f := o.Ellipsoid.Flattening

				ecefr := o.CoordFrame.MulVec(i.ECEF)

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

			p := i.ECEF
			opts := i.Opts

			wantNv, wantD, err := client.FromECEF(ctx, p, opts...)
			if err != nil {
				t.Fatal(err)
			}

			gotNv, gotD := FromECEF(p, opts...)

			if !scalar.EqualWithinAbs(gotNv.X, wantNv.X, 1e-14) {
				t.Errorf(
					"FromECEF(%v) = X: %v; want X: %v",
					p,
					gotNv.X,
					wantNv.X,
				)
			}
			if !scalar.EqualWithinAbs(gotNv.Y, wantNv.Y, 1e-14) {
				t.Errorf(
					"FromECEF(%v) = Y: %v; want Y: %v",
					p,
					gotNv.Y,
					wantNv.Y,
				)
			}
			if !scalar.EqualWithinAbs(gotNv.Z, wantNv.Z, 1e-14) {
				t.Errorf(
					"FromECEF(%v) = Z: %v; want Z: %v",
					p,
					gotNv.Z,
					wantNv.Z,
				)
			}
			if !scalar.EqualWithinAbs(gotD, wantD, 1e-8) {
				t.Errorf(
					"FromECEF(%v) = D: %v; want D: %v",
					p,
					gotD,
					wantD,
				)
			}
		})
	})
}

func Test_ToECEF(t *testing.T) {
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

			nv := rapidgen.UnitVector().Draw(t, "nv")
			d := rapidgen.Depth(o.Ellipsoid).Draw(t, "d")

			want, err := client.ToECEF(ctx, nv, d, opts...)
			if err != nil {
				t.Fatal(err)
			}

			got := ToECEF(nv, d, opts...)

			if !scalar.EqualWithinAbs(got.X, want.X, 1e-7) {
				t.Errorf(
					"ToECEF(%v, %v) = X: %v; want X: %v",
					nv,
					d,
					got.X,
					want.X,
				)
			}
			if !scalar.EqualWithinAbs(got.Y, want.Y, 1e-7) {
				t.Errorf(
					"ToECEF(%v, %v) = Y: %v; want Y: %v",
					nv,
					d,
					got.Y,
					want.Y,
				)
			}
			if !scalar.EqualWithinAbs(got.Z, want.Z, 1e-7) {
				t.Errorf(
					"ToECEF(%v, %v) = Z: %v; want Z: %v",
					nv,
					d,
					got.Z,
					want.Z,
				)
			}
		})
	})
}
