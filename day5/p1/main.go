package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seeds := strs2ints(strings.Fields(scanner.Text())[1:])
	fmt.Println("seeds:", seeds)
	scanner.Scan()
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println("line:", txt)
	}
}