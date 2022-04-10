package character

import (
	"testing"

	"github.com/kemokemo/kuronan-dash/internal/move"
)

func TestTension_Add(t *testing.T) {
	type fields struct {
		max    int
		border int
	}
	type args struct {
		val1 move.State
		val2 move.State
		want int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"dash", fields{max: 10, border: 2}, args{val1: move.Dash, val2: move.Dash, want: 2}},
		{"over", fields{max: 1, border: 2}, args{val1: move.Dash, val2: move.Dash, want: 1}},
		{"walk", fields{max: 10, border: 2}, args{val1: move.Walk, val2: move.Walk, want: 1}},
		{"Ascending (not add)", fields{max: 10, border: 2}, args{val1: move.Ascending, val2: move.Ascending, want: 0}},
		{"Descending (not add)", fields{max: 10, border: 2}, args{val1: move.Descending, val2: move.Descending, want: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTension(tt.fields.max, tt.fields.border)
			tr.AddByState(tt.args.val1)
			tr.AddByState(tt.args.val2)
			got := tr.Get()
			if got != tt.args.want {
				t.Errorf("Get(): %v, want: %v", got, tt.args.want)
			}
		})
	}
}

func TestTension_AddByAttack(t *testing.T) {
	type fields struct {
		max     int
		languor int
	}
	type args struct {
		brokenNum int
		want      int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"normal", fields{max: 30, languor: 2}, args{brokenNum: 5, want: 12}},
		{"normal-2", fields{max: 10, languor: 2}, args{brokenNum: 1, want: 2}},
		{"over", fields{max: 10, languor: 2}, args{brokenNum: 5, want: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTension(tt.fields.max, tt.fields.languor)
			tr.AddByAttack(tt.args.brokenNum)
			got := tr.Get()
			if got != tt.args.want {
				t.Errorf("Get(): %v, want: %v", got, tt.args.want)
			}
		})
	}
}
