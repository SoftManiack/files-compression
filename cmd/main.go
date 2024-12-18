package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./images/banana.jpeg")

	fmt.Printf("Type %T", file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	fmt.Println(wr)

	//internal.LRECompression(data)

}
