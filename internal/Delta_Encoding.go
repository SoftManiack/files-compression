package internal

import (
	"fmt"
	"unsafe"
)

func DeltaEncoding() {

	fmt.Println("delta encoding")

	input := []int64{100000001, 100000002, 100000002, 100000002, 100000002, 100000003, 100000007, 100000008, 100000010, 100000011, 100000022, 100000045, 100000009, 100000001}

	fmt.Printf("input: %T, %d\n", input, int(unsafe.Sizeof(input[0]))*len(input))

	var x int64 = input[0]

	arr := []int8{}

	for i := 1; i < len(input); i++ {

		arr = append(arr, int8(input[i]-input[0]))
	}

	fmt.Printf("n: %T, %d\n", x, unsafe.Sizeof(x))
	fmt.Printf("arr: %T, %d\n", arr, int(unsafe.Sizeof(arr[0]))*len(arr))

}
