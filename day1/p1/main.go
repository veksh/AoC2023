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
	re := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(f)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("got %s", line)
		digits := re.FindAllString(line, -1)
    log.Printf(" first %s last %s", digits[0], digits[len(digits)-1])
		// alt with re1 := regexp.MustCompile(`(\d).*(\d)`)
		// ss := re1.FindStringSubmatch("p1a2ra33no4a")
		// first, last := ss[1], ss[2]
		first, last := int(digits[0][0] - '0'), int(digits[len(digits)-1][0] - '0')
		res += first*10 + last
	}
	log.Printf("done, res: %d", res)  // 55017 is ok
}
