package testapi

import "gonum.org/v1/gonum/mat"

func marshalMatrix(m mat.Matrix) [][]float64 {
	r, c := m.Dims()
	data := make([][]float64, r)

	for i := 0; i < r; i++ {
		data[i] = make([]float64, c)

		for j := 0; j < c; j++ {
			data[i][j] = m.At(i, j)
		}
	}

	return data
}

func unmarshalVector(data [][]float64) *mat.VecDense {
	return mat.NewVecDense(3, []float64{data[0][0], data[1][0], data[2][0]})
}
