package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// FromECEF converts an ECEF position vector to an n-vector and depth.
func (c *Client) FromECEF(
	ctx context.Context,
	v nvector.Vector,
	e nvector.Ellipsoid,
	f nvector.Matrix,
) (nvector.Position, error) {
	type res struct {
		V [][]float64 `json:"n_EB_E"`
		D float64     `json:"depth"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "p_EB_E2n_EB_E", map[string]any{
		"p_EB_E": marshalVector(v),
		"a":      e.SemiMajorAxis,
		"f":      e.Flattening,
		"R_Ee":   marshalMatrix(f),
	})
	if err != nil {
		return nvector.Position{}, err
	}

	return nvector.Position{Vector: unmarshalVector(data.V), Depth: data.D}, nil
}

// ToECEF converts an n-vector and depth to an ECEF position vector.
func (c *Client) ToECEF(
	ctx context.Context,
	v nvector.Position,
	e nvector.Ellipsoid,
	f nvector.Matrix,
) (nvector.Vector, error) {
	return call(ctx, c, unmarshalVector, "n_EB_E2p_EB_E", map[string]any{
		"n_EB_E": marshalVector(v.Vector),
		"depth":  v.Depth,
		"a":      e.SemiMajorAxis,
		"f":      e.Flattening,
		"R_Ee":   marshalMatrix(f),
	})
}
