package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	fileName := fmt.Sprintf("data/day%s_%s.txt", args[0], args[1])
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch args[0] {
	case "1":
		output := dayOne(scanner)	
		if args[1] == "test" {
			if output != 11 {
				log.Fatalf("Invalid: expected 11, got %d", output)
			} else {
				log.Printf("You got it tiger! Output: %d", output)
			}
		} else {
			log.Printf("You got it tiger! Output: %d", output)
		}
	}
}
