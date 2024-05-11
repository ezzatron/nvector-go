package rapidgen

import (
	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
	"pgregory.net/rapid"
)

// Depth creates a rapid generator for depths relative to an ellipsoid.
func Depth(e options.Ellipsoid) *rapid.Generator[float64] {
	// Semi-minor axis
	b := e.SemiMajorAxis * (1 - e.Flattening)

	return rapid.Float64Range(-b, b)
}

// EcefVector creates a rapid generator for ECEF vectors relative to an
// ellipsoid.
func EcefVector(e options.Ellipsoid) *rapid.Generator[r3.Vec] {
	// Semi-minor axis
	b := e.SemiMajorAxis * (1 - e.Flattening)

	return VectorRange(e.SemiMajorAxis-b, e.SemiMajorAxis+b)
}

// Ellipsoid creates a rapid generator for ellipsoids.
func Ellipsoid() *rapid.Generator[options.Ellipsoid] {
	return rapid.SampledFrom([]options.Ellipsoid{
		{
			SemiMajorAxis: 6378137,
			Flattening:    1 / 298.257222101,
		},
		{
			SemiMajorAxis: 6378135,
			Flattening:    1 / 298.26,
		},
		{
			SemiMajorAxis: 6378137,
			Flattening:    1 / 298.257223563,
		},
		{
			SemiMajorAxis: 6378137,
			Flattening:    0,
		},
	})
}
