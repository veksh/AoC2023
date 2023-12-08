package main

import (
	"bufio"
	"fmt"
	"os"
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

func readData(fh *os.File) (turns string, graph map[string][2]string) {
  scanner := bufio.NewScanner(fh)
  scanner.Scan()
  turns = scanner.Text()
  scanner.Scan()
  graph = map[string][2]string{}
  for scanner.Scan() {
  	// AAA = (BBB, CCC)
  	// 0123456789012345
  	line := scanner.Text()
  	graph[line[0:3]] = [2]string{line[7:10], line[12:15]}
  }
	return turns, graph
}

func solve1(turns string, graph map[string][2]string) (res int) {
	currNode, currTurnNo := "AAA", 0
	for currNode != "ZZZ" {
		res += 1
		if turns[currTurnNo] == 'L' {
			currNode = graph[currNode][0]
		} else {
			currNode = graph[currNode][1]
		}
		currTurnNo = (currTurnNo + 1) % len(turns)
		fmt.Printf("round %d: node %s (%v), next turn %s\n", res, currNode, graph[currNode], string(turns[currTurnNo]))
	}
	return res
}

func solve2(turns string, graph map[string][2]string) (res int) {
	currNodes, currTurnNo := []string{}, 0
	for k := range(graph) {
		if k[2] == 'A' {
			currNodes = append(currNodes, k)
		}
	}
	fmt.Println("startNodes:", currNodes)
	lent := len(turns)
	for {
		res += 1
		allZ := true
		currTurn := 0
		if turns[currTurnNo] == 'R' {
			currTurn = 1
		}
		for i, currNode := range(currNodes) {
			currNodes[i] = graph[currNode][currTurn]
			if currNodes[i][2] != 'Z' {
				allZ = false
			}
		}
		currTurnNo = (currTurnNo + 1) % lent
		if allZ {
			break
		}
		// if res % 1_000_000 == 0 {
			fmt.Println("iteration", res, "nodes", currNodes)
		// }
	}
	return res
}


func main() {
	turns, graph := readData(getFH("input.txt"))
	fmt.Println(turns)
	fmt.Println(graph)
	// fmt.Println("part1:", solve1(turns, graph))
	fmt.Println("part2:", solve2(turns, graph))
}