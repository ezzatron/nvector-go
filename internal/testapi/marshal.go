package testapi

import (
	"gonum.org/v1/gonum/spatial/r3"
)

func marshalMatrix(m *r3.Mat) [][]float64 {
	return [][]float64{
		{m.At(0, 0), m.At(0, 1), m.At(0, 2)},
		{m.At(1, 0), m.At(1, 1), m.At(1, 2)},
		{m.At(2, 0), m.At(2, 1), m.At(2, 2)},
	}
}

func marshalVector(v r3.Vec) [][]float64 {
	return [][]float64{{v.X}, {v.Y}, {v.Z}}
}

func unmarshalAs[J any](data J) J {
	return data
}

func unmarshalVector(data [][]float64) r3.Vec {
	return r3.Vec{X: data[0][0], Y: data[1][0], Z: data[2][0]}
}
