package images

func loadFieldparts() error {
	var err error

	// Prairie field
	TilePrairie, err = loadSingleImage(tilePrairie_png)
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
	MountainNear, err = loadSingleImage(mountainNear_png)
	if err != nil {
		return err
	}
	MountainFar, err = loadSingleImage(mountainFar_png)
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

	return nil
}
