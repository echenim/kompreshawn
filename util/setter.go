package util

//SetBit function for setting the bit value of the stream data
func SetBit(d int, position uint) int {
	mask := 1 << position
	return d | mask
}

//IsBitSet function check if the bit value is set and if true return the appropriate value
func IsBitSet(d int, position uint) bool {
	d >>= position
	return (d & 1) != 0
}
