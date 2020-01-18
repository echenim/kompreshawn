package util

//ResetBit function for reseting or clear the bit value of the data stream
func ResetBit(d int, position uint) int {
	mask := 1 << position
	return d &^ mask
}
