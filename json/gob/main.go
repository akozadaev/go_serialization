package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	book := Book{
		Title:  "Go Programming Language",
		Author: "Gopher",
	}

	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(book)
	if err != nil {
		fmt.Println(err)
	}

	var book2 Book

	dec := gob.NewDecoder(&buffer)

	err = dec.Decode(&book2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Исходный %v\n", book)
	fmt.Printf("Восстановленный %v", book2)
}
