package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strs2ints(s []string) []int {
	res := []int{}
	for _, ss := range(s) {
		n, _ := strconv.Atoi(ss)
		res = append(res, n)
	}
	return res
}

func readData(fname string) (seedIntervals [][]int, maps [][][]int) {
	f := os.Stdin
	if fname != "stdin" {
		var err error
		f, err = os.Open(fname)
		if err != nil {
			fmt.Println("error opening file:", err)
			os.Exit(1)
		}
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

func mapIntervals2(intervals [][]int, ranges[][]int) [][]int {
	res := [][]int{}
	fmt.Println("mapping src", intervals, "via", ranges)
	queue := intervals
	for len(queue) > 0 {
		newQ := [][]int{}
		for _, se := range(queue) {
			iStart, iLength := se[0], se[1]
			hit := false
			for _, rng := range(ranges) {
				rDest, rStart, rLength := rng[0], rng[1], rng[2]
				if iStart <= rStart + rLength - 1 && iStart + iLength - 1 >= rStart {
					oStart := max(iStart, rStart)
					oLength := min(iStart + iLength, rStart + rLength) - oStart
					mStart := oStart - rStart + rDest
					fmt.Printf(" hit for (%d + %d) map (%d + %d): overlap (%d + %d) -> (%d + %d)\n",
						iStart, iLength, rStart, rLength, oStart, oLength, mStart, oLength)
					res = append(res, []int{mStart, oLength})
					if oLength < iLength {
						fmt.Printf("  partial match: %d of %d\n", oLength, iLength)
						addl := 0
						if ll := oStart - iStart; ll > 0 {
							newQ = append(newQ, []int{iStart, ll})
							fmt.Printf("  appending l (%d + %d)\n", iStart, ll)
							addl += ll
						}
						if rl := iStart + iLength - (oStart + oLength); rl > 0 {
							newQ = append(newQ, []int{oStart + oLength, rl})
							fmt.Printf("  appending r (%d + %d)\n", oStart + oLength, rl)
							addl += rl
						}
						if oLength + addl != iLength {
							fmt.Printf("  oLength mismatch: still %d vs %d\n", oLength + addl, iLength)
							os.Exit(1)
						}
					}
					hit = true
					break
				}
			}
			if !hit {
				fmt.Printf(" unmatched (%d + %d)\n", iStart, iLength)
				res = append(res, []int{iStart, iLength})
			}
		}
		queue = newQ
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
	fileName := "stdin"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	intervals, maps := readData(fileName)
	fmt.Println("range maps:", maps)
	fmt.Println("seed intervals:", intervals)
	iss := intervalSum(intervals)
	for _, m := range(maps) {
		intervals = mapIntervals2(intervals, m)
		isn := intervalSum(intervals)
		if isn != iss {
			fmt.Println("sum mismatch:", isn, "vs", iss)
			os.Exit(1)
		}
		// check overlaps
		sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
		fmt.Println("new intervals:", intervals)
		prev := intervals[0][0] - 1
		for i,in := range(intervals) {
			if in[0] <= prev {
				fmt.Println("*** warn: interval", i, "overlaps with prev")
				os.Exit(1)
			}
			prev = in[0] + in[1] - 1
		}
	}
	fmt.Println("final intervals:", intervals)
	fmt.Println("answer is", intervals[0][0]) // 183085156 is still wrong :)
}