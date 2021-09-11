package images

func loadBackground() error {
	var err error

	TitleBackground, err = loadSingleImage(bg_title_png)
	if err != nil {
		return err
	}
	SelectBackground, err = loadSingleImage(bg_select_png)
	if err != nil {
		return err
	}
	SkyBackground, err = loadSingleImage(bg_prairie_png)
	if err != nil {
		return err
	}

	return nil
}
