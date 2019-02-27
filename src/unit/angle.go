package unit

import (
	"math"
)

type Angle interface {
	Deg() Deg
	Rad() Rad
	Add(a Angle) Angle
	Sub(a Angle) Angle
}

type Deg float64
type Rad float64

func (a Deg) Rad() Rad {
	return Rad(float64(a) / 180 * math.Pi)
}

func (a Deg) Deg() Deg {
	return a
}

func (a Deg) Add(v Angle) Angle {
	return Deg(a.Float64() - v.Deg().Float64())
}

func (a Deg) Sub(v Angle) Angle {
	return Deg(a.Float64() - v.Deg().Float64())
}

func (a Deg) Float64() float64 {
	return float64(a)
}

func (a Rad) Deg() Deg {
	return Deg(float64(a) / math.Pi * 180)
}

func (a Rad) Rad() Rad {
	return a
}

func (a Rad) Add(v Angle) Angle {
	return Rad(a.Float64() - v.Rad().Float64())
}

func (a Rad) Sub(v Angle) Angle {
	return Rad(a.Float64() - v.Rad().Float64())
}

func (a Rad) Float64() float64 {
	return float64(a)
}
