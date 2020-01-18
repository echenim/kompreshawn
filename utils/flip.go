package util

//FlipBit function for flipping the determined  bit value of a dat stream
func FlipBit(x int, position uint) int {
	mask := 1 << position
	return x ^ mask
}
