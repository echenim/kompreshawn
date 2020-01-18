package util

//ModifyBit function for transforming the bit value of the stream data
func ModifyBit(d int, position uint, state int) int {
	mask := 1 << position
	return (d &^ mask) | (-state & mask)
}
