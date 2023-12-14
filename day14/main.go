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
	res := 0
	for c := 0; c < len(field[0]); c++ {
		lastO := -1
		for r := 0; r < len(field); r++ {
			sym := field[r][c]
			switch sym {
			case 'O':
				newPos := r
				if newPos > lastO + 1 {
					newPos = lastO + 1
					field[newPos][c] = 'O'
					field[r][c] = '.'
				}
				lastO = newPos
				res += len(field) - newPos
			case '#':
				lastO = r
			case '.':
				continue
			}
		}
	}
	lo.ForEach(field, func (r []byte, i int) {fmt.Printf("%d: %v\n", i, string(r))})
	return res
}

func main() {
	platform := readData(getFH("input.txt"))
	lo.ForEach(platform, func (r []byte, i int) {fmt.Printf("%d: %v\n", i, string(r))})
	fmt.Println("----")
	fmt.Println("ans1:", solve(platform))
}