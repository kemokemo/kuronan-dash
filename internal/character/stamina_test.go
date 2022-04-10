package character

import (
	"reflect"
	"testing"

	"github.com/kemokemo/kuronan-dash/internal/move"
)

func TestNewStamina(t *testing.T) {
	type args struct {
		max       int
		endurance int
	}
	tests := []struct {
		name string
		args args
		want *Stamina
	}{
		{"normal", args{50, 5}, &Stamina{max: 50, endurance: 5, val: 50, valRate: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStamina(tt.args.max, tt.args.endurance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStamina() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStamina_Restore(t *testing.T) {
	type fields struct {
		max       int
		val       int
		endurance int
		valRate   int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"normal", fields{100, 60, 5, 4}, args{5}, 65},
		{"too much", fields{100, 98, 5, 5}, args{10}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stamina{
				max:       tt.fields.max,
				val:       tt.fields.val,
				endurance: tt.fields.endurance,
				valRate:   tt.fields.valRate,
			}
			s.Add(tt.args.val)
			got := s.GetStamina()
			if got != tt.want {
				t.Errorf("GetStamina() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStamina_Consumes(t *testing.T) {
	type fields struct {
		max       int
		endurance int
	}
	type args struct {
		state move.State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  int
		want2  int
	}{
		{"Walk", fields{100, 2}, args{move.Walk}, 100, 99},
		{"Walk2", fields{100, 3}, args{move.Walk}, 100, 100},
		{"Dash", fields{100, 2}, args{move.Dash}, 99, 98},
		{"Dash2", fields{100, 3}, args{move.Dash}, 100, 99},
		{"Ascending (not consume)", fields{100, 2}, args{move.Ascending}, 100, 100},
		{"Descending (not consume)", fields{100, 2}, args{move.Descending}, 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStamina(tt.fields.max, tt.fields.endurance)
			s.ConsumesByState(tt.args.state)
			got := s.GetStamina()
			if got != tt.want1 {
				t.Errorf("GetStamina() = %v, want1 %v", got, tt.want1)
			}

			s.ConsumesByState(tt.args.state)
			got = s.GetStamina()
			if got != tt.want2 {
				t.Errorf("GetStamina() = %v, want2 %v", got, tt.want2)
			}

			s.Initialize()

			s.ConsumesByState(tt.args.state)
			got = s.GetStamina()
			if got != tt.want1 {
				t.Errorf("GetStamina() = %v, want1 %v", got, tt.want1)
			}

			s.ConsumesByState(tt.args.state)
			got = s.GetStamina()
			if got != tt.want2 {
				t.Errorf("GetStamina() = %v, want2 %v", got, tt.want2)
			}
		})
	}
}
