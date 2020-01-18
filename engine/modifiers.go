package engine

//BitModifier for transforming the bit value of the stream
func BitModifier(d int, position uint, state int) int {
	mask := 1 << position
	return (d &^ mask) | (-state & mask)
}
