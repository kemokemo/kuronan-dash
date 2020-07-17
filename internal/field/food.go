package field

// Food is the interface for filed food item. The character can eat them and restore stamina.
type Food interface {
	// Eat eats this food. This func reteruns the value to restore character's stamina.
	Eat() int
}
