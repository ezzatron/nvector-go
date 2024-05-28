package rapidgen

import (
	"github.com/ezzatron/nvector-go"
	"pgregory.net/rapid"
)

// Depth creates a rapid generator for depths relative to an ellipsoid.
func Depth(e nvector.Ellipsoid) *rapid.Generator[float64] {
	return rapid.Float64Range(-e.SemiMinorAxis, e.SemiMinorAxis)
}

// EcefVector creates a rapid generator for ECEF position vectors relative to an
// ellipsoid.
func EcefVector(e nvector.Ellipsoid) *rapid.Generator[nvector.Vector] {
	return VectorRange(
		e.SemiMajorAxis-e.SemiMinorAxis,
		e.SemiMajorAxis+e.SemiMinorAxis,
	)
}

// Ellipsoid creates a rapid generator for ellipsoids.
func Ellipsoid() *rapid.Generator[nvector.Ellipsoid] {
	return rapid.SampledFrom([]nvector.Ellipsoid{
		nvector.WGS84,
		nvector.GRS80,
		nvector.WGS72,
	})
}
