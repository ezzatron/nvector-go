package nvector

import (
	"gonum.org/v1/gonum/spatial/r3"
)

// Delta finds a delta ECEF position vector from a reference n-vector, a depth,
// and a target n-vector.
//
// fromDepth and toDepth are in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EA_E_and_n_EB_E2p_AB_E.m
func Delta(
	from r3.Vec,
	fromDepth float64,
	to r3.Vec,
	toDepth float64,
	opts ...Option,
) r3.Vec {
	// Function 1. in Section 5.4 in Gade (2010):
	return r3.Sub(ToECEF(to, toDepth, opts...), ToECEF(from, fromDepth, opts...))
}

// Destination finds a n-vector from a reference n-vector, a depth, and a delta
// ECEF position vector.
//
// fromDepth and toDepth are in meters below the ellipsoid surface.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EA_E_and_p_AB_E2n_EB_E.m
func Destination(
	from r3.Vec,
	fromDepth float64,
	delta r3.Vec,
	opts ...Option,
) (to r3.Vec, toDepth float64) {
	// Function 2. in Section 5.4 in Gade (2010):
	return FromECEF(r3.Add(ToECEF(from, fromDepth, opts...), delta), opts...)
}
