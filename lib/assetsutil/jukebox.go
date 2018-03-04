package assetsutil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/audio"
	mp3 "github.com/hajimehoshi/go-mp3"
)

// RequestCard is used to insert discs.
type RequestCard struct {
	MusicName string
	FilePath  string
}

// JukeBox loads all music files and play any time you want.
type JukeBox struct {
	context *audio.Context
	current *disc
	discs   map[string]*disc
}

type disc struct {
	name   string
	file   *os.File
	player *audio.Player
}

// NewJukeBox creates a new JukeBox instance.
func NewJukeBox(con *audio.Context) *JukeBox {
	return &JukeBox{
		context: con,
		discs:   make(map[string]*disc),
	}
}

// InsertDiscs loads music files specified by the arguments.
// If the MusicName of the RequestCard is "", name it from the file name.
func (j *JukeBox) InsertDiscs(cards []RequestCard) error {
	for _, card := range cards {
		f, err := os.Open(card.FilePath)
		if err != nil {
			return err
		}
		s, err := mp3.NewDecoder(f)
		if err != nil {
			return err
		}
		p, err := audio.NewPlayer(j.context, s)
		if err != nil {
			return err
		}
		if card.MusicName == "" {
			card.MusicName = getFileNameWithoutExt(card.FilePath)
		}
		d := j.discs[card.MusicName]
		if d != nil {
			return fmt.Errorf("music '%s' is duplicated", card.MusicName)
		}
		j.discs[card.MusicName] = &disc{name: card.MusicName, file: f, player: p}
	}
	return nil
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// NowPlaying returns the name of the currently playing music.
// If disc is not selected or not playing, "-" is returned.
func (j *JukeBox) NowPlaying() string {
	if j.current == nil {
		return "-"
	}
	if !j.current.player.IsPlaying() {
		return "-"
	}
	return j.current.name
}

// SelectDisc selects one disc with the name specified by the argument.
func (j *JukeBox) SelectDisc(name string) error {
	disc := j.discs[name]
	if disc == nil {
		return fmt.Errorf("disc '%s' is not inserted in this JukeBox", name)
	}

	if j.current != nil && j.current.player.IsPlaying() {
		err := j.current.player.Pause()
		if err != nil {
			return err
		}
	}
	j.current = disc
	return nil
}

// Play plays a preselected disc.
func (j *JukeBox) Play() error {
	if j.current == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if j.current.player.IsPlaying() {
		return nil
	}
	err := j.current.player.Rewind()
	if err != nil {
		return err
	}
	err = j.context.Update()
	if err != nil {
		return err
	}
	return j.current.player.Play()
}

// Pause pauses music.
func (j *JukeBox) Pause() error {
	if j.current == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if !j.current.player.IsPlaying() {
		return nil
	}
	return j.current.player.Pause()
}

// Stop stops music.
func (j *JukeBox) Stop() error {
	if j.current == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if !j.current.player.IsPlaying() {
		return nil
	}
	err := j.current.player.Pause()
	if err != nil {
		return err
	}
	return j.current.player.Rewind()
}

// Close closes all resources that JukeBox used.
func (j *JukeBox) Close() error {
	var err error
	var e error
	for i := range j.discs {
		e = j.discs[i].player.Close()
		if e != nil {
			err = fmt.Errorf("%v %v", err, e)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
