package view

// Vector is the vector to define the position or velocity of the objects.
type Vector struct {
	X float64
	Y float64
}

// Add returns the added vector.
func (v *Vector) Add(vec Vector) {
	v.X += vec.X
	v.Y += vec.Y
}
