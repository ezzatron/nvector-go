package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// Delta finds a delta ECEF position vector from a reference n-vector, a depth,
// and a target n-vector.
func (c *Client) Delta(
	ctx context.Context,
	from nvector.Position,
	to nvector.Position,
	e nvector.Ellipsoid,
	f nvector.Matrix,
) (nvector.Vector, error) {
	return call(
		ctx,
		c,
		unmarshalVector,
		"n_EA_E_and_n_EB_E2p_AB_E",
		map[string]any{
			"n_EA_E": marshalVector(from.Vector),
			"z_EA":   from.Depth,
			"n_EB_E": marshalVector(to.Vector),
			"z_EB":   to.Depth,
			"a":      e.SemiMajorAxis,
			"f":      e.Flattening,
			"R_Ee":   marshalMatrix(f),
		},
	)
}

// Destination finds a n-vector from a reference n-vector, a depth, and a delta
// ECEF position vector.
func (c *Client) Destination(
	ctx context.Context,
	from nvector.Position,
	delta nvector.Vector,
	e nvector.Ellipsoid,
	f nvector.Matrix,
) (nvector.Position, error) {
	type res struct {
		V [][]float64 `json:"n_EB_E"`
		D float64     `json:"z_EB"`
	}

	data, err := call(
		ctx,
		c,
		unmarshalAs[res],
		"n_EA_E_and_p_AB_E2n_EB_E",
		map[string]any{
			"n_EA_E": marshalVector(from.Vector),
			"z_EA":   from.Depth,
			"p_AB_E": marshalVector(delta),
			"a":      e.SemiMajorAxis,
			"f":      e.Flattening,
			"R_Ee":   marshalMatrix(f),
		},
	)
	if err != nil {
		return nvector.Position{}, err
	}

	return nvector.Position{Vector: unmarshalVector(data.V), Depth: data.D}, nil
}
