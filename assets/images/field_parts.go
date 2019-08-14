package images

func loadFieldparts() error {
	var err error

	TilePrairie, err = loadSingleImage(tilePrairie_png)
	if err != nil {
		return err
	}
	Prairies, err = loadSingleImage(prairies_png)
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

	return nil
}
