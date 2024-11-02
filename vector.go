package geometry

import "math"

func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

type Vector struct {
	X, Y float64
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

func (v Vector) Mult(scale float64) Vector {
	return Vector{v.X * scale, v.Y * scale}
}

func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}
