package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json: "name"`
	Age  string `json: "age"`
}

func main() {
	p := &Person{"Alexey", "45"}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Println("Marshal error: ", err)
	}

	fmt.Println("Data: ", jsonData)
	//fmt.Println("Data: ", string(jsonData))

	var p2 = Person{"Alexey", "45"}
	fmt.Println(p2)
	jsonData2, err := json.Marshal(p2)
	fmt.Println("Data: ", string(jsonData2))

	err = json.Unmarshal(jsonData, p)
	if err != nil {
		log.Println("Unmarshal error: ", err)
	}
	fmt.Println("Data: ", p)
}
