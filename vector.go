package nvector

import (
	"math"
)

// Vector is a 3D vector.
type Vector struct {
	X, Y, Z float64
}

// Add returns v+w.
func (v Vector) Add(w Vector) Vector {
	return Vector{
		v.X + w.X,
		v.Y + w.Y,
		v.Z + w.Z,
	}
}

// Cross returns the cross product of v and w.
func (v Vector) Cross(w Vector) Vector {
	return Vector{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// Dot returns the dot product of v and w.
func (v Vector) Dot(w Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Norm returns the Euclidean norm of v.
func (v Vector) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a vector in the same direction as v but with norm 1.
func (v Vector) Normalize() Vector {
	return v.Scale(1 / v.Norm())
}

// Scale returns v scaled by s.
func (v Vector) Scale(s float64) Vector {
	return Vector{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}

// Sub returns v-w.
func (v Vector) Sub(w Vector) Vector {
	return Vector{
		v.X - w.X,
		v.Y - w.Y,
		v.Z - w.Z,
	}
}

// Transform returns v transformed by m.
func (v Vector) Transform(m Matrix) Vector {
	return Vector{
		m.XX*v.X + m.XY*v.Y + m.XZ*v.Z,
		m.YX*v.X + m.YY*v.Y + m.YZ*v.Z,
		m.ZX*v.X + m.ZY*v.Y + m.ZZ*v.Z,
	}
}
