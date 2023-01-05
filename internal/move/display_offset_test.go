package move

import (
	"testing"
)

func Test_displayOffset_GetXAxisOffset(t *testing.T) {
	d := &displayOffset{}
	tests := []struct {
		name  string
		state State
		want  float64
	}{
		{"Dash-1", Dash, 0.0},
		{"Dash-2", Dash, 0.0},
		{"Dash-to-Walk-1", Walk, -1 * gradually},
		{"Dash-to-Walk-2", Walk, -1 * gradually},
		{"Dash-to-Walk-3", Walk, -1 * gradually},
		{"Dash-to-Walk-4", Walk, -1 * gradually},
		{"Dash-to-Walk-5", Walk, -1 * gradually},
		{"Dash-to-Walk-6", Walk, -1 * gradually},
		{"Walk-to-SkillDash-1", SkillEffect, -1 * gradually}, // 無視するStateなので直前と同じ値になる
		{"Walk-to-SkillDash-2", SkillDash, moreSpeedy},
		{"Walk-to-SkillDash-3", SkillDash, moreSpeedy},
		{"Walk-to-SkillDash-4", SkillDash, moreSpeedy},
		{"Walk-to-SkillDash-5", SkillDash, moreSpeedy},
		{"Walk-to-SkillDash-6", SkillDash, moreSpeedy},
		{"Walk-to-SkillDash-7", SkillDash, moreSpeedy},
		{"SkillDash-to-SkillWalk-1", SkillWalk, -1 * gradually},
		{"SkillDash-to-SkillWalk-2", SkillWalk, -1 * gradually},
		{"SkillDash-to-SkillWalk-3", SkillWalk, -1 * gradually},
		{"SkillWalk-to-SkillAscending-1", SkillAscending, gradually},
		{"SkillWalk-to-SkillAscending-2", SkillAscending, gradually},
		{"SkillWalk-to-SkillAscending-3", SkillAscending, gradually},
		{"SkillWalk-to-SkillAscending-4", SkillAscending, gradually},
		{"SkillWalk-to-SkillAscending-5", SkillAscending, gradually},
		{"SkillAscending-to-SkillDash-1", SkillDash, 0},
		{"SkillDash-to-Walk-1", Walk, -1 * gradually},
		{"SkillDash-to-Walk-2", Walk, -1 * gradually},
		{"SkillDash-to-Walk-3", Walk, -1 * gradually},
		{"SkillDash-to-Walk-4", Walk, -1 * gradually},
		{"SkillDash-to-Walk-5", Walk, -1 * gradually},
		{"SkillDash-to-Walk-6", Walk, -1 * gradually},
		{"Walk-to-Dash-1", Dash, -1 * gradually},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d.Update(tt.state)
			if got := d.GetXAxisOffset(); got != tt.want {
				t.Errorf("displayOffset.GetXAxisOffset() = %v, want %v", got, tt.want)
				t.Errorf("target %v, here %v", d.targetOffset, d.currentAllOffset)
			}
		})
	}
}
