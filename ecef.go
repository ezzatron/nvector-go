package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// FromECEF converts an ECEF position vector to an n-vector and depth.
//
// The depth is returned in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/p_EB_E2n_EB_E.m
func FromECEF(ecef r3.Vec, opts ...Option) (r3.Vec, float64) {
	o := options.New(opts)
	a := o.Ellipsoid.SemiMajorAxis
	f := o.Ellipsoid.Flattening

	// CoordFrame selects correct E-axes
	ecefr := o.CoordFrame.MulVec(ecef)

	// e2 = eccentricity^2
	e2 := 2*f - math.Pow(f, 2)

	// The following code implements equation (23) from Gade (2010):
	R2 := math.Pow(ecefr.Y, 2) + math.Pow(ecefr.Z, 2)
	// R = component of ecef in the equatorial plane
	R := math.Sqrt(R2)

	p := R2 / math.Pow(a, 2)
	q := (1 - e2) / math.Pow(a, 2) * math.Pow(ecefr.X, 2)
	r := (p + q - math.Pow(e2, 2)) / 6

	s := math.Pow(e2, 2) * p * q / (4 * math.Pow(r, 3))
	t := math.Cbrt(1 + s + math.Sqrt(s*(2+s)))
	u := r * (1 + t + 1/t)
	v := math.Sqrt(math.Pow(u, 2) + math.Pow(e2, 2)*q)

	w := e2 * (u + v - q) / (2 * v)
	k := math.Sqrt(u+v+math.Pow(w, 2)) - w
	d := k * R / (k + e2)

	// Calculate height:
	height := (k + e2 - 1) / k * math.Sqrt(math.Pow(d, 2)+math.Pow(ecefr.X, 2))

	temp := 1 / math.Sqrt(math.Pow(d, 2)+math.Pow(ecefr.X, 2))

	nv := r3.Vec{
		X: temp * ecefr.X,
		Y: temp * k / (k + e2) * ecefr.Y,
		Z: temp * k / (k + e2) * ecefr.Z,
	}

	// Ensure unit length:
	return r3.Unit(o.CoordFrame.MulVecTrans(nv)), -height
}

// ToECEF converts an n-vector and depth to an ECEF position vector.
//
// d is given in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EB_E2p_EB_E.m
func ToECEF(nv r3.Vec, d float64, opts ...Option) r3.Vec {
	o := options.New(opts)
	a := o.Ellipsoid.SemiMajorAxis
	f := o.Ellipsoid.Flattening

	// CoordFrame selects correct E-axes
	nvr := o.CoordFrame.MulVec(nv)

	// semi-minor axis:
	b := a * (1 - f)

	// The following code implements equation (22) in Gade (2010):

	denom := math.Sqrt(math.Pow(nvr.X, 2) +
		math.Pow(nvr.Y, 2)/math.Pow(1-f, 2) +
		math.Pow(nvr.Z, 2)/math.Pow(1-f, 2))

	// We first calculate the position at the origin of coordinate system L, which
	// has the same n-vector as B (n_EL_E = n_EB_E), but lies at the surface of
	// the Earth (z_EL = 0).

	pELE := r3.Vec{
		X: b / denom * nvr.X,
		Y: b / denom * nvr.Y / math.Pow(1-f, 2),
		Z: b / denom * nvr.Z / math.Pow(1-f, 2),
	}

	return o.CoordFrame.MulVecTrans(r3.Sub(pELE, r3.Scale(d, nvr)))
}
