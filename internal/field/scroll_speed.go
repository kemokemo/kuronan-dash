package field

// ScrollSpeed is the speed of scrolling the field.
type ScrollSpeed int

const (
	// VeryFast is the fastest scrolling.
	VeryFast ScrollSpeed = iota
	// Fast is the fast scrolling.
	Fast
	// Normal is the normal scrolling.
	Normal
	// Slow is the slow scrolling.
	Slow
	// VerySlow is the slowest scrolling.
	VerySlow
)
