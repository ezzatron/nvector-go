package nvector

import (
	"math"
)

// FromECEF converts an ECEF position vector to an n-vector and depth.
//
// f is the coordinate frame in which the vectors are decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/p_EB_E2n_EB_E.m
func FromECEF(v Vector, e Ellipsoid, f Matrix) Position {
	// f selects correct E-axes
	v = v.Transform(f)

	// e2 = eccentricity^2
	e2 := 2*e.Flattening - math.Pow(e.Flattening, 2)

	// The following code implements equation (23) from Gade (2010):
	R2 := math.Pow(v.Y, 2) + math.Pow(v.Z, 2)
	// R = component of v in the equatorial plane
	R := math.Sqrt(R2)

	x2 := math.Pow(v.X, 2)

	p := R2 / math.Pow(e.SemiMajorAxis, 2)
	q := (1 - e2) / math.Pow(e.SemiMajorAxis, 2) * x2
	r := (p + q - math.Pow(e2, 2)) / 6

	s := math.Pow(e2, 2) * p * q / (4 * math.Pow(r, 3))
	t := math.Cbrt(1 + s + math.Sqrt(s*(2+s)))
	u := r * (1 + t + 1/t)
	// This variable was named v in the original code, but that clashes with the
	// input variable name.
	V := math.Sqrt(math.Pow(u, 2) + math.Pow(e2, 2)*q)

	w := e2 * (u + V - q) / (2 * V)
	k := math.Sqrt(u+V+math.Pow(w, 2)) - w
	d := k * R / (k + e2)

	// Calculate height:
	hf := math.Sqrt(math.Pow(d, 2) + x2)
	h := (k + e2 - 1) / k * hf

	temp := 1 / hf

	return Position{
		// Ensure unit length:
		Vector{
			temp * v.X,
			temp * k / (k + e2) * v.Y,
			temp * k / (k + e2) * v.Z,
		}.Transform(f.Transpose()).Normalize(),
		-h,
	}
}

// ToECEF converts an n-vector and depth to an ECEF position vector.
//
// f is the coordinate frame in which the vectors are decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EB_E2p_EB_E.m
func ToECEF(v Position, e Ellipsoid, f Matrix) Vector {
	ef := e.Flattening

	// f selects correct E-axes
	nv := v.Vector.Transform(f)

	// semi-minor axis:
	eb := e.SemiMinorAxis

	// The following code implements equation (22) in Gade (2010):

	denom := math.Sqrt(math.Pow(nv.X, 2) +
		math.Pow(nv.Y, 2)/math.Pow(1-ef, 2) +
		math.Pow(nv.Z, 2)/math.Pow(1-ef, 2))

	// We first calculate the position at the origin of coordinate system L, which
	// has the same n-vector as B (n_EL_E = n_EB_E), but lies at the surface of
	// the Earth (depth = 0).

	x := eb / denom * nv.X
	y := eb / denom * nv.Y / math.Pow(1-ef, 2)
	z := eb / denom * nv.Z / math.Pow(1-ef, 2)

	return Vector{
		x - nv.X*v.Depth,
		y - nv.Y*v.Depth,
		z - nv.Z*v.Depth,
	}.Transform(f.Transpose())
}

// Position is an n-vector with an associated depth. Together, they represent a
// concrete position relative to the ellipsoid, similar to an ECEF position
// vector.
type Position struct {
	// Vector is the n-vector.
	Vector Vector
	// Depth is the depth in meters below the ellipsoid.
	Depth float64
}
