package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/mat"
)

// FromLatLon converts a geodetic latitude and longitude to an n-vector.
func (c *Client) FromLatLon(
	ctx context.Context,
	lat, lon float64,
	opts ...options.Option,
) (mat.Vector, error) {
	o := options.New(opts)

	return call(ctx, c, unmarshalVector, "lat_lon2n_E", map[string]any{
		"latitude":  lat,
		"longitude": lon,
		"R_Ee":      marshalMatrix(o.CoordFrame),
	})
}

// ToLatLon converts an n-vector to a geodetic latitude and longitude.
func (c *Client) ToLatLon(
	ctx context.Context,
	nv mat.Vector,
	opts ...options.Option,
) (float64, float64, error) {
	o := options.New(opts)

	type latLon struct {
		Lat float64 `json:"latitude"`
		Lon float64 `json:"longitude"`
	}

	r, err := call(ctx, c, unmarshalAs[latLon], "n_E2lat_lon", map[string]any{
		"n_E":  marshalVector(nv),
		"R_Ee": marshalMatrix(o.CoordFrame),
	})
	if err != nil {
		return 0, 0, err
	}

	return r.Lat, r.Lon, nil
}
