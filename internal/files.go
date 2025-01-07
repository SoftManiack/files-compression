package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
)

func ReadByteFile(name string) ([]byte, error) {

	path := fmt.Sprintf("./images/%s", name)

	file, err := os.Open(path)

	//fmt.Printf("Type %T", file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	wr := bytes.Buffer{}

	sc := bufio.NewScanner(file)

	//fmt.Println(reflect.TypeOf(wr))

	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	data := make([]byte, wr.Len())

	fmt.Println(reflect.TypeOf(data))

	_, err = wr.Read(data)

	fmt.Println(reflect.TypeOf(data))

	if err != nil {
		return nil, err
	}

	return data, nil
}
