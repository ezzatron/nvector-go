package nvector_test

import (
	"context"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"pgregory.net/rapid"
)

func Test_FromGeodeticCoordinates(t *testing.T) {
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
			c := rapidgen.GeodeticCoordinates().Draw(t, "coords")
			f := rapidgen.RotationMatrix().Draw(t, "coordFrame")

			want, err := client.FromGeodeticCoordinates(ctx, c, f)
			if err != nil {
				t.Fatal(err)
			}

			got := FromGeodeticCoordinates(c, f)

			if eq, ineq := equality.EqualToVector(got, want, 1e-14); !eq {
				equality.ReportInequalities(t, ineq)
			}
		})
	})
}

func Test_ToGeodeticCoordinates(t *testing.T) {
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
			v := rapidgen.UnitVector().Draw(t, "nVector")
			f := rapidgen.RotationMatrix().Draw(t, "coordFrame")

			want, err := client.ToGeodeticCoordinates(ctx, v, f)
			if err != nil {
				t.Fatal(err)
			}

			got := ToGeodeticCoordinates(v, f)

			latEq, latIneq := equality.EqualToRadians(
				got.Latitude,
				want.Latitude,
				1e-14,
			)
			lonEq, lonIneq := equality.EqualToRadians(
				got.Longitude,
				want.Longitude,
				1e-14,
			)

			if !latEq {
				equality.ReportInequality(t, "latitude", latIneq)
			}
			if !lonEq {
				equality.ReportInequality(t, "longitude", lonIneq)
			}
		})
	})
}
