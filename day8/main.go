package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFH() *os.File {
  fh := os.Stdin
  if len(os.Args) > 1 {
    var err error
    fh, err = os.Open(os.Args[1])
    if err != nil {
      fmt.Println("error opening file:", err)
      os.Exit(1)
    }
  }
  return fh
}

func readData(fh *os.File) (turns string, graph map[string][2]string) {
  scanner := bufio.NewScanner(fh)
  scanner.Scan()
  turns = scanner.Text()
  scanner.Scan()
  graph = map[string][2]string{}
  for scanner.Scan() {
  	// AAA = (BBB, CCC)
  	// 0123456789012345
  	line := scanner.Text()
  	graph[line[0:3]] = [2]string{line[7:10], line[12:15]}
  }
	return turns, graph
}

func solve1(turns string, graph map[string][2]string) (res int) {
	fmt.Println(turns)
	fmt.Println(graph)
	return 0
}

func main() {
	turns, graph := readData(getFH())
	fmt.Println(solve1(turns, graph))
}