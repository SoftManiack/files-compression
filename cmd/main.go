package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./images/banana.jpeg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 64)

	internal.LRECompression(data)

}
