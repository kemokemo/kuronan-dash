package field

// genPosSet is the settings to generate positions of objects.
type genPosSet struct {
	amount      int // Number of objects to be created
	randomRough int // Random numbers for rough placement
	randomFine  int // Random numbers for fine adjustment of the placement position
}
