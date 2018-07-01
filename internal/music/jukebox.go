package music

import (
	"fmt"
)

// JukeBox loads all music files and play any time you want.
type JukeBox struct {
	cd      *disc
	discMap map[DiscType]*disc
}

// NewJukeBox creates a new JukeBox instance.
// Please call the Close method when you no longer use this instance.
func NewJukeBox() (*JukeBox, error) {
	jb := JukeBox{
		discMap: make(map[DiscType]*disc),
	}
	var err, e error
	for _, dt := range DiscTypeList {
		jb.discMap[dt], e = newDisc(dt)
		if e != nil {
			err = fmt.Errorf("%v %v", err, e)
		}
		if dt == Title {
			jb.cd = jb.discMap[dt]
		}
	}
	return &jb, err
}

// NowPlayingName returns the name of the currently playing music.
// If disc is not selected or not playing, "-" is returned.
func (j *JukeBox) NowPlayingName() string {
	if j.cd == nil {
		return "nil"
	}
	if !j.cd.player.IsPlaying() {
		return "(Not playing..)"
	}
	return j.cd.name
}

// SelectDisc selects one disc by args.
func (j *JukeBox) SelectDisc(dt DiscType) error {
	disc, exist := j.discMap[dt]
	if !exist {
		return fmt.Errorf("disc '%v' is not loaded in this JukeBox", dt)
	}

	if j.cd != nil && j.cd.player.IsPlaying() {
		err := j.cd.player.Pause()
		if err != nil {
			return err
		}
	}
	j.cd = disc
	return nil
}

// Play plays a preselected disc.
func (j *JukeBox) Play() error {
	if j.cd == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if j.cd.player.IsPlaying() {
		return nil
	}
	err := j.cd.player.Rewind()
	if err != nil {
		return err
	}
	err = audioContext.Update()
	if err != nil {
		return err
	}
	return j.cd.player.Play()
}

// Pause pauses music.
func (j *JukeBox) Pause() error {
	if j.cd == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if !j.cd.player.IsPlaying() {
		return nil
	}
	return j.cd.player.Pause()
}

// Stop stops music. (pause and rewind)
func (j *JukeBox) Stop() error {
	if j.cd == nil {
		return fmt.Errorf("disc is not selected, please select disc with 'SelectDisc' method")
	}

	if !j.cd.player.IsPlaying() {
		return nil
	}
	err := j.cd.player.Pause()
	if err != nil {
		return err
	}
	return j.cd.player.Rewind()
}

// Close closes inner resources.
func (j *JukeBox) Close() error {
	var err, e error
	for i := range j.discMap {
		e = j.discMap[i].close()
		if e != nil {
			err = fmt.Errorf("%v %v", err, e)
		}
	}
	return err
}
