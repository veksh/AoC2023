package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/samber/lo"
)

func getFH(defFileName string) *os.File {
	if defFileName == "" && len(os.Args) == 1 {
		return os.Stdin
	}
  fileName := defFileName
  if len(os.Args) > 1 {
  	fileName = os.Args[1]
  }
  fh, err := os.Open(fileName)
  if err != nil {
    fmt.Println("error opening file:", err)
    os.Exit(1)
  }
  return fh
}

func readData(fh *os.File) [][]int {
  scanner := bufio.NewScanner(fh)
  res := [][]int{}
  for scanner.Scan() {
  	res = append(res, lo.Map([]byte(scanner.Text()), func (b byte, _ int) int {return int(b - '0')} ))
  }
	return res
}

func solve1(field [][]int) int {
	return len(field)
}

func main() {
	maze := readData(getFH("input.txt"))
	fmt.Println(maze)
	fmt.Println("ans1:", solve1(maze))
}