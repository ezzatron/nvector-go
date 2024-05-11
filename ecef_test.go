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
	"gonum.org/v1/gonum/mat"
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
				ECEF mat.Vector
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				opts := rapidgen.Options().Draw(t, "opts")
				o := options.New(opts)

				ecef := rapidgen.EcefVector(o.Ellipsoid).Draw(t, "ecef")

				return inputs{
					Opts: opts,
					ECEF: ecef,
				}
			}).Filter(func(i inputs) bool {
				o := options.New(i.Opts)
				a := o.Ellipsoid.SemiMajorAxis
				f := o.Ellipsoid.Flattening

				ecefr := mat.NewVecDense(3, nil)
				ecefr.MulVec(o.CoordFrame, i.ECEF)

				// filter vectors where the x or yz components are zero after rotation
				// this causes a division by zero in the Python implementation
				if ecefr.AtVec(0) == 0 || ecefr.AtVec(1)+ecefr.AtVec(2) == 0 {
					return false
				}

				// filter a case that makes the Python implementation try to find the
				// square root of a negative number
				// not sure why this happens, the math is beyond me
				e2 := 2*f - math.Pow(f, 2)
				R2 := math.Pow(ecefr.AtVec(1), 2) + math.Pow(ecefr.AtVec(2), 2)
				p := R2 / math.Pow(a, 2)
				q := (1 - e2) / math.Pow(a, 2) * math.Pow(ecefr.AtVec(0), 2)
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

			nvEqual := mat.EqualApprox(gotNv, wantNv, 1e-15)
			dEqual := scalar.EqualWithinAbs(gotD, wantD, 1e-8)

			if !nvEqual || !dEqual {
				t.Errorf(
					"FromECEF(%v) = %v, %v; want %v, %v, (%v, %v)",
					p,
					gotNv,
					gotD,
					wantNv,
					wantD,
					nvEqual,
					dEqual,
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

			if !mat.EqualApprox(got, want, 1e-7) {
				t.Errorf(
					"ToECEF(%v, %v) = %v; want %v",
					nv,
					d,
					got,
					want,
				)
			}
		})
	})
}
