package nvector_test

import (
	"context"
	"math"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"gonum.org/v1/gonum/floats/scalar"
	"gonum.org/v1/gonum/spatial/r3"
	"pgregory.net/rapid"
)

func Test_EulerXYZToRotMat(t *testing.T) {
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
			x := rapidgen.Radians().Draw(t, "x")
			y := rapidgen.Radians().Draw(t, "y")
			z := rapidgen.Radians().Draw(t, "z")

			want, err := client.EulerXYZToRotMat(ctx, x, y, z)
			if err != nil {
				t.Fatal(err)
			}

			got := EulerXYZToRotMat(x, y, z)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !scalar.EqualWithinAbs(got.At(i, j), want.At(i, j), 1e-15) {
						t.Errorf(
							"EulerXYZToRotMat(%v, %v, %v) = %v,%v: %v; want %v,%v: %v",
							x,
							y,
							z,
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

func Test_EulerZYXToRotMat(t *testing.T) {
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
			z := rapidgen.Radians().Draw(t, "z")
			y := rapidgen.Radians().Draw(t, "y")
			x := rapidgen.Radians().Draw(t, "x")

			want, err := client.EulerZYXToRotMat(ctx, z, y, x)
			if err != nil {
				t.Fatal(err)
			}

			got := EulerZYXToRotMat(z, y, x)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !scalar.EqualWithinAbs(got.At(i, j), want.At(i, j), 1e-15) {
						t.Errorf(
							"EulerZYXToRotMat(%v, %v, %v) = %v,%v: %v; want %v,%v: %v",
							z,
							y,
							x,
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

func Test_RotMatToEulerXYZ(t *testing.T) {
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

			wantX, wantY, wantZ, err := client.RotMatToEulerXYZ(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			gotX, gotY, gotZ := RotMatToEulerXYZ(r)

			if !equality.EqualToRadians(gotX, wantX, 1e-15) {
				t.Errorf("RotMatToEulerXYZ(%v) = X: %v; want X: %v", r, gotX, wantX)
			}
			if !equality.EqualToRadians(gotY, wantY, 1e-15) {
				t.Errorf("RotMatToEulerXYZ(%v) = Y: %v; want Y: %v", r, gotY, wantY)
			}
			if !equality.EqualToRadians(gotZ, wantZ, 1e-15) {
				t.Errorf("RotMatToEulerXYZ(%v) = Z: %v; want Z: %v", r, gotZ, wantZ)
			}
		})
	})

	t.Run("handles Euler angle singularity", func(t *testing.T) {
		r := r3.NewMat([]float64{
			0, 0, 1,
			0, 1, 0,
			1, 0, 0,
		})

		wantX, wantY, wantZ := 0.0, math.Pi/2, 0.0
		gotX, gotY, gotZ := RotMatToEulerXYZ(r)

		if !equality.EqualToRadians(gotX, wantX, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = X: %v; want X: %v", r, gotX, wantX)
		}
		if !equality.EqualToRadians(gotY, wantY, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = Y: %v; want Y: %v", r, gotY, wantY)
		}
		if !equality.EqualToRadians(gotZ, wantZ, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = Z: %v; want Z: %v", r, gotZ, wantZ)
		}
	})
}

func Test_RotMatToEulerZYX(t *testing.T) {
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

			wantZ, wantY, wantX, err := client.RotMatToEulerZYX(ctx, r)
			if err != nil {
				t.Fatal(err)
			}

			gotZ, gotY, gotX := RotMatToEulerZYX(r)

			if !equality.EqualToRadians(gotZ, wantZ, 1e-15) {
				t.Errorf("RotMatToEulerZYX(%v) = Z: %v; want Z: %v", r, gotZ, wantZ)
			}
			if !equality.EqualToRadians(gotY, wantY, 1e-15) {
				t.Errorf("RotMatToEulerZYX(%v) = Y: %v; want Y: %v", r, gotY, wantY)
			}
			if !equality.EqualToRadians(gotX, wantX, 1e-15) {
				t.Errorf("RotMatToEulerZYX(%v) = X: %v; want X: %v", r, gotX, wantX)
			}
		})
	})

	t.Run("handles Euler angle singularity", func(t *testing.T) {
		r := r3.NewMat([]float64{
			0, 0, 1,
			0, 1, 0,
			1, 0, 0,
		})

		wantX, wantY, wantZ := -0.0, math.Pi/2, 0.0
		gotX, gotY, gotZ := RotMatToEulerXYZ(r)

		if !equality.EqualToRadians(gotX, wantX, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = X: %v; want X: %v", r, gotX, wantX)
		}
		if !equality.EqualToRadians(gotY, wantY, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = Y: %v; want Y: %v", r, gotY, wantY)
		}
		if !equality.EqualToRadians(gotZ, wantZ, 1e-15) {
			t.Errorf("RotMatToEulerXYZ(%v) = Z: %v; want Z: %v", r, gotZ, wantZ)
		}
	})
}
