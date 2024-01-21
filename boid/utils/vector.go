package utils

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vector) Mul(v2 Vector) Vector {
	return Vector{v1.X * v2.X, v1.Y * v2.Y}
}

func (v1 Vector) AddNum(val float64) Vector {
	return Vector{v1.X + val, v1.Y + val}
}

func (v1 Vector) MulNum(val float64) Vector {
	return Vector{v1.X * val, v1.Y * val}
}

func (v1 Vector) DivNum(val float64) Vector {
	return Vector{v1.X / val, v1.Y / val}
}

func (v1 Vector) Limit(lower, upper float64) Vector {
	return Vector{math.Min(math.Max(v1.X, lower), upper), math.Min(math.Max(v1.Y, lower), upper)}
}

func (v1 Vector) Dist(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
