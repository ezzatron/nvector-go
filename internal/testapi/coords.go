package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// FromGeodeticCoordinates converts geodetic coordinates to an n-vector.
func (c *Client) FromGeodeticCoordinates(
	ctx context.Context,
	gc nvector.GeodeticCoordinates,
	f nvector.Matrix,
) (nvector.Vector, error) {
	return call(ctx, c, unmarshalVector, "lat_lon2n_E", map[string]any{
		"latitude":  gc.Latitude,
		"longitude": gc.Longitude,
		"R_Ee":      marshalMatrix(f),
	})
}

// ToGeodeticCoordinates converts an n-vector to geodetic coordinates.
func (c *Client) ToGeodeticCoordinates(
	ctx context.Context,
	v nvector.Vector,
	f nvector.Matrix,
) (nvector.GeodeticCoordinates, error) {
	type res struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "n_E2lat_lon", map[string]any{
		"n_E":  marshalVector(v),
		"R_Ee": marshalMatrix(f),
	})
	if err != nil {
		return nvector.GeodeticCoordinates{}, err
	}

	return nvector.GeodeticCoordinates(data), nil
}
