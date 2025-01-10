package internal

import (
	"fmt"
)

func LRECompression(data []byte) []byte {
	//Run length encoding

	var compressed []byte

	fmt.Printf("uncompressed length %d \n", len(data))

	var count uint8 = 1
	var value uint8 = 0

	for i := 0; i < len(data); i++ {

		if i == 0 {
			value = data[i]
		} else if data[i] == value {
			count++
		} else {

			compressed = append(compressed, value)
			compressed = append(compressed, count)

			value = data[i]
			count = 0
		}
	}

	fmt.Println(data[0:10])

	fmt.Printf("compressed length %d \n", len(compressed))
	fmt.Println(compressed[0:10])

	return compressed

}

func LREDecompression(data []byte) []byte {

	return nil
}
