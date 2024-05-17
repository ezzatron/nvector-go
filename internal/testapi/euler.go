package testapi

import (
	"context"

	"github.com/ezzatron/nvector-go"
)

// XYZToRotationMatrix converts Euler angles in XYZ order to a rotation matrix.
func (c *Client) XYZToRotationMatrix(
	ctx context.Context,
	a nvector.EulerXYZ,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "xyz2R", map[string]any{
		"x": a.X,
		"y": a.Y,
		"z": a.Z,
	})
}

// ZYXToRotationMatrix converts Euler angles in ZYX order to a rotation matrix.
func (c *Client) ZYXToRotationMatrix(
	ctx context.Context,
	a nvector.EulerZYX,
) (nvector.Matrix, error) {
	return call(ctx, c, unmarshalMatrix, "zyx2R", map[string]any{
		"z": a.Z,
		"y": a.Y,
		"x": a.X,
	})
}

// RotationMatrixToXYZ converts a rotation matrix to Euler angles in XYZ order.
func (c *Client) RotationMatrixToXYZ(
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

// RotationMatrixToZYX converts a rotation matrix to Euler angles in ZYX order.
func (c *Client) RotationMatrixToZYX(
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
