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

// row, col; RDLU
var RDLU = [][2]int {
	{ 0,  1}, // R 0
	{ 1,  0}, // D 1
	{ 0, -1}, // L 2
	{-1,  0}, // U 3
}

type cell struct {
	row  int
	col  int
	dir  byte
	run  byte
}

const MAXRUN = 3
var VOID = struct{}{}

func solve1(field [][]int) int {
	maxR, maxC := len(field)-1, len(field[0])-1
	seen := map[cell]int{}
	queue := map[cell]int {
		{row: 0, col: 1, dir: 0, run: 1}: 0,
		{row: 1, col: 0, dir: 1, run: 1}: 0,
	}
	res := 1_000_000
	for len(queue) > 0 {
		newq := map[cell]int{}
		for c,cost := range(queue) {
			if c.row < 0 || c.row > maxR || c.col < 0 || c.col > maxC {
				continue
			}
			cost += field[c.row][c.col]
			if prevcost, ok := seen[c]; ok {
				if prevcost < cost {
					continue
				}
			}
			seen[c] = cost
			if c.row == maxR && c.col == maxR {
				fmt.Println("reached, cost", cost)
				res = min(res, cost)
				continue
			}
			newnei := []cell{}
			if c.run < MAXRUN {
				newnei = append(newnei, cell{
					row: c.row + RDLU[c.dir][0],
					col: c.col + RDLU[c.dir][1],
					dir: c.dir,
					run: c.run + 1,
				})
			}
			for _, turn := range([]byte{(c.dir + 1) % 4, (c.dir + 3) % 4}) {
				newnei = append(newnei, cell{
						row: c.row + RDLU[turn][0],
						col: c.col + RDLU[turn][1],
						dir: turn,
						run: 1,
				})
			}
			for _, nn := range(newnei) {
				if prevcost, ok := newq[nn]; ok {
					if prevcost < cost {
						continue
					}
				}
				newq[nn] = cost
			}
		}
		queue = newq
	}
	return res
}

func main() {
	maze := readData(getFH("input.txt"))
	fmt.Println(maze)
	fmt.Println("ans1:", solve1(maze))
}