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

func Test_EulerXYZToRotationMatrix(t *testing.T) {
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

			want, err := client.EulerXYZToRotationMatrix(ctx, a)
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

func Test_EulerZYXToRotationMatrix(t *testing.T) {
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

			want, err := client.EulerZYXToRotationMatrix(ctx, a)
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

func Test_RotationMatrixToEulerXYZ(t *testing.T) {
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
				return cy > (math.Nextafter(1, 2)-1)*10
			}).Draw(t, "r")

			want, err := client.RotationMatrixToEulerXYZ(ctx, r)
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
		tests := []struct {
			name string
			y    float64
		}{
			{"positive y", math.Pi / 2},
			{"negative y", -math.Pi / 2},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := EulerXYZToRotationMatrix(EulerXYZ{X: 0, Y: tt.y, Z: 0.3})

				want := EulerXYZ{0, tt.y, 0.3}
				got := RotationMatrixToEulerXYZ(r)

				if eq, ineq := equality.EqualToEulerAnglesXYZ(got, want, 1e-15); !eq {
					equality.ReportInequalities(t, ineq)
				}
			})
		}
	})
}

func Test_RotationMatrixToEulerZYX(t *testing.T) {
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
				return cy > (math.Nextafter(1, 2)-1)*10
			}).Draw(t, "r")

			want, err := client.RotationMatrixToEulerZYX(ctx, r)
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
		tests := []struct {
			name string
			y    float64
		}{
			{"positive y", math.Pi / 2},
			{"negative y", -math.Pi / 2},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := EulerZYXToRotationMatrix(EulerZYX{Z: 0.3, Y: tt.y, X: 0})

				want := EulerZYX{0.3, tt.y, 0}
				got := RotationMatrixToEulerZYX(r)

				if eq, ineq := equality.EqualToEulerAnglesZYX(got, want, 1e-15); !eq {
					equality.ReportInequalities(t, ineq)
				}
			})
		}
	})
}
