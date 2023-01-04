package move

import (
	"log"
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
		{"SkillDash-to-SkillWalk-1", SkillWalk, gradually},
		{"SkillDash-to-SkillWalk-2", SkillWalk, gradually},
		{"SkillDash-to-SkillWalk-3", SkillWalk, gradually},
		{"SkillWalk-to-Ascending-1", Ascending, -1 * gradually},
		{"Ascending-to-SkillDash-1", SkillDash, speedy},
		{"SkillDash-to-Walk-1", Walk, -1 * moreSpeedy},
		{"SkillDash-to-Walk-2", Walk, -1 * moreSpeedy},
		{"SkillDash-to-Walk-3", Walk, -1 * moreSpeedy},
		{"Walk-to-Dash-1", Dash, gradually},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d.Update(tt.state)
			log.Printf("name %v, total %v, target %v", tt.name, d.currentAllOffset, d.targetOffset)
			if got := d.GetXAxisOffset(); got != tt.want {
				t.Errorf("displayOffset.GetXAxisOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
