package view

// HitRectangle is the wrapper of image.Rectangle to check hits.
type HitRectangle struct {
	min, max *Vector
}

// NewHitRectangle returns a new HitRectangle.
func NewHitRectangle(min, max Vector) *HitRectangle {
	return &HitRectangle{
		min: &Vector{min.X, min.Y},
		max: &Vector{max.X, max.Y},
	}
}

// Add adds the v to the rectangle to check hits.
func (hr HitRectangle) Add(v *Vector) {
	hr.min.X += v.X
	hr.min.Y += v.Y
	hr.max.X += v.X
	hr.max.Y += v.Y
}

// Overlaps returns whether this HitRectangle hits the arg.
func (hr HitRectangle) Overlaps(hrP *HitRectangle) bool {
	if hrP.max.X < hr.min.X || hr.max.X < hrP.min.X {
		return false
	}
	if hrP.max.Y < hr.min.Y || hr.max.Y < hrP.min.Y {
		return false
	}
	return true
}
