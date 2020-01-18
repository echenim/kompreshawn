package util

//FlipBit function for flipping the determined  bit value of a dat stream
func FlipBit(x int, position uint) int {
	mask := 1 << position
	return x ^ mask
}

//LowestBitNotSet function for lowest bit not set
func LowestBitNotSet(d int) int {
	return ^d & (d + 1)
}
