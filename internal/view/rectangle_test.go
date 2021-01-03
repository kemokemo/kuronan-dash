package view

import (
	"reflect"
	"testing"
)

func TestHitRectangle_GetRect(t *testing.T) {
	type fields struct {
		min Vector
		max Vector
	}

	basePos := fields{Vector{0.0, 0.0}, Vector{10.0, 10.0}}
	xPos1 := fields{Vector{15.0, 0.0}, Vector{25.0, 10.0}}
	yPos1 := fields{Vector{0.0, 15.0}, Vector{10.0, 25.0}}
	Pos1 := fields{Vector{20.1, 20.2}, Vector{40.1, 40.2}}

	tests := []struct {
		name        string
		fields1     fields
		fields2     fields
		add         Vector
		wantOverlap bool
	}{
		{"Move X-axis, not collide", basePos, xPos1, Vector{4.9, 0.0}, false},
		{"Move X-axis, collides with edge", basePos, xPos1, Vector{5.0, 0.0}, true},
		{"Move X-axis, overlaps", basePos, xPos1, Vector{8.0, 0.0}, true},
		{"Move X-axis, pass and collide with edge", basePos, xPos1, Vector{25.0, 0.0}, true},
		{"Move X-axis, pass and not collide", basePos, xPos1, Vector{25.1, 0.0}, false},

		{"Move Y-axis, not collide", basePos, yPos1, Vector{0.0, 4.9}, false},
		{"Move Y-axis, collides with edge", basePos, yPos1, Vector{0.0, 5.0}, true},
		{"Move Y-axis, overlaps", basePos, yPos1, Vector{0.0, 8.0}, true},
		{"Move Y-axis, pass and collide with edge", basePos, yPos1, Vector{0.0, 25.0}, true},
		{"Move Y-axis, pass and not collide", basePos, yPos1, Vector{0.0, 25.1}, false},

		{"Move, not collide", basePos, Pos1, Vector{10.0, 10.1}, false},
		{"Move, collides with x-edge", basePos, Pos1, Vector{10.1, 20.3}, true},
		{"Move, collides with y-edge", basePos, Pos1, Vector{16.2, 10.2}, true},
		{"Move, overlaps", basePos, Pos1, Vector{15.6, 18.9}, true},
		{"Move, pass and collide with x-edge", basePos, Pos1, Vector{40.1, 30.8}, true},
		{"Move, pass and collide with y-edge", basePos, Pos1, Vector{39.3, 40.2}, true},
		{"Move, pass and not collide", basePos, Pos1, Vector{40.2, 40.3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hr1 := NewHitRectangle(tt.fields1.min, tt.fields1.max)
			hr1.Add(tt.add)
			hr2 := NewHitRectangle(tt.fields2.min, tt.fields2.max)

			if got := hr1.Overlaps(hr2); !reflect.DeepEqual(got, tt.wantOverlap) {
				t.Errorf("HitRectangle.Overlaps() = %v, want %v", got, tt.wantOverlap)
			}
		})
	}
}
