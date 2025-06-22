package geometry

import "math"

func NewVector(x, y float32) Vector {
	return Vector{x, y}
}

type Vector struct {
	X, Y float32
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

func (v Vector) Mult(scale float32) Vector {
	return Vector{v.X * scale, v.Y * scale}
}

func (v Vector) Dot(other Vector) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y}
}

func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector) Angle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v Vector) Rotate(r float32, center Vector) Vector {
	newX := center.X + (v.X-center.X)*float32(math.Cos(float64(r))) - (v.Y-center.Y)*float32(math.Sin(float64(r)))
	newY := center.Y + (v.X-center.X)*float32(math.Sin(float64(r))) + (v.Y-center.Y)*float32(math.Cos(float64(r)))

	return Vector{newX, newY}
}

func Direction(r float32) Vector {
	r64 := float64(r)
	return Vector{
		X: float32(math.Cos(r64)),
		Y: float32(math.Sin(r64)),
	}
}
