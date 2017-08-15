package kuronandash

import (
	"reflect"
	"testing"

	"github.com/hajimehoshi/ebiten/audio"
)

func TestNewJukeBox(t *testing.T) {
	context, err := audio.NewContext(44100)
	if err != nil {
		t.Errorf("Failed to create new audio context. %v", err)
		return
	}
	type args struct {
		con *audio.Context
	}
	tests := []struct {
		name string
		args args
		want *JukeBox
	}{
		{
			name: "normal-01",
			args: args{con: context},
			want: &JukeBox{
				context: context,
				discs:   make(map[string]*disc),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJukeBox(tt.args.con); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJukeBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFileNameWithoutExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Unix and Linux
		{name: "normal-01", args: args{path: "/path/to/file.name.hoge.name"}, want: "file.name.hoge"},
		{name: "normal-02", args: args{path: "/path/to/file.name"}, want: "file"},
		{name: "normal-03", args: args{path: "/path/to/file"}, want: "file"},
		// Windows
		{name: "normal-04", args: args{path: "C:\\path\\to\\file.name.hoge.name"}, want: "file.name.hoge"},
		{name: "normal-05", args: args{path: "C:\\path\\to\\file.name"}, want: "file"},
		{name: "normal-06", args: args{path: "C:\\path\\to\\file"}, want: "file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileNameWithoutExt(tt.args.path); got != tt.want {
				t.Errorf("getFileNameWithoutExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
