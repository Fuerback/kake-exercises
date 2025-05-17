package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BakeryInput struct {
	BakeryName      string                 `json:"bakeryName"`
	Characteristics []CharacteristicsInput `json:"characteristics"`
}

type CharacteristicsInput struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.json")
		return
	}

	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var input BakeryInput
	err = json.Unmarshal(data, &input)
	if err != nil {
		fmt.Println("Error unmarshalling file:", err)
		return
	}

	generateCombinations(input)
}

func generateCombinations(input BakeryInput) {
}
