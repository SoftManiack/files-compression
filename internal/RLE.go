package internal

import (
	"fmt"
)

func LRECompression(data []byte) []byte {
	//Run length encoding

	var compressed []byte

	fmt.Printf("uncompressed length %d \n", len(data))

	fmt.Printf("compressed length %d \n", len(compressed))

	var count int16 = 1
	var value uint8 = 0

	for i := 0; i < len(data); i++ {

		if i == 0 {
			value = data[i]
		} else if value == 0 && data[i] == value {
			count++
		} else {

			compressed = append(compressed, value)
			compressed = append(compressed, count)

			value = data[i]
			count = 0
		}
	}

	fmt.Println(data)

	fmt.Println(compressed)

	return compressed

}

func LREDecompression() {

}
