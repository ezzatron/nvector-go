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
