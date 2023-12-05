package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func mapIntervals(intervals [][]int, ranges[][]int) [][]int {
	res := [][]int{}
	fmt.Println("mapping src", intervals, "via", ranges)
	for _, se := range(intervals) {
		iStart, iLength := se[0], se[1]
		mappedParts := [][]int{}
		for _, rng := range(ranges) {
			rDest, rStart, rLength := rng[0], rng[1], rng[2]
			if iStart <= rStart + rLength && iStart + iLength >= rStart {
				oStart := max(iStart, rStart)
				oLength := min(iStart + iLength, rStart + rLength) - oStart
				mStart := oStart - rStart + rDest
				fmt.Printf(" hit for (%d + %d) map (%d + %d): overlap (%d + %d) -> (%d + %d)\n",
					iStart, iLength, rStart, rLength, oStart, oLength, mStart, oLength)
				res = append(res, []int{mStart, oLength})
				mappedParts = append(mappedParts, []int{oStart, oLength})
			}
		}
		if len(mappedParts) > 0 {
			sort.Slice(mappedParts, func(i, j int) bool {return mappedParts[i][0] < mappedParts[j][0]})
			prevEnd := iStart
			mappedParts = append(mappedParts, []int{iStart + iLength - 1, 1})
			for _, mp := range(mappedParts) {
				if mp[0] > prevEnd {
					fmt.Printf(" unmapped: (%d + %d)\n", prevEnd, mp[0] - prevEnd)
					res = append(res, []int{prevEnd, mp[0] - prevEnd})
				}
				prevEnd = mp[0] + mp[1]
			}
		} else {
			fmt.Printf(" interval (%d + %d) is fully unmapped\n", iStart, iLength)
			res = append(res, []int{iStart, iLength})
		}
	}
	return res
}

func main() {
	intervals, maps := readData(FILE_NAME)
	fmt.Println("range maps:", maps)
	fmt.Println("seed intervals:", intervals)
	for _, m := range(maps) {
		intervals = mapIntervals(intervals, m)
		fmt.Println(intervals)
	}
	res := intervals[0][0]
	for _,ip := range(intervals) {
		if ip[0] < res {
			res = ip[0]
		}
	}
	fmt.Println("answer is", res)
}