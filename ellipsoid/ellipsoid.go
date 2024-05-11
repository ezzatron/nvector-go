package ellipsoid

import "github.com/ezzatron/nvector-go/internal/options"

// Ellipsoid is a reference ellipsoid.
type Ellipsoid = options.Ellipsoid

// GRS80 is the Geodetic Reference System 1980 ellipsoid.
//
// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L45
func GRS80() Ellipsoid {
	return Ellipsoid{
		SemiMajorAxis: 6378137,
		Flattening:    1 / 298.257222101,
	}
}

// WGS72 is the World Geodetic System 1972 ellipsoid.
//
// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L47
func WGS72() Ellipsoid {
	return Ellipsoid{
		SemiMajorAxis: 6378135,
		Flattening:    1 / 298.26,
	}
}

// WGS84 is the World Geodetic System 1984 ellipsoid.
//
// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L39
func WGS84() Ellipsoid {
	return Ellipsoid{
		SemiMajorAxis: 6378137,
		Flattening:    1 / 298.257223563,
	}
}

// WGS84Sphere is a sphere with the same semi-major axis as the WGS-84 ellipsoid.
//
// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L39
func WGS84Sphere() Ellipsoid {
	return Ellipsoid{
		SemiMajorAxis: 6378137,
		Flattening:    0,
	}
}
