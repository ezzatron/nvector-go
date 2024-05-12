package testapi

import (
	"context"

	"gonum.org/v1/gonum/spatial/r3"
)

// EulerXYZToRotMat converts Euler angles in XYZ order to a rotation matrix.
func (c *Client) EulerXYZToRotMat(
	ctx context.Context,
	x, y, z float64,
) (*r3.Mat, error) {
	return call(ctx, c, unmarshalMatrix, "xyz2R", map[string]any{
		"x": x,
		"y": y,
		"z": z,
	})
}

// EulerZYXToRotMat converts Euler angles in ZYX order to a rotation matrix.
func (c *Client) EulerZYXToRotMat(
	ctx context.Context,
	z, y, x float64,
) (*r3.Mat, error) {
	return call(ctx, c, unmarshalMatrix, "zyx2R", map[string]any{
		"z": z,
		"y": y,
		"x": x,
	})
}

// RotMatToEulerXYZ converts a rotation matrix to Euler angles in XYZ order.
func (c *Client) RotMatToEulerXYZ(
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

// RotMatToEulerZYX converts a rotation matrix to Euler angles in ZYX order.
func (c *Client) RotMatToEulerZYX(
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
