package view

import (
	"testing"
)

func TestViewport_Move(t *testing.T) {
	v := Viewport{}
	v.SetSize(256, 128)
	x, y := v.Position()
	if x != 0 || y != 0 {
		t.Errorf("Initial position get (%v,%v), want (0,0)", x, y)
	}

	v.Move(Left)
	x, y = v.Position()
	if x != -8 || y != 0 {
		t.Errorf("Left moved position get (%v,%v), want (-8,0)", x, y)
	}

	v.Move(Upper)
	x, y = v.Position()
	if x != -8 || y != -4 {
		t.Errorf("Upper moved position get (%v,%v), want (-8,-4)", x, y)
	}

	v.Move(LowerRight)
	x, y = v.Position()
	if x != 0 || y != 0 {
		t.Errorf("LowerRight moved position get (%v,%v), want (0,0)", x, y)
	}

	v.Move(Lower)
	x, y = v.Position()
	if x != 0 || y != 4 {
		t.Errorf("Lower moved position get (%v,%v), want (0,4)", x, y)
	}

	v.Move(Right)
	x, y = v.Position()
	if x != 8 || y != 4 {
		t.Errorf("Right moved position get (%v,%v), want (8,4)", x, y)
	}

	v.Move(UpperLeft)
	x, y = v.Position()
	if x != 0 || y != 0 {
		t.Errorf("UpperLeft moved position get (%v,%v), want (0,0)", x, y)
	}

	v.Move(UpperRight)
	x, y = v.Position()
	if x != 8 || y != -4 {
		t.Errorf("UpperRight moved position get (%v,%v), want (8,-4)", x, y)
	}

	v.Move(None)
	x, y = v.Position()
	if x != 8 || y != -4 {
		t.Errorf("None moved position get (%v,%v), want (8,-4)", x, y)
	}

	v.Move(LowerLeft)
	x, y = v.Position()
	if x != 0 || y != 0 {
		t.Errorf("LowerLeft moved position get (%v,%v), want (0,0)", x, y)
	}
}
