package main

import (
	"encoding/xml"
	"fmt"
)

c
func main() {
	book := Book{
		"Война и мир",
		"Толстой",
	}
	xmlData2, err := xml.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xmlData2)
	//fmt.Println(string(xmlData2))

	var book2 Book
	err = xml.Unmarshal(xmlData2, &book2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book2.title)
	fmt.Println(book2.Author)
}
