package internal

import (
	"fmt"
)

func LRECompression(data []byte) []byte {
	//Run length encoding

	var compressed []byte

	var bits_data []bits
	var bits_compressed []bits

	fmt.Printf("uncompressed length %d \n", len(bits_data))

	var count uint8 = 1
	var value uint8 = 0

	for i := 0; i < len(bits_data); i++ {

		if i == 0 {
			value = bits_data[i]
			count--
		} else {
			count++
			bits_compressed = append(bits_compressed, value)
			bits_compressed = append(bits_compressed, count)

			value = bits_data[i]
			count = 0
		}
	}

	fmt.Printf("compressed length %d \n", len(bits_compressed))

	return compressed

}

func LREDecompression(data []byte) []byte {

	return nil
}
