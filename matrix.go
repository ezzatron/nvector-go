package nvector

// Matrix is a 3x3 matrix.
type Matrix struct {
	XX, XY, XZ float64
	YX, YY, YZ float64
	ZX, ZY, ZZ float64
}

// Multiply multiplies two matrices.
func (m Matrix) Multiply(n Matrix) Matrix {
	return Matrix{
		m.XX*n.XX + m.XY*n.YX + m.XZ*n.ZX,
		m.XX*n.XY + m.XY*n.YY + m.XZ*n.ZY,
		m.XX*n.XZ + m.XY*n.YZ + m.XZ*n.ZZ,

		m.YX*n.XX + m.YY*n.YX + m.YZ*n.ZX,
		m.YX*n.XY + m.YY*n.YY + m.YZ*n.ZY,
		m.YX*n.XZ + m.YY*n.YZ + m.YZ*n.ZZ,

		m.ZX*n.XX + m.ZY*n.YX + m.ZZ*n.ZX,
		m.ZX*n.XY + m.ZY*n.YY + m.ZZ*n.ZY,
		m.ZX*n.XZ + m.ZY*n.YZ + m.ZZ*n.ZZ,
	}
}

// Transpose returns the transpose of m.
func (m Matrix) Transpose() Matrix {
	return Matrix{
		m.XX, m.YX, m.ZX,
		m.XY, m.YY, m.ZY,
		m.XZ, m.YZ, m.ZZ,
	}
}
