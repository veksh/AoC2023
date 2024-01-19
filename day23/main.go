package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readMaze() []string {
	fh, err := os.Open(os.Args[1]); if err != nil {log.Fatal("cannot open file", err)}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	maze := []string{}
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func main() {
	maze := readMaze()
	fmt.Println(strings.Join(maze, "\n"))
}