package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const INPUT = "../input.txt"

func main() {
	log.Print("start")
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	// re := regexp.MustCompile(`\d`)
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	w2n := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	scanner := bufio.NewScanner(f)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("got %s", line)
		digits := re.FindAllString(line, -1)
		// alt with re1 := regexp.MustCompile(`(\d).*(\d)`)
		// ss := re1.FindStringSubmatch("p1a2ra33no4a")
		// first, last := ss[1], ss[2]
		log.Printf(" first %s last %s", digits[0], digits[len(digits)-1])
		res += w2n[digits[0]]*10 + w2n[digits[len(digits)-1]]
	}
	log.Printf("done, res: %d", res)  // 55017 is ok
}