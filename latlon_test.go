package nvector_test

import (
	"context"
	"testing"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/rapidgen"
	"github.com/ezzatron/nvector-go/internal/testapi"
	"gonum.org/v1/gonum/floats/scalar"
	"gonum.org/v1/gonum/mat"
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

			if !mat.EqualApprox(got, want, 1e-14) {
				t.Errorf(
					"FromLatLon(%v, %v) = %v; want %v",
					latitude,
					longitude,
					got,
					want,
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

			latEqual := scalar.EqualWithinAbsOrRel(gotLat, wantLat, 1e-14, 1e-14)
			lonEqual := scalar.EqualWithinAbsOrRel(gotLon, wantLon, 1e-14, 1e-14)

			if !latEqual || !lonEqual {
				t.Errorf(
					"ToLatLon(%v) = (%v, %v); want (%v, %v)",
					nv,
					gotLat,
					gotLon,
					wantLat,
					wantLon,
				)
			}
		})
	})
}
