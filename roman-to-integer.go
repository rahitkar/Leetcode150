package main

import (
	"fmt"
)

var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// with split func
//func romanToInt(str string) int {
//	splitStr := strings.Split(str, "")
//	value := 0
//	lastIndex := len(splitStr) - 1
//	for i := 0; i < lastIndex; i++ {
//		int1 := symbol[splitStr[i]]
//		int2 := symbol[splitStr[i+1]]
//		if int1 >= int2 {
//			value += int1
//		} else {
//			if int1*5 == int2 || int1*10 == int2 {
//				value -= int1
//			}
//		}
//	}
//	return value + symbol[splitStr[lastIndex]]
//}

func getSymbolValue(val uint8) int {
	return symbol[fmt.Sprintf("%c", val)]
}

func romanToInt(str string) int {
	lastIndex := len(str) - 1
	value := getSymbolValue(str[lastIndex])
	for i := 0; i < lastIndex; i++ {
		int1 := getSymbolValue(str[i])
		int2 := getSymbolValue(str[i+1])
		if int1 >= int2 {
			value += int1
		} else {
			if int1*5 == int2 || int1*10 == int2 {
				value -= int1
			}
		}
	}
	return value
}

func main() {
	println(romanToInt("II"))
	println(romanToInt("XII"))
	println(romanToInt("XVII"))
	println(romanToInt("XXVII"))
	println(romanToInt("XIV"))
	println(romanToInt("CMM"))
}
