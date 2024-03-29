package images

func loadFieldParts() error {
	var err error

	// General
	Goal_back, err = loadSingleImage(goal_back_png)
	if err != nil {
		return err
	}
	Goal_front, err = loadSingleImage(goal_front_png)
	if err != nil {
		return err
	}
	MapBackground, err = loadSingleImage(map_background_png)
	if err != nil {
		return err
	}

	// Prairie field
	TilePrairie, err = loadSingleImage(tile_prairie_png)
	if err != nil {
		return err
	}
	Grass1, err = loadSingleImage(grass1_png)
	if err != nil {
		return err
	}
	Grass2, err = loadSingleImage(grass2_png)
	if err != nil {
		return err
	}
	Grass3, err = loadSingleImage(grass3_png)
	if err != nil {
		return err
	}
	MountainNear, err = loadSingleImage(mountain_near_png)
	if err != nil {
		return err
	}
	MountainFar, err = loadSingleImage(mountain_far_png)
	if err != nil {
		return err
	}
	CloudNear, err = loadSingleImage(cloud_near_png)
	if err != nil {
		return err
	}
	CloudFar, err = loadSingleImage(cloud_far_png)
	if err != nil {
		return err
	}
	RockNormal, err = loadSingleImage(rock_normal_png)
	if err != nil {
		return err
	}

	// UI
	CharaWindow, err = loadSingleImage(chara_window_png)
	if err != nil {
		return err
	}
	CharaSelectButton, err = loadSingleImage(chara_select_button_png)
	if err != nil {
		return err
	}
	PauseButton, err = loadSingleImage(pause_button_png)
	if err != nil {
		return err
	}
	StartButton, err = loadSingleImage(start_button_png)
	if err != nil {
		return err
	}
	UpButton, err = loadSingleImage(up_button_png)
	if err != nil {
		return err
	}
	DownButton, err = loadSingleImage(down_button_png)
	if err != nil {
		return err
	}
	AttackButton, err = loadSingleImage(attack_button_png)
	if err != nil {
		return err
	}
	SkillButton, err = loadSingleImage(skill_button_png)
	if err != nil {
		return err
	}
	VolumeOnButton, err = loadSingleImage(button_volume_on_png)
	if err != nil {
		return err
	}
	VolumeOffButton, err = loadSingleImage(button_volume_off_png)
	if err != nil {
		return err
	}
	StartTitleButton, err = loadSingleImage(button_start_title_png)
	if err != nil {
		return err
	}
	RatingStar, err = loadSingleImage(star_png)
	if err != nil {
		return err
	}

	return nil
}
