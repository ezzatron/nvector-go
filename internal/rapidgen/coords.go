package rapidgen

import (
	"math"

	"github.com/ezzatron/nvector-go"
	"pgregory.net/rapid"
)

// GeodeticCoordinates creates a rapid generator for geodetic coordinates.
func GeodeticCoordinates() *rapid.Generator[nvector.GeodeticCoordinates] {
	return rapid.Custom(func(t *rapid.T) nvector.GeodeticCoordinates {
		return nvector.GeodeticCoordinates{
			Latitude:  rapid.Float64Range(-math.Pi/2, math.Pi/2).Draw(t, "latitude"),
			Longitude: rapid.Float64Range(-math.Pi, math.Pi).Draw(t, "longitude"),
		}
	})
}
