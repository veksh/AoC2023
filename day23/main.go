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

func getNeigh(maze []string, pos rc, prevPos rc) (good_neigh []rc, is_cross bool) {
	sym := maze[pos.r][pos.c]
	// mandatory slide, cannot be crossroads
	if d, ok := moves[move(sym)]; ok {
		return []rc{pos.add(d)}, false
	}
	good_neigh = []rc{}
	not_walls := 0
	for _, d := range(moves) {
		n := pos.add(d)
		// actually it is walled so no need except for finish line
		if n.r > len(maze)-1 || n.r < 0 || n.c > len(maze[0])-1 || n.c < 0 {
			continue
		}
		// only # are impassable
		if maze[n.r][n.c] == '#' {
			continue
		}
		not_walls += 1
		// dont look back
		if n == prevPos {
			continue
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
// there are not dead ends
// in p2 slopes are just like normal paths, so vertices are bidirectional
func buildGraph(maze []string) map[rc](map[rc]int) {
	start, first, end := rc{0, 1}, rc{1, 1}, rc{len(maze)-1, len(maze[0]) - 2}
	fmt.Println("start at", start, "end at", end)
	edges := map[rc]map[rc]int{}
	seen := map[rc]void{}
	q := []pathVector{{start, first}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		fmt.Println("looking at", v.next, "from", v.prev)
		// do not go out if path was already taken
		if _, ok := seen[v.next]; ok {
			fmt.Println(" already seen it")
			continue
		}
		seen[v.next] = void{}
		curr, nexts := v.prev, []rc{v.next}
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
			fmt.Println(" edge with", curr, "len", edge_len, "nexts", nexts)
			edges[edge_start][curr] = edge_len
			// add reverse edge and do not go out from this node via this path
			seen[prev] = void{}
			if edge_start != start {
				if _, ok := edges[curr]; !ok {
					edges[curr] = map[rc]int{}
				}
				edges[curr][edge_start] = edge_len
			}
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

func findLongestAll(from rc, end rc, graph map[rc](map[rc]int), seen map[rc]void, dist int, res *int) {
	if _, ok := seen[from]; ok {
		return
	}
	if from == end {
		if dist > *res {
			*res = dist
		}
		return
	}
	seen[from] = void{}
	for neigh, length := range(graph[from]) {
		findLongestAll(neigh, end, graph, seen, dist + length, res)
	}
	delete(seen, from)
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

func ans2(maze []string) int {
	g := buildGraph(maze)
	printGraph(g)
	a2 := 0
	findLongestAll(rc{0, 1}, rc{len(maze)-1, len(maze[0]) - 2}, g, map[rc]void{}, 0, &a2)
	return a2
}

func main() {
	maze := readMaze()
	// fmt.Println("ans1:", ans1(maze))
	fmt.Println("ans2:", ans2(maze))
}