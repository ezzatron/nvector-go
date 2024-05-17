package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// FromRotationMatrix converts a rotation matrix to an n-vector.
func (c *Client) FromRotationMatrix(
	ctx context.Context,
	r nvector.Matrix,
) (nvector.Vector, error) {
	return call(ctx, c, unmarshalVector, "R_EN2n_E", map[string]any{
		"R_EN": marshalMatrix(r),
	})
}

// ToRotationMatrix converts an n-vector to a rotation matrix.
func (c *Client) ToRotationMatrix(
	ctx context.Context,
	v nvector.Vector,
	f nvector.Matrix,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "n_E2R_EN", map[string]any{
		"n_E":  marshalVector(v),
		"R_Ee": marshalMatrix(f),
	})
}

// ToRotationMatrixUsingWanderAzimuth converts an n-vector and a wander azimuth
// angle to a rotation matrix.
func (c *Client) ToRotationMatrixUsingWanderAzimuth(
	ctx context.Context,
	v nvector.Vector,
	w float64,
	f nvector.Matrix,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "n_E_and_wa2R_EL", map[string]any{
		"n_E":            marshalVector(v),
		"wander_azimuth": w,
		"R_Ee":           marshalMatrix(f),
	})
}
