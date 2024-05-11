package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/mat"
)

// FromECEF converts an ECEF vector to an n-vector and depth.
//
// The depth is returned in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/p_EB_E2n_EB_E.m
func FromECEF(ecef mat.Vector, opts ...Option) (mat.Vector, float64) {
	o := options.New(opts)
	a := o.Ellipsoid.SemiMajorAxis
	f := o.Ellipsoid.Flattening

	// CoordFrame selects correct E-axes
	ecefr := &mat.VecDense{}
	ecefr.MulVec(o.CoordFrame, ecef)

	// e2 = eccentricity^2
	e2 := 2*f - math.Pow(f, 2)

	// The following code implements equation (23) from Gade (2010):
	R2 := math.Pow(ecefr.AtVec(1), 2) + math.Pow(ecefr.AtVec(2), 2)
	// R = component of ecef in the equatorial plane
	R := math.Sqrt(R2)

	p := R2 / math.Pow(a, 2)
	q := (1 - e2) / math.Pow(a, 2) * math.Pow(ecefr.AtVec(0), 2)
	r := (p + q - math.Pow(e2, 2)) / 6

	s := math.Pow(e2, 2) * p * q / (4 * math.Pow(r, 3))
	t := math.Cbrt(1 + s + math.Sqrt(s*(2+s)))
	u := r * (1 + t + 1/t)
	v := math.Sqrt(math.Pow(u, 2) + math.Pow(e2, 2)*q)

	w := e2 * (u + v - q) / (2 * v)
	k := math.Sqrt(u+v+math.Pow(w, 2)) - w
	d := k * R / (k + e2)

	// Calculate height:
	height := (k + e2 - 1) /
		k * math.Sqrt(math.Pow(d, 2)+math.Pow(ecefr.AtVec(0), 2))

	temp := 1 / math.Sqrt(math.Pow(d, 2)+math.Pow(ecefr.AtVec(0), 2))

	nv := mat.NewVecDense(3, []float64{
		temp * ecefr.AtVec(0),
		temp * k / (k + e2) * ecefr.AtVec(1),
		temp * k / (k + e2) * ecefr.AtVec(2),
	})

	// Ensure unit length:
	nv.MulVec(o.CoordFrame.T(), nv)
	nv.ScaleVec(1/nv.Norm(2), nv)

	return nv, -height
}

// ToECEF converts an n-vector and depth to an ECEF vector.
//
// d is given in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EB_E2p_EB_E.m
func ToECEF(nv mat.Vector, d float64, opts ...Option) mat.Vector {
	o := options.New(opts)
	a := o.Ellipsoid.SemiMajorAxis
	f := o.Ellipsoid.Flattening

	// CoordFrame selects correct E-axes
	nvr := &mat.VecDense{}
	nvr.MulVec(o.CoordFrame, nv)

	// semi-minor axis:
	b := a * (1 - f)

	// The following code implements equation (22) in Gade (2010):

	denom := math.Sqrt(math.Pow(nvr.AtVec(0), 2) +
		math.Pow(nvr.AtVec(1), 2)/math.Pow(1-f, 2) +
		math.Pow(nvr.AtVec(2), 2)/math.Pow(1-f, 2))

	// We first calculate the position at the origin of coordinate system L, which
	// has the same n-vector as B (n_EL_E = n_EB_E), but lies at the surface of
	// the Earth (z_EL = 0).

	pELE := mat.NewVecDense(3, []float64{
		b / denom * nvr.AtVec(0),
		b / denom * nvr.AtVec(1) / math.Pow(1-f, 2),
		b / denom * nvr.AtVec(2) / math.Pow(1-f, 2),
	})

	nvrs := &mat.VecDense{}
	nvrs.ScaleVec(d, nvr)

	pEBE := &mat.VecDense{}
	pEBE.SubVec(pELE, nvrs)
	pEBE.MulVec(o.CoordFrame.T(), pEBE)

	return pEBE
}
