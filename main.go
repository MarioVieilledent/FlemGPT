package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Reading file...")
	content := readFile("input.txt")

	learn(content)

	fmt.Println("Writing JSON output...")
	writeFile("output.json", string(jsonOutput))

	fmt.Println("Generation...")
	playGround()
}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(data)
}

func writeFile(filename string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
