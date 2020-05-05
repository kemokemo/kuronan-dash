package field

// genPosSet is the settings to generate positions of objects.
type genPosSet struct {
	amount      int // Number of objects to be created
	randomRough int // Random numbers for rough placement
	randomFine  int // Random numbers for fine adjustment of the placement position
}

// genVelSet is the settings to generate velocities of objects.
type genVelSet struct {
	x      float64
	y      float64
	random bool
}
