package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "github.com/samber/lo"
)

type gate struct {
	name string
	kind byte
	inputs  []string
	outputs []string
}

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {log.Fatalf("cannot open: %v", err)}
	defer fh.Close()
	lr := bufio.NewReader(fh)
	for {
		str, err := lr.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(strings.Trim(strings.Replace(str, ",", "", -1), "\n"), " ")
		kind, name, outs := string(parts[0][0]), parts[0][1:], parts[2:]
		fmt.Printf("%s: %s, outs %v\n", name, kind, outs)
	}
}