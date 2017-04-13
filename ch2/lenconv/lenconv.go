// Package lenconv performs various conversions for length.
package lenconv

import (
	"fmt"
)

// Feet describes a length measurement in feet
type Feet float64

// Meters describes a length measurement in meters
type Meters float64

// FToM converts a measurement in Feet to one in Meters
func FToM(f Feet) Meters {
	return Meters(f / 3.28)
}

// MToF converts a measurement in Meters to one in Feet
func MToF(m Meters) Feet {
	return Feet(m * 3.28)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gf", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%gm", m)
}
