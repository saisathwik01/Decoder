package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"Name"`
	RollNo int    `json:"Roll"`
}

func main() {
	f, err := os.Open("output.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		panic(err)
	}
	defer f.Close()

	var p Person
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&p)
	if err != nil {
		fmt.Println("Error decoding person: ", err)
		panic(err)
	}

	fmt.Printf("Name: %s, RollNo: %d\n", p.Name, p.RollNo)
}
