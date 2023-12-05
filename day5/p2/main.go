package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FILE_NAME = "../input_test.txt"

func strs2ints(s []string) []int {
	res := []int{}
	for _, ss := range(s) {
		n, _ := strconv.Atoi(ss)
		res = append(res, n)
	}
	return res
}

func readData(fname string) (seedIntervals [][]int, maps [][][]int) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seeds := strs2ints(strings.Fields(scanner.Text())[1:])
	seedIntervals = [][]int{}
	for i := 0; i <= len(seeds)/2; i += 2 {
		seedIntervals = append(seedIntervals, []int{seeds[i], seeds[i+1]})
	}
	sort.Slice(seedIntervals, func(i, j int) bool {return seedIntervals[i][0] < seedIntervals[j][0]})
	fmt.Println("seeds:", seeds, "intervals:", seedIntervals)
	scanner.Scan()

	maps = [][][]int{}

	for {
		if ! scanner.Scan() {
			break
		}
		txt := scanner.Text()
		fmt.Println("map name:", txt)
		rangeMap := [][]int{}
		for {
			if !scanner.Scan() {
				break
			}
			txt = scanner.Text()
			if txt == "" {
				break
			}
			rangeMap = append(rangeMap, strs2ints(strings.Fields(txt)))
		}
		sort.Slice(rangeMap, func(i, j int) bool {return rangeMap[i][0] < rangeMap[i][0]})
		fmt.Println("  ranges:", rangeMap)
		maps = append(maps, rangeMap)
	}
	return seedIntervals, maps
}

func mapOne(n int, ranges [][]int) int {
	for _, r := range(ranges) {
		if n >= r[1] && n < r[1] + r[2] {
			return n - r[1] + r[0]
		}
	}
	return n
}

// func mapInterval(start int, end int, ranges[][]int) []int {

// }

func main() {
	seedIntervals, maps := readData(FILE_NAME)
	fmt.Println("seed intervals:", seedIntervals)
	fmt.Println("range maps:", maps)
	// for _, m := range(maps) {

	// }

	// for _, s := range(seedIntervals) {
	// 	f := s
	// 	for _, m := range(maps) {
	// 		f = mapOne(f, m)
	// 	}
	// 	fmt.Printf("%d => %d\n", s, f)
	// 	if res == -1 || f < res {
	// 		res = f
	// 	}
	// }
	// fmt.Println("ans:", res)
}