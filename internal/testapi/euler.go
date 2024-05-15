package testapi

import (
	"context"

	"gonum.org/v1/gonum/spatial/r3"
)

// XYZToRotationMat converts Euler angles in XYZ order to a rotation matrix.
func (c *Client) XYZToRotationMat(
	ctx context.Context,
	x, y, z float64,
) (*r3.Mat, error) {
	return call(ctx, c, unmarshalMatrix, "xyz2R", map[string]any{
		"x": x,
		"y": y,
		"z": z,
	})
}

// ZYXToRotationMat converts Euler angles in ZYX order to a rotation matrix.
func (c *Client) ZYXToRotationMat(
	ctx context.Context,
	z, y, x float64,
) (*r3.Mat, error) {
	return call(ctx, c, unmarshalMatrix, "zyx2R", map[string]any{
		"z": z,
		"y": y,
		"x": x,
	})
}

// RotationMatToXYZ converts a rotation matrix to Euler angles in XYZ order.
func (c *Client) RotationMatToXYZ(
	ctx context.Context,
	r *r3.Mat,
) (x, y, z float64, err error) {
	type res struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "R2xyz", map[string]any{
		"R_AB": marshalMatrix(r),
	})
	if err != nil {
		return 0, 0, 0, err
	}

	return data.X, data.Y, data.Z, nil
}

// RotationMatToZYX converts a rotation matrix to Euler angles in ZYX order.
func (c *Client) RotationMatToZYX(
	ctx context.Context,
	r *r3.Mat,
) (z, y, x float64, err error) {
	type res struct {
		Z float64 `json:"z"`
		Y float64 `json:"y"`
		X float64 `json:"x"`
	}

	data, err := call(ctx, c, unmarshalAs[res], "R2zyx", map[string]any{
		"R_AB": marshalMatrix(r),
	})
	if err != nil {
		return 0, 0, 0, err
	}

	return data.Z, data.Y, data.X, nil
}
