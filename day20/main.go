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
		fmt.Println("got", strings.Trim(str, "\n"))
	}
}