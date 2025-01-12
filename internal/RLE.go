package internal

import (
	"fmt"
)

func LRECompression(data []byte) []byte {
	//Run length encoding

	var compressed []byte

	var bits_data string
	var bits_compressed string

	for i := 0; i < len(data); i++ {

		bits_data += byteToBits(data[i])
	}

	fmt.Println(data[0:8])
	fmt.Println(bits_data[0])

	//fmt.Println(bits_data)

	fmt.Printf("uncompressed length %d \n", len(bits_data))

	var count uint8 = 1
	var value uint8 = 0

	for i := 0; i < 10; i++ {

		fmt.Println("*")
		fmt.Println(bits_data[i])
		if i == 0 {
			value = bits_data[i]
			count--
		} else {
			count++
			bits_compressed = bits_compressed + string(value) + string(count)

			value = bits_data[i]
			count = 0
		}
	}

	fmt.Printf("compressed length %d \n", len(bits_compressed))
	fmt.Println(bits_compressed)

	return compressed

}

func LREDecompression(data []byte) []byte {

	return nil
}

func byteToBits(b byte) string {
	var bits string
	for i := 7; i >= 0; i-- {
		if (b & (1 << i)) != 0 {
			bits += "1"
		} else {
			bits += "0"
		}
	}
	return bits
}
