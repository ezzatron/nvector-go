package nvector

// Delta finds a delta ECEF position vector from a reference n-vector, a depth,
// and a target n-vector.
//
// f is the coordinate frame in which the n-vectors are decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EA_E_and_n_EB_E2p_AB_E.m
func Delta(from Position, to Position, e Ellipsoid, f Matrix) Vector {
	// Function 1. in Section 5.4 in Gade (2010):
	return ToECEF(to, e, f).Sub(ToECEF(from, e, f))
}

// Destination finds a n-vector from a reference n-vector, a depth, and a delta
// ECEF position vector.
//
// f is the coordinate frame in which the n-vectors are decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_EA_E_and_p_AB_E2n_EB_E.m
func Destination(from Position, delta Vector, e Ellipsoid, f Matrix) Position {
	// Function 2. in Section 5.4 in Gade (2010):
	return FromECEF(ToECEF(from, e, f).Add(delta), e, f)
}
