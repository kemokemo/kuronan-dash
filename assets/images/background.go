package images

func loadBackground() error {
	var err error

	TitleBackground, err = loadSingleImage(title_bg_png)
	if err != nil {
		return err
	}
	SelectBackground, err = loadSingleImage(select_bg_png)
	if err != nil {
		return err
	}
	SkyBackground, err = loadSingleImage(sky_bg_png)
	if err != nil {
		return err
	}

	return nil
}
