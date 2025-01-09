package main

import (
	internal "files_compression/internal"
	"fmt"
)

func main() {

	data, err := internal.ReadByteFile("banana.jpeg")

	if err != nil {
		fmt.Println(err)
	} else {
		_ := internal.LRECompression(data)
	}

	//internal.LRECompression()

	/*
		/wr1 := bytes.Buffer{}
		newData := []byte("New data")
		wr1.Write(newData)

		p := make([]byte, 0)
		wr1.Read(p)

		fmt.Println(p)
	*/

	//internal.LRECompression(data)

}

// Разобарться в чтении байтов
