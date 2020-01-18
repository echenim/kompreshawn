package engine

//SetBit function for setting the bit value
func SetBit(d int, position uint) int {
	mask := 1 << position
	return d | mask
}
