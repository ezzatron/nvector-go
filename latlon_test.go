package nvector_test

import (
	"context"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/equality"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"gonum.org/v1/gonum/floats/scalar"
	"pgregory.net/rapid"
)

func Test_FromLatLon(t *testing.T) {
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
			latitude := rapidgen.Latitude().Draw(t, "latitude")
			longitude := rapidgen.Longitude().Draw(t, "longitude")
			opts := rapidgen.Options().Draw(t, "opts")

			want, err := client.FromLatLon(ctx, latitude, longitude, opts...)
			if err != nil {
				t.Fatal(err)
			}

			got := FromLatLon(latitude, longitude, opts...)

			if !scalar.EqualWithinAbs(got.X, want.X, 1e-14) {
				t.Errorf(
					"FromLatLon(%v, %v) = X: %v; want X: %v",
					latitude,
					longitude,
					got.X,
					want.X,
				)
			}
			if !scalar.EqualWithinAbs(got.Y, want.Y, 1e-14) {
				t.Errorf(
					"FromLatLon(%v, %v) = Y: %v; want Y: %v",
					latitude,
					longitude,
					got.Y,
					want.Y,
				)
			}
			if !scalar.EqualWithinAbs(got.Z, want.Z, 1e-14) {
				t.Errorf(
					"FromLatLon(%v, %v) = Z: %v; want Z: %v",
					latitude,
					longitude,
					got.Z,
					want.Z,
				)
			}
		})
	})
}

func Test_ToLatLon(t *testing.T) {
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
			nv := rapidgen.UnitVector().Draw(t, "nv")
			opts := rapidgen.Options().Draw(t, "opts")

			wantLat, wantLon, err := client.ToLatLon(ctx, nv, opts...)
			if err != nil {
				t.Fatal(err)
			}

			gotLat, gotLon := ToLatLon(nv, opts...)

			if !equality.EqualToRadians(gotLat, wantLat, 1e-14) {
				t.Errorf(
					"ToLatLon(%v) = lat: %v; want lat: %v",
					nv,
					gotLat,
					wantLat,
				)
			}
			if !equality.EqualToRadians(gotLon, wantLon, 1e-14) {
				t.Errorf(
					"ToLatLon(%v) = lon: %v; want lon: %v",
					nv,
					gotLon,
					wantLon,
				)
			}
		})
	})
}
