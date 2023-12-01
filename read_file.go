package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"regexp"
)

// format: <name>: <num1> <num2> <num3>
const FILE_NAME string = "input.txt"

type outRec struct {
	name string
	nums []int
}

func readWithScanner(inputFileName string) []outRec {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("problem opening input: %s", err)
	}
	defer f.Close()
	res := []outRec{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		log.Print("got '", line, "'")

		// simple split
		flds := strings.Fields(line)
		rec := outRec{name: strings.Trim(flds[0], ":"), nums: make([]int, len(flds) - 1)}
		for i,ns := range(flds[1:]) {
			if n,err := strconv.Atoi(ns); err == nil {
				rec.nums[i] = n
			}
		}

		// alt: pos + regexp
		name, rest, _ := strings.Cut(line, ":")
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllString(rest, -1)
		rec = outRec{name: name, nums: make([]int, len(nums))}
		for i,ns := range(nums) {
			if n,err := strconv.Atoi(ns); err == nil {
				rec.nums[i] = n
			}
		}

		res = append(res, rec)
	}
	return res
}

func readWithScanf(inputFileName string) []outRec {
	// alt for f : os.Stdin
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("problem opening input: %s", err)
	}
	defer f.Close()
	res := []outRec{}
	for {
		// if record format is constant known in advance and i
		// var name string
		// var n1, n2, n3 int
		// if _, err := fmt.Fscanln(f, &name, &n1, &n2, &n3); err == nil {
		// 	res = append(res, outRec{name: strings.Trim(name, ":"), nums: []int{n1, n2, n3}})
		// } else {
		// 	break
		// }

		// if count is given: fmt.Fscanln(in, &testCount)

		// bad if number of args is not known: 1st symbol of next line is eaten
		var name string
		_, err := fmt.Fscan(f, &name)
		if err != nil {
			break
		}
		rec := outRec{name: strings.TrimSuffix(name, ":"), nums: []int{}}
		var num int
		for {
			if n, err := fmt.Fscan(f, &num); n == 1 && err == nil {
				rec.nums = append(rec.nums, num)
			} else {
				log.Print(" got error: ", err)
				// skipping newlines somehow: "expected integer" and eating 1st symbol of next line
				break
			}
		}
		res = append(res, rec)
	}
	return res
}

func main() {
	fmt.Printf("scanner: %v\n", readWithScanner(FILE_NAME))
	fmt.Printf("scanf:   %v\n", readWithScanf(FILE_NAME))
}
