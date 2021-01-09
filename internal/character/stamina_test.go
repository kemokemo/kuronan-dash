package character

import (
	"reflect"
	"testing"
)

func TestStaminaConsumes(t *testing.T) {
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
		{"normal", fields{100, 90, 2, 1}, args{1}, 89},
		{"not reduced", fields{100, 50, 2, 2}, args{1}, 50},
		{"big consume", fields{90, 20, 2, 2}, args{3}, 19},
		{"zero", fields{90, 0, 2, 1}, args{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stamina{
				max:       tt.fields.max,
				val:       tt.fields.val,
				endurance: tt.fields.endurance,
				valRate:   tt.fields.valRate,
			}
			s.consumes(tt.args.val)
			got := s.GetStamina()
			if got != tt.want {
				t.Errorf("GetStamina() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
