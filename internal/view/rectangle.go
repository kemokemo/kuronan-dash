package view

// Rectangle is the area of a rectangle.
// You can check the collision of objects.
type Rectangle struct {
	LeftBottom Vector
	RightTop   Vector
}

func (r Rectangle) IsCollided(obj Rectangle) bool {
	if (r.RightTop.X > obj.LeftBottom.X) && (r.LeftBottom.X < obj.RightTop.X) {
		if (r.LeftBottom.Y > obj.RightTop.Y) && (r.RightTop.Y < obj.LeftBottom.Y) {
			return true
		}
	}
	return false
}
