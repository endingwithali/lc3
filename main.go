package main

import (
	"fmt"
)

func sext(value uint16, signIndex uint16) uint16 {
	if (value>>signIndex)&1 == 0 {
		return value
	}
	mask := uint16(0)
	for i := signIndex; i < 16; i++ {
		mask |= 1 << i
	}
	return mask | value
}

func sign_extend(sign_bit_idx uint16, input_value uint16) uint16 {

	if (input_value>>sign_bit_idx)&1 == 0 {
		return input_value
	}
	mask := uint16(0)
	for i := sign_bit_idx; i < 16; i++ {
		mask |= 1 << i
	}
	return mask | input_value
}

func main() {
	test1 := uint16(0b10001)
	result := sext(test1, 4)
	result2 := sign_extend(4, test1)
	fmt.Printf("value: %b\n", test1)
	fmt.Printf("SEXT value: %b\n", result)
	fmt.Printf("signExtend value: %b\n", result2)
}

// func main() {
// 	testValue := 0b1011
// 	sextVal := sext(testValue, 4)
// 	fmt.Printf("value: %b\n", testValue)
// 	fmt.Printf("SEXT value: %b\n", sextVal)
// 	fmt.Printf("expected value: 1111111111111011\n")
// 	testValue2 := 0b0011
// 	sextVal2 := sext(testValue2, 4)
// 	fmt.Printf("value: %b\n", testValue2)
// 	fmt.Printf("SEXT value: %b\n", sextVal2)
// 	fmt.Printf("expected value: 11\n")
// }

// func sext(value int, len int) int {
// 	maskedValue := value & 0xFFFF
// 	shift := uint(16 - len)
// 	return (maskedValue << shift) >> shift

// }

// testing bitmapings
func bitmaptests() {
	testValue := 0b0000000011111111
	firstFour := bitRangeExtraction(testValue, 0, 3)
	lastFour := bitRangeExtraction(testValue, 11, 15)
	fifthToLast := bitExtraction(testValue, 10)
	fmt.Printf("value: %b\n", testValue)
	fmt.Printf("firstFour value: %b\n", firstFour)
	fmt.Printf("Last Four value: %b\n", lastFour)
	fmt.Printf("Fifth to lastvalue: %b\n", fifthToLast)
	// fmt.Print(m[0:4] == "hell")

}

func bitRangeExtraction(value int, startIndex int, endIndex int) int {
	maskedValue := value & 0xFFFF
	mask := (1 << (endIndex - startIndex + 1)) - 1
	return (maskedValue >> startIndex) & mask
}
func bitExtraction(value int, index int) int {
	maskedValue := value & 0xFFFF
	return (maskedValue >> (15 - index)) & 1
}

/*

learning bits
var number int16
	// number = 0b11100001
	// // fmt.Println(number << 3)
	// output := number >> 4
	// fmt.Printf("%b\n", number)




	Trying to understand bitestrings
		var holder bitset.BitSet
	var holder2 bitset.BitSet
	var number int16
	var num1 int16
	number = 0b111000011110000
	num1 = 0b1
	holder2 = *holder2.Set(uint(num1))
	holder = *holder.Set(uint(number))
	fmt.Println("test")
	fmt.Println(holder.String())
	fmt.Println(holder2.DumpAsBits())
	fmt.Printf(holder2[0])

*/
