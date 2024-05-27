package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// EulerXYZToRotationMatrix converts Euler angles in XYZ order to a rotation matrix.
func (c *Client) EulerXYZToRotationMatrix(
	ctx context.Context,
	a nvector.EulerXYZ,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "xyz2R", map[string]any{
		"x": a.X,
		"y": a.Y,
		"z": a.Z,
	})
}

// EulerZYXToRotationMatrix converts Euler angles in ZYX order to a rotation matrix.
func (c *Client) EulerZYXToRotationMatrix(
	ctx context.Context,
	a nvector.EulerZYX,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "zyx2R", map[string]any{
		"z": a.Z,
		"y": a.Y,
		"x": a.X,
	})
}

// RotationMatrixToEulerXYZ converts a rotation matrix to Euler angles in XYZ order.
func (c *Client) RotationMatrixToEulerXYZ(
	ctx context.Context,
	r nvector.Matrix,
) (nvector.EulerXYZ, error) {
	type res struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "R2xyz", map[string]any{
		"R_AB": marshalMatrix(r),
	})
	if err != nil {
		return nvector.EulerXYZ{}, err
	}

	return nvector.EulerXYZ(data), nil
}

// RotationMatrixToEulerZYX converts a rotation matrix to Euler angles in ZYX order.
func (c *Client) RotationMatrixToEulerZYX(
	ctx context.Context,
	r nvector.Matrix,
) (nvector.EulerZYX, error) {
	type res struct {
		Z float64 `json:"z"`
		Y float64 `json:"y"`
		X float64 `json:"x"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "R2zyx", map[string]any{
		"R_AB": marshalMatrix(r),
	})
	if err != nil {
		return nvector.EulerZYX{}, err
	}

	return nvector.EulerZYX(data), nil
}
