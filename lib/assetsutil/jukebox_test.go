package assetsutil

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/hajimehoshi/ebiten/audio"
)

func TestNewJukeBox(t *testing.T) {
	type args struct {
		con *audio.Context
	}
	tests := []struct {
		name string
		args args
		want *JukeBox
	}{
		{name: "normal-01", args: args{con: testContext}, want: &JukeBox{context: testContext, discs: make(map[string]*disc)}},
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
		{name: "normal-01", args: args{path: filepath.Join("path", "to", "file.name.hoge.name")}, want: "file.name.hoge"},
		{name: "normal-02", args: args{path: filepath.Join("path", "to", "file.name")}, want: "file"},
		{name: "normal-03", args: args{path: filepath.Join("path", "to", "file")}, want: "file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileNameWithoutExt(tt.args.path); got != tt.want {
				t.Errorf("getFileNameWithoutExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertDiscs(t *testing.T) {
	juke := NewJukeBox(testContext)
	type args struct {
		cards []RequestCard
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Normal", args: args{cards: []RequestCard{
			RequestCard{
				MusicName: "stage01",
				FilePath:  "./TestData/music/sample01.mp3",
			},
		}},
			wantErr: false,
		},
		{name: "File is not found", args: args{cards: []RequestCard{
			RequestCard{
				FilePath: "./TestData/music/sample00.mp3",
			},
		}},
			wantErr: true,
		},
		{name: "MP3 file is corrupted", args: args{cards: []RequestCard{
			RequestCard{
				FilePath: "./TestData/music/corrupted01.mp3",
			},
		}},
			wantErr: true,
		},
		{name: "MusicName is duplicated", args: args{cards: []RequestCard{
			RequestCard{
				MusicName: "sample",
				FilePath:  "./TestData/music/sample01.mp3",
			},
			RequestCard{
				MusicName: "sample",
				FilePath:  "./TestData/music/sample02.mp3",
			},
		}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := juke.InsertDiscs(tt.args.cards); (err != nil) != tt.wantErr {
				t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// MusicName will be fila name
	err := juke.InsertDiscs([]RequestCard{
		RequestCard{
			FilePath: "./TestData/music/sample01.mp3",
		}})
	if err != nil {
		t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.SelectDisc("sample01")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
}

func TestNowPlaying(t *testing.T) {
	juke := NewJukeBox(testContext)
	err := juke.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "stage01", FilePath: "./TestData/music/sample01.mp3"},
	})
	if err != nil {
		t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.SelectDisc("stage01")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Play()
	if err != nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, false)
		return
	}

	notPlaying := NewJukeBox(testContext)
	err = notPlaying.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "stage02", FilePath: "./TestData/music/sample02.mp3"},
	})
	if err != nil {
		t.Errorf("notPlaying.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	err = notPlaying.SelectDisc("stage02")
	if err != nil {
		t.Errorf("notPlaying.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}

	notSelected := NewJukeBox(testContext)
	err = notSelected.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "sample", FilePath: "./TestData/music/sample02.mp3"},
	})
	if err != nil {
		t.Errorf("notSelected.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}

	tests := []struct {
		name string
		juke *JukeBox
		want string
	}{
		{name: "Normal playing", juke: juke, want: "stage01"},
		{name: "Not playing", juke: notPlaying, want: "-"},
		{name: "Not selected", juke: notSelected, want: "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.juke.NowPlaying(); got != tt.want {
				t.Errorf("juke.NowPlaying() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectDisc(t *testing.T) {
	juke := NewJukeBox(testContext)
	err := juke.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "stage01", FilePath: "./TestData/music/sample01.mp3"},
		RequestCard{MusicName: "stage02", FilePath: "./TestData/music/sample02.mp3"},
	})
	if err != nil {
		t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.SelectDisc("stage01")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Play()
	if err != nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.SelectDisc("stage02")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
	// error pattern
	err = juke.SelectDisc("boss")
	if err == nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, true)
		return
	}
}

func TestPlay(t *testing.T) {
	juke := NewJukeBox(testContext)
	err := juke.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "stage01", FilePath: "./TestData/music/sample01.mp3"},
		RequestCard{MusicName: "stage02", FilePath: "./TestData/music/sample02.mp3"},
	})
	if err != nil {
		t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	// Not selected error
	err = juke.Play()
	if err == nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, true)
		return
	}
	err = juke.Pause()
	if err == nil {
		t.Errorf("juke.Pause() error = %v, wantErr %v", err, true)
		return
	}
	err = juke.Stop()
	if err == nil {
		t.Errorf("juke.Stop() error = %v, wantErr %v", err, true)
		return
	}

	err = juke.SelectDisc("stage01")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Play()
	if err != nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Pause()
	if err != nil {
		t.Errorf("juke.Pause() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Play()
	if err != nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Stop()
	if err != nil {
		t.Errorf("juke.Stop() error = %v, wantErr %v", err, false)
		return
	}
}

func TestClose(t *testing.T) {
	juke := NewJukeBox(testContext)
	err := juke.InsertDiscs([]RequestCard{
		RequestCard{MusicName: "stage01", FilePath: "./TestData/music/sample01.mp3"},
		RequestCard{MusicName: "stage02", FilePath: "./TestData/music/sample02.mp3"},
	})
	if err != nil {
		t.Errorf("juke.InsertDiscs() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.SelectDisc("stage01")
	if err != nil {
		t.Errorf("juke.SelectDisc() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Play()
	if err != nil {
		t.Errorf("juke.Play() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Stop()
	if err != nil {
		t.Errorf("juke.Stop() error = %v, wantErr %v", err, false)
		return
	}
	err = juke.Close()
	if err != nil {
		t.Errorf("juke.Close() error = %v, wantErr %v", err, false)
		return
	}
}
