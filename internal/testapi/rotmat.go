package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// FromRotMat converts a rotation matrix to an n-vector.
func (c *Client) FromRotMat(ctx context.Context, r *r3.Mat) (r3.Vec, error) {
	return call(ctx, c, unmarshalVector, "R_EN2n_E", map[string]any{
		"R_EN": marshalMatrix(r),
	})
}

// ToRotMat converts n-vector to a rotation matrix.
func (c *Client) ToRotMat(
	ctx context.Context,
	nv r3.Vec,
	opts ...options.Option,
) (*r3.Mat, error) {
	o := options.New(opts)

	return call(ctx, c, unmarshalMatrix, "n_E2R_EN", map[string]any{
		"n_E":  marshalVector(nv),
		"R_Ee": marshalMatrix(o.CoordFrame),
	})
}
