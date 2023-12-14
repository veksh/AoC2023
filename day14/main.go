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

func readData(fh *os.File) [][]byte {
  scanner := bufio.NewScanner(fh)
  res := [][]byte{}
  for scanner.Scan() {
  	// "." == 35, "O" == 79, "#" == 35
  	res = append(res, []byte(scanner.Text()))
  }
	return res
}

func solve(field [][]byte) int {
	return 0
}

func main() {
	platform := readData(getFH("input.txt"))
	lo.ForEach(platform, func (r []byte, i int) {fmt.Printf("%d: %v\n", i, string(r))})

	// outer:
	// for i := 0; i < len(starmap[0]); i++ {
	// 	for j := 0; j < len(starmap); j++ {
	// 		if starmap[j][i] == '#' {
	// 			continue outer
	// 		}
	// 	}
	// 	emptycols = append(emptycols, i)
	// }
	fmt.Println("ans1:", solve(platform))
}