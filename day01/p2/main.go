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
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	digstr := "(one|two|three|four|five|six|seven|eight|nine|[1-9])"
	re1 := regexp.MustCompile(digstr + ".*" + digstr)
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
		log.Printf(" first %s last %s: + %d", digits[0], digits[len(digits)-1], w2n[digits[0]]*10 + w2n[digits[len(digits)-1]])
		// res += w2n[digits[0]]*10 + w2n[digits[len(digits)-1]]
		// alt, just to check
		ss := re1.FindStringSubmatch(line)
		if len(ss) != 3 {
		  // lfnt5 was 55 for the p1: first == last
			log.Printf(" *** check: only 1, ss %v", ss)
		} else {
  		first, last := ss[1], ss[2]
	  	if first != digits[0] || last != digits[len(digits)-1] {
		  	// startfrom6,must_end_in_one_:xxxonetwoneetc: 1st stumbles on "twone"
			  log.Printf(" *** check: mismatch, ss %v", ss[1:])
			  res += w2n[ss[1]]*10 + w2n[ss[2]]
			  continue
			}
		}
		res += w2n[digits[0]]*10 + w2n[digits[len(digits)-1]]
	}
	log.Printf("done, res: %d", res)  // 53551 is not ok, 53539 is it :)
}