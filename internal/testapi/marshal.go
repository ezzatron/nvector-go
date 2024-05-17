package testapi

import (
	"github.com/ezzatron/nvector-go"
)

func marshalMatrix(m nvector.Matrix) [][]float64 {
	return [][]float64{
		{m.XX, m.XY, m.XZ},
		{m.YX, m.YY, m.YZ},
		{m.ZX, m.ZY, m.ZZ},
	}
}

func marshalVector(v nvector.Vector) [][]float64 {
	return [][]float64{{v.X}, {v.Y}, {v.Z}}
}

func unmarshalAs[J any](data J) J {
	return data
}

func unmarshalMatrix(data [][]float64) nvector.Matrix {
	return nvector.Matrix{
		XX: data[0][0], XY: data[0][1], XZ: data[0][2],
		YX: data[1][0], YY: data[1][1], YZ: data[1][2],
		ZX: data[2][0], ZY: data[2][1], ZZ: data[2][2],
	}
}

func unmarshalVector(data [][]float64) nvector.Vector {
	return nvector.Vector{X: data[0][0], Y: data[1][0], Z: data[2][0]}
}
