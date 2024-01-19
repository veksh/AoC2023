package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rc  struct {r, c int}
type drc struct {dr, dc int}

func (r rc)add(d drc) rc {
	return rc{r.r + d.dr, r.c + d.dc}
}

func readMaze() []string {
	fName := "input_test.txt"
	if len(os.Args) >= 2 {
		fName = os.Args[1]
	}
	fh, err := os.Open(fName); if err != nil {log.Fatal("cannot open file", err)}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	maze := []string{}
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

// values also used as steps/directions
var slides = map[byte]drc {
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
	'^': {-1, 0},
}

func getNeigh(maze []string, pos rc) []rc {
	sym := maze[pos.r][pos.c]
	// mandatory slide
	if d, ok := slides[sym]; ok {
		return []rc{pos.add(d)}
	}
	res := []rc{}
	for s, d := range(slides) {
		n := pos.add(d)
		if n.r > len(maze)-1 || n.r < 0 || n.c > len(maze[0])-1 || n.c < 0 {
			continue
		}
		ns := maze[n.r][n.c]
		if ns == '#' {
			continue
		}
		// do not go against the slide
		if (s == '<' && ns == '>') || (s == '>' && ns == '<') || (s == '^' && ns == 'v') || (s == 'v' && ns == '^') {
			continue
		}
		res = append(res, n)
	}
	return res
}

func main() {
	maze := readMaze()
	// fmt.Println(getNeigh(maze, rc{5, 3}))
	for r, row := range(maze) {
		for c, sym := range(row) {
			if sym == '#' {
				continue
			}
			if len(getNeigh(maze, rc{r, c})) > 2 {
				fmt.Printf("cross at %d:%d", r, c)
			}
		}
	}
}