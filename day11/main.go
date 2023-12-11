package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
  starmap := [][]byte{}
  for scanner.Scan() {
  	// void "." == 35, star "#" == 46
  	starmap = append(starmap, []byte(scanner.Text()))
  }
	return starmap
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1*a
}

func main() {
	starmap := readData(getFH("input.txt"))
	emptycols := []int{}

	outer:
	for i := 0; i < len(starmap[0]); i++ {
		for j := 0; j < len(starmap); j++ {
			if starmap[j][i] == '#' {
				continue outer
			}
		}
		emptycols = append(emptycols, i)
	}
	fmt.Println("emptycols:", emptycols)

  // {row, col} ie {y, x}
	allstars := [][2]int{}
	radd := 0
	for r := 0; r < len(starmap); r++ {
		if slices.Index(starmap[r], '#') == -1 {
			radd += 1
			continue
		}
		cadd := 0
		for c := 0; c < len(starmap[0]); c++ {
			if starmap[r][c] == '#' {
				allstars = append(allstars, [2]int{r + radd, c + cadd})
			} else {
				if slices.Index(emptycols, c) != -1 {
					cadd += 1
				}
			}
		}
	}
	fmt.Println("allstars:", allstars)

	res1 := 0
	for i, s1 := range(allstars) {
		for _, s2 := range(allstars[i+1:]) {
			res1 += abs(s2[0] - s1[0]) + abs(s2[1] - s1[1])
		}
	}
	fmt.Println("ans 1:", res1)
}