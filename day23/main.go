package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/samber/lo"
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
	return maze // start == 0, 1; end == N,N-1
}

// moves are also slides values also used as steps/directions
type move byte
var moves = map[move]drc {
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
	'^': {-1, 0},
}

func getNeigh(maze []string, pos rc) (good_neigh []rc) {
	sym := maze[pos.r][pos.c]
	// mandatory slide
	if d, ok := moves[move(sym)]; ok {
		return []rc{pos.add(d)}
	}
	res := []rc{}
	for s, d := range(moves) {
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
	// fmt.Println(getNeigh(maze, rc{13, 13}))
	for r, row := range(maze) {
		for c, sym := range(row) {
			if sym == '#' {
				continue
			}
			// all are signposted
			if len(getNeigh(maze, rc{r, c})) > 2 {
				fmt.Printf("non-trivial cross at %d:%d", r, c)
			}
			if len(lo.Filter(getNeigh(maze, rc{r, c}), func(pos rc, _ int) bool {return maze[pos.r][pos.c] != '.'})) == 2 {
				fmt.Printf("cross at %d:%d\n", r, c)
			}
		}
	}
}