package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "../input.txt"

func strs2ints(s []string) []int {
	res := []int{}
	for _, ss := range(s) {
		n, _ := strconv.Atoi(ss)
		res = append(res, n)
	}
	return res
}

func readData(fname string) (seeds []int, maps [][][]int) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seeds = strs2ints(strings.Fields(scanner.Text())[1:])
	fmt.Println("seeds:", seeds)
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
		fmt.Println("  ranges:", rangeMap)
		maps = append(maps, rangeMap)
	}
	return seeds, maps
}

func mapOne(n int, ranges [][]int) int {
	for _, r := range(ranges) {
		if n >= r[1] && n < r[1] + r[2] {
			return n - r[1] + r[0]
		}
	}
	return n
}

func main() {
	seeds, maps := readData(FILE_NAME)
	res := -1
	for _, s := range(seeds) {
		f := s
		for _, m := range(maps) {
			f = mapOne(f, m)
		}
		fmt.Printf("%d => %d\n", s, f)
		if res == -1 || f < res {
			res = f
		}
	}
	fmt.Println("ans:", res)
}