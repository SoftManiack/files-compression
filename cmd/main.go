package main

import (
	"bufio"
	"bytes"
	"files_compression/internal"
	"fmt"
	"os"
	"reflect"
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

	fmt.Println(reflect.TypeOf(wr))

	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	fmt.Println(reflect.TypeOf(wr))

	internal.LRECompression(wr.Read())

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
