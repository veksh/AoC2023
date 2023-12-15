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

func solve1(field [][]byte) int {
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

func spin(field [][]byte) int {
	// north
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
			case '#':
				lastO = r
			case '.':
				continue
			}
		}
	}
	// lo.ForEach(field, func (r []byte, i int) {fmt.Printf("north %d: %v\n", i, string(r))})
	// west
	for r := 0; r < len(field); r++ {
		lastC := -1
		for c := 0; c < len(field[0]); c++ {
			sym := field[r][c]
			switch sym {
			case 'O':
				newPos := c
				if newPos > lastC + 1 {
					newPos = lastC + 1
					field[r][newPos] = 'O'
					field[r][c] = '.'
				}
				lastC = newPos
			case '#':
				lastC = c
			case '.':
				continue
			}
		}
	}
	// lo.ForEach(field, func (r []byte, i int) {fmt.Printf("west %d: %v\n", i, string(r))})
	// south
	for c := 0; c < len(field[0]); c++ {
		lastO := len(field)
		for r := len(field)-1; r >= 0; r-- {
			sym := field[r][c]
			switch sym {
			case 'O':
				newPos := r
				if newPos < lastO - 1 {
					newPos = lastO - 1
					field[newPos][c] = 'O'
					field[r][c] = '.'
				}
				lastO = newPos
			case '#':
				lastO = r
			case '.':
				continue
			}
		}
	}
	// lo.ForEach(field, func (r []byte, i int) {fmt.Printf("south %d: %v\n", i, string(r))})
	// east
	for r := 0; r < len(field); r++ {
		lastC := len(field[0])
		for c := len(field[0])-1; c >= 0; c-- {
			sym := field[r][c]
			switch sym {
			case 'O':
				newPos := c
				if newPos < lastC - 1 {
					newPos = lastC - 1
					field[r][newPos] = 'O'
					field[r][c] = '.'
				}
				lastC = newPos
			case '#':
				lastC = c
			case '.':
				continue
			}
		}
	}
	// lo.ForEach(field, func (r []byte, i int) {fmt.Printf("east %d: %v\n", i, string(r))})
	// load on north
	res := 0
	for c := 0; c < len(field[0]); c++ {
		for r := 0; r < len(field); r++ {
			if field[r][c] == 'O' {
				res += len(field) - r
			}
		}
	}
	return res
}

func main() {
	platform := readData(getFH("input.txt"))
	lo.ForEach(platform, func (r []byte, i int) {fmt.Printf("%d: %v\n", i, string(r))})
	fmt.Println("----")
	// fmt.Println("ans1:", solve1(platform))
	for i := 1; i <= 5; i++ {
		fmt.Println("# spin:", i)
		load := spin(platform)
		lo.ForEach(platform, func (r []byte, i int) {fmt.Printf("%d: %v\n", i, string(r))})
		fmt.Println("- load:", load)
	}
}