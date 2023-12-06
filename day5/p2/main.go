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
	fmt.Println("seeds:", seeds)
	fmt.Println("seed intervals:", seedIntervals)
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
				if oStart < 0 || oLength <= 0 {
					fmt.Printf(" *** warn: pathological split")
				}
				res = append(res, []int{mStart, oLength})
				mappedParts = append(mappedParts, []int{oStart, oLength})
			}
		}
		if len(mappedParts) > 0 {
			if mappedParts[0][1] == iLength {
				fmt.Printf("  fully mapped: (%d + %d)\n", mappedParts[0][0], mappedParts[0][1])
				continue
			}
			sort.Slice(mappedParts, func(i, j int) bool {return mappedParts[i][0] < mappedParts[j][0]})
			prevEnd := iStart
			mappedParts = append(mappedParts, []int{iStart + iLength, 0})
			sumLen := 0
			for _, mp := range(mappedParts) {
				sumLen += mp[1]
				if uLen := mp[0] - prevEnd; uLen > 0 {
					fmt.Printf("  unmapped: (%d + %d)\n", prevEnd, uLen)
				  if prevEnd < 0 || uLen <= 0 {
					  fmt.Printf(" *** warn: pathological unmapped")
				  }
					res = append(res, []int{prevEnd, uLen})
					sumLen += uLen
				}
				prevEnd = mp[0] + mp[1]
			}
			if sumLen != iLength {
				fmt.Printf(" *** warn: %d unaccounted\n", sumLen - iLength)
			}
		} else {
			fmt.Printf(" interval (%d + %d) is fully unmapped\n", iStart, iLength)
			res = append(res, []int{iStart, iLength})
		}
	}
	return res
}

func intervalSum(intervals [][]int) (res int) {
	for _, interval := range(intervals) {
		res += interval[1]
	}
	return res
}

func main() {
	intervals, maps := readData(FILE_NAME)
	fmt.Println("range maps:", maps)
	fmt.Println("seed intervals:", intervals)
	iss := intervalSum(intervals)
	for _, m := range(maps) {
		intervals = mapIntervals(intervals, m)
		isn := intervalSum(intervals)
		if isn != iss {
			fmt.Println("sum mismatch:", isn, "vs", iss)
			os.Exit(-1)
		}
		// check overlaps
		sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
		fmt.Println("new intervals:", intervals)
		prev := intervals[0][0] - 1
		for i,in := range(intervals) {
			if in[0] <= prev {
				fmt.Println("*** warn: interval", i, "overlaps with prev")
				os.Exit(-1)
			}
			prev = in[0] + in[1] - 1
		}
	}
	fmt.Println("final intervals:", intervals)
	fmt.Println("answer is", intervals[0][0]) // 183085156 is still wrong :)
}