package util

//SetBit function for setting the bit value of the stream data
func SetBit(d int, position uint) int {
	mask := 1 << position
	return d | mask
}

//ModifyBit function for transforming the bit value of the stream data
func ModifyBit(d int, position uint, state int) int {
	mask := 1 << position
	return (d &^ mask) | (-state & mask)
}

//ResetBit function for reseting or clear the bit value of the data stream
func ResetBit(d int, position uint) int {
	mask := 1 << position
	return d &^ mask
}
