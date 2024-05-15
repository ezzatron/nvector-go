package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// FromECEF converts an ECEF position vector to an n-vector and depth.
func (c *Client) FromECEF(
	ctx context.Context,
	ecef r3.Vec,
	opts ...options.Option,
) (r3.Vec, float64, error) {
	o := options.New(opts)

	type res struct {
		Nv [][]float64 `json:"n_EB_E"`
		D  float64     `json:"depth"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "p_EB_E2n_EB_E", map[string]any{
		"p_EB_E": marshalVector(ecef),
		"a":      o.Ellipsoid.SemiMajorAxis,
		"f":      o.Ellipsoid.Flattening,
		"R_Ee":   marshalMatrix(o.CoordFrame),
	})
	if err != nil {
		return r3.Vec{}, 0, err
	}

	return unmarshalVector(data.Nv), data.D, nil
}

// ToECEF converts an n-vector and depth to an ECEF position vector.
func (c *Client) ToECEF(
	ctx context.Context,
	nv r3.Vec,
	d float64,
	opts ...options.Option,
) (r3.Vec, error) {
	o := options.New(opts)

	return call(ctx, c, unmarshalVector, "n_EB_E2p_EB_E", map[string]any{
		"n_EB_E": marshalVector(nv),
		"depth":  d,
		"a":      o.Ellipsoid.SemiMajorAxis,
		"f":      o.Ellipsoid.Flattening,
		"R_Ee":   marshalMatrix(o.CoordFrame),
	})
}
