package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// Delta finds a delta ECEF position vector from a reference n-vector, a depth,
// and a target n-vector.
func (c *Client) Delta(
	ctx context.Context,
	from r3.Vec,
	fromDepth float64,
	to r3.Vec,
	toDepth float64,
	opts ...options.Option,
) (delta r3.Vec, err error) {
	o := options.New(opts)

	return call(
		ctx,
		c,
		unmarshalVector,
		"n_EA_E_and_n_EB_E2p_AB_E",
		map[string]any{
			"n_EA_E": marshalVector(from),
			"z_EA":   fromDepth,
			"n_EB_E": marshalVector(to),
			"z_EB":   toDepth,
			"a":      o.Ellipsoid.SemiMajorAxis,
			"f":      o.Ellipsoid.Flattening,
			"R_Ee":   marshalMatrix(o.CoordFrame),
		},
	)
}

// Destination finds a n-vector from a reference n-vector, a depth, and a delta
// ECEF position vector.
func (c *Client) Destination(
	ctx context.Context,
	from r3.Vec,
	fromDepth float64,
	delta r3.Vec,
	opts ...options.Option,
) (to r3.Vec, toDepth float64, err error) {
	o := options.New(opts)

	type res struct {
		To      [][]float64 `json:"n_EB_E"`
		ToDepth float64     `json:"z_EB"`
	}

	data, err := call(
		ctx,
		c,
		unmarshalAs[res],
		"n_EA_E_and_p_AB_E2n_EB_E",
		map[string]any{
			"n_EA_E": marshalVector(from),
			"z_EA":   fromDepth,
			"p_AB_E": marshalVector(delta),
			"a":      o.Ellipsoid.SemiMajorAxis,
			"f":      o.Ellipsoid.Flattening,
			"R_Ee":   marshalMatrix(o.CoordFrame),
		},
	)
	if err != nil {
		return r3.Vec{}, 0, err
	}

	to = unmarshalVector(data.To)
	toDepth = data.ToDepth

	return to, toDepth, nil
}
