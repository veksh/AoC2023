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

var part1 = true

func getNeigh(maze []string, pos rc, prevPos rc) (good_neigh []rc, is_cross bool) {
	sym := maze[pos.r][pos.c]
	// mandatory slide, cannot be crossroads
	if d, ok := moves[move(sym)]; ok {
		return []rc{pos.add(d)}, false
	}
	good_neigh = []rc{}
	not_walls := 0
	for s, d := range(moves) {
		n := pos.add(d)
		// actually it is walled so no need except for finish line
		if n.r > len(maze)-1 || n.r < 0 || n.c > len(maze[0])-1 || n.c < 0 {
			continue
		}
		ns := maze[n.r][n.c]
		if ns == '#' {
			continue
		}
		not_walls += 1
		// dont look back
		if n == prevPos {
			continue
		}
		// do not go against the slide
		if (s == '<' && ns == '>') || (s == '>' && ns == '<') || (s == '^' && ns == 'v') || (s == 'v' && ns == '^') {
			// only in part1 :)
			if part1 {
				continue
			}
		}
		good_neigh = append(good_neigh, n)
	}
	// also cross if first or last row (entry or exit)
	return good_neigh, not_walls != 2 || pos.r == 0 || pos.r == len(maze) - 1
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
		fmt.Println("looking at", q[0].next, "from", q[0].prev)
		curr, nexts := q[0].prev, []rc{q[0].next}
		q = q[1:]
		edge_start, edge_len := curr, 0
		if _, ok := edges[edge_start]; !ok {
			edges[edge_start] = map[rc]int{}
		}
		var prev rc
		cross := false
		for !cross {
			prev, curr = curr, nexts[0]
			nexts, cross = getNeigh(maze, curr, prev)
			edge_len += 1
		}
		if len(nexts) == 0 {
			// hope we've reached the end
			if curr != end {
				fmt.Println("*** cul-de-sac at", curr)
				continue
			}
			fmt.Println(" reached the end from", edge_start)
			edges[edge_start][end] = edge_len
		} else {
			// crossroads
			fmt.Println(" edge to", curr, "len", edge_len, "nexts", nexts)
			edges[edge_start][curr] = edge_len
			if _, ok := seen[curr]; ok {
				fmt.Println("  already seen")
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

func findLongest(from rc, end rc, graph map[rc](map[rc]int)) int {
	if from == end {
		return 0
	}
	res := 0
	for neigh, length := range(graph[from]) {
		res = max(res, length + findLongest(neigh, end, graph))
	}
	return res
}

func findLongestNOC(from rc, end rc, graph map[rc](map[rc]int), seen map[rc]void) int {
	if from == end {
		return 0
	}
	res := 0
	seen[from] = void{}
	for neigh, length := range(graph[from]) {
		if _, ok := seen[neigh]; ok {
			continue
		}
		res = max(res, length + findLongestNOC(neigh, end, graph, seen))
	}
	delete(seen, from)
	return res
}

func printGraph(graph map[rc](map[rc]int)) {
	fmt.Println("edges:")
	for src, edges := range(graph) {
		fmt.Printf("from %v\n", src)
		for dst, len := range(edges) {
			fmt.Printf(" to %v len %d\n", dst, len)
		}
	}
}

func ans1(maze []string) int {
	g := buildGraph(maze)
	printGraph(g)
	return findLongest(rc{0, 1}, rc{len(maze)-1, len(maze[0]) - 2}, g)
}

func main() {
	maze := readMaze()
	fmt.Println("ans1:", ans1(maze))
	// ans2 := findLongestNOC(rc{0, 1}, rc{len(maze)-1, len(maze[0]) - 2}, g, map[rc]void{})
	// fmt.Println("ans2:", ans2) // 6581 is too high :)
}