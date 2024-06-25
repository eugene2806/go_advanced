package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Professor struct {
	Name       string     `json:"name"`
	ScienceID  string     `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScienceID: "9798761",
		IsWorking: true,
		University: University{
			Name: "BMSTU",
			City: "Moscow",
		},
	}

	//1. Превратим профессора в последовательность байт
	ByteArr, err := json.MarshalIndent(prof1, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	json.Marshal(ByteArr)

	fmt.Println(string(ByteArr))
	err = os.WriteFile("output.json", ByteArr, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
