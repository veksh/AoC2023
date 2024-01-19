package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "github.com/samber/lo"
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

// moves are also slides values also used as steps/directions
type move byte
var moves = map[move]drc {
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
	'^': {-1, 0},
}

func getNeigh(maze []string, pos rc, prevPos rc) (good_neigh []rc) {
	sym := maze[pos.r][pos.c]
	// mandatory slide
	if d, ok := moves[move(sym)]; ok {
		return []rc{pos.add(d)}
	}
	res := []rc{}
	for s, d := range(moves) {
		n := pos.add(d)
		if n == prevPos {
			continue
		}
		// actually it is walled so no need except for finish line
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

type pathVector struct {
	prev rc
	next rc
}

type void struct{}

// convert maze to adjacency graph: vertex -> dst -> length
// no dead ends?
func buildGraph(maze []string) map[rc](map[rc]int) {
	start, first, end := rc{0, 1}, rc{1, 1}, rc{len(maze)-1, len(maze[0]) - 2}
	fmt.Println("start at", start, "end at", end)
	edges := map[rc]map[rc]int{}
	seen := map[rc]void{}
	q := []pathVector{{start, first}}
	for len(q) > 0 {
		curr, nexts := q[0].prev, []rc{q[0].next}
		q = q[1:]
		edge_start, edge_len := curr, 0
		if _, ok := edges[edge_start]; !ok {
			edges[edge_start] = map[rc]int{}
		}
		for len(nexts) == 1 {
			curr, nexts = nexts[0], getNeigh(maze, nexts[0], curr)
			edge_len += 1
		}
		if len(nexts) == 0 {
			// hope we've reached the end
			if curr != end {
				fmt.Println("*** cul-de-sac at", curr)
				continue
			}
			edges[edge_start][end] = edge_len
		} else {
			// crossroads
			fmt.Println("crossroads at", curr, "nexts", nexts)
			edges[edge_start][curr] = edge_len
			if _, ok := seen[curr]; ok {
				fmt.Println(" already seen")
				continue
			}
			seen[curr] = void{}
			for _, n := range(nexts) {
				q = append(q, pathVector{curr, n})
			}
		}
	}
	return edges
}

func main() {
	maze := readMaze()
	// fmt.Println(getNeigh(maze, rc{13, 13}))
	// for r, row := range(maze) {
	// 	for c, sym := range(row) {
	// 		if sym == '#' {
	// 			continue
	// 		}
	// 		// all are signposted
	// 		if len(getNeigh(maze, rc{r, c}, rc{0, 0})) > 2 {
	// 			fmt.Printf("non-trivial cross at %d:%d", r, c)
	// 		}
	// 		if len(lo.Filter(getNeigh(maze, rc{r, c}, rc{0, 0}), func(pos rc, _ int) bool {return maze[pos.r][pos.c] != '.'})) == 2 {
	// 			fmt.Printf("cross at %d:%d\n", r, c)
	// 		}
	// 	}
	// }
	g := buildGraph(maze)
	fmt.Println("edges:")
	for src, edges := range(g) {
		fmt.Printf("from %v\n", src)
		for dst, len := range(edges) {
			fmt.Printf(" to %v len %d\n", dst, len)
		}
	}
}