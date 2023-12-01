package main

import (
	"bufio"
	"log"
	"os"
)

const INPUT = "../input.txt"

func main() {
	log.Print("start")
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("got %s", line)
	}
	log.Print("done")
}