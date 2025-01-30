package internal

import (
	"fmt"
	"reflect"
	"strconv"
)

func LRECompression(data []byte) []byte {
	//Run length encoding

	var compressed []byte

	var bits_data string = ""
	var bits_compressed string = ""

	for i := 0; i < len(data); i++ {

		bits_data += byteToBits(data[i])
	}

	fmt.Println(reflect.TypeOf(bits_data))
	fmt.Println(data[0:8])
	fmt.Println(reflect.TypeOf(bits_data[0]))
	fmt.Println(string(bits_data[0:8]))
	fmt.Println(string(bits_data[8:16]))

	//fmt.Println(bits_data)

	fmt.Printf("uncompressed length %d \n", len(bits_data))

	var count int = 1
	var value string = "0"

	fmt.Println(bits_data[11:100])
	for i := 0; i < len(bits_data); i++ {

		//fmt.Print(string(bits_data[i]))

		if i == 0 {
			value = string(bits_data[i])
			count += 1
		} else {
			if value != string(bits_data[i]) {

				if count > 9 {
					bits_compressed = bits_compressed + "-" + value + strconv.Itoa(count) + "-"
				} else {
					bits_compressed = bits_compressed + value + strconv.Itoa(count)
				}
				value = string(bits_data[i])
				count = 1

			} else {
				count += 1
			}
		}

		//fmt.Println(string(bits_data[i]))
	}

	fmt.Println(bits_compressed[0:100])

	fmt.Printf("compressed length %d \n", len(bits_compressed))

	// bit convert to byte

	for _, n := range bits_compressed {

		compressed = append(compressed, byte(n))
	}

	//fmt.Println(compressed[0:16])

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
