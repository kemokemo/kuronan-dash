package view

// Vector is the vector to define the position or velocity of the objects.
type Vector struct {
	X float64
	Y float64
}

// Add returns the added vector.
func (v *Vector) Add(vec Vector) Vector {
	return Vector{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}
