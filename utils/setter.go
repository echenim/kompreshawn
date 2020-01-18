package util

//SetBit function for setting the bit value of the stream data
func SetBit(d int, position uint) int {
	mask := 1 << position
	return d | mask
}
