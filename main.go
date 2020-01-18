package main

import (
	"fmt"

	bit "github.com/echenim/kompreshawn/util"
)

func main() {
	// fmt.Printf("%08b %d \n", 10, 10)
	// k := bit.FlipBit(10, 1)
	// fmt.Printf("%08b %d", k, k)

	fmt.Printf("%08b %d \n", 10, 10)
	k := bit.LowestBitNotSet(12)
	fmt.Printf("%08b %d", k, k)

}
