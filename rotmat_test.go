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

func Test_FromRotMat(t *testing.T) {
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

			want, err := client.FromRotMat(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			got := FromRotMat(r)

			if !scalar.EqualWithinAbs(got.X, want.X, 1e-15) {
				t.Errorf("FromRotMat(%v) = X: %v; want X: %v", r, got.X, want.X)
			}
			if !scalar.EqualWithinAbs(got.Y, want.Y, 1e-15) {
				t.Errorf("FromRotMat(%v) = Y: %v; want Y: %v", r, got.Y, want.Y)
			}
			if !scalar.EqualWithinAbs(got.Z, want.Z, 1e-15) {
				t.Errorf("FromRotMat(%v) = Z: %v; want Z: %v", r, got.Z, want.Z)
			}
		})
	})
}

func Test_ToRotMat(t *testing.T) {
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
				Nv   r3.Vec
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				return inputs{
					Opts: rapidgen.Options().Draw(t, "opts"),
					Nv:   rapidgen.UnitVector().Draw(t, "nv"),
				}
			}).Filter(func(i inputs) bool {
				o := options.New(i.Opts)

				// Avoid situations where very close to poles
				// Python implementation rounds to zero in these cases, which causes
				// the Y axis to be [0, 1, 0] instead of the calculated value,
				// producing very different results.
				nvr := o.CoordFrame.MulVec(i.Nv)
				yDirNorm := math.Hypot(-nvr.Z, nvr.Y)
				if yDirNorm > 0 && yDirNorm <= 1e-100 {
					return false
				}

				return true
			}).Draw(t, "inputs")

			nv := i.Nv
			opts := i.Opts

			want, err := client.ToRotMat(ctx, nv, opts...)
			if err != nil {
				t.Fatal(err)
			}

			got := ToRotMat(nv, opts...)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !scalar.EqualWithinAbs(got.At(i, j), want.At(i, j), 1e-14) {
						t.Errorf(
							"ToRotMat(%v, %v) = %v,%v: %v; want %v,%v: %v",
							nv,
							opts,
							i,
							j,
							got.At(i, j),
							i,
							j,
							want.At(i, j),
						)
					}
				}
			}
		})
	})

	t.Run("it handles the poles", func(t *testing.T) {
		nv := r3.Vec{X: 0, Y: 0, Z: 1}
		want := r3.NewMat([]float64{
			-1, 0, 0,
			0, 1, -0,
			0, 0, -1,
		})
		got := ToRotMat(nv)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if !scalar.EqualWithinAbs(got.At(i, j), want.At(i, j), 1e-14) {
					t.Errorf(
						"ToRotMat(%v) = %v,%v: %v; want %v,%v: %v",
						nv,
						i,
						j,
						got.At(i, j),
						i,
						j,
						want.At(i, j),
					)
				}
			}
		}
	})
}

func Test_WithWanderAzimuthToRotMat(t *testing.T) {
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
				Nv   r3.Vec
				Wa   float64
			}

			i := rapid.Custom(func(t *rapid.T) inputs {
				return inputs{
					Opts: rapidgen.Options().Draw(t, "opts"),
					Nv:   rapidgen.UnitVector().Draw(t, "nv"),
					Wa:   rapidgen.Radians().Draw(t, "wa"),
				}
			}).Filter(func(i inputs) bool {
				// Avoid situations where components of the xyz2R matrix are close
				// to zero. The Python implementation rounds to zero in these cases,
				// which produces very different results.
				lat, lon := ToLatLon(i.Nv, i.Opts...)
				r := EulerXYZToRotMat(lon, -lat, i.Wa)
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						n := r.At(i, j)
						if n != 0 && n < 1e-15 && n > -1e-15 {
							return false
						}
					}
				}

				return true
			}).Draw(t, "inputs")

			nv := i.Nv
			wa := i.Wa
			opts := i.Opts

			want, err := client.WithWanderAzimuthToRotMat(ctx, nv, wa, opts...)
			if err != nil {
				t.Fatal(err)
			}

			got := WithWanderAzimuthToRotMat(nv, wa, opts...)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !scalar.EqualWithinAbs(got.At(i, j), want.At(i, j), 1e-14) {
						t.Errorf(
							"WithWanderAzimuthToRotMat(%v, %v, %v) = %v,%v: %v; want %v,%v: %v",
							nv,
							wa,
							opts,
							i,
							j,
							got.At(i, j),
							i,
							j,
							want.At(i, j),
						)
					}
				}
			}
		})
	})
}
