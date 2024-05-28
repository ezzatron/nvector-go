package nvector

// Ellipsoid is a reference ellipsoid.
type Ellipsoid struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	Flattening    float64
}

var (
	// GRS80 is the Geodetic Reference System 1980 ellipsoid.
	//
	// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L45
	GRS80 = Ellipsoid{6378137, 6356752.314140356, 1 / 298.257222101}

	// WGS72 is the World Geodetic System 1972 ellipsoid.
	//
	// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L47
	WGS72 = Ellipsoid{6378135, 6356750.520016094, 1 / 298.26}

	// WGS84 is the World Geodetic System 1984 ellipsoid.
	//
	// See: https://github.com/chrisveness/geodesy/blob/761587cd748bd9f7c9825195eba4a9fc5891b859/latlon-ellipsoidal-datum.js#L39
	WGS84 = Ellipsoid{6378137, 6356752.314245179, 1 / 298.257223563}
)

// Sphere returns a spherical ellipsoid with the given radius.
func Sphere(radius float64) Ellipsoid {
	return Ellipsoid{radius, radius, 0}
}
