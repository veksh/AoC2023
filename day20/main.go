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
	kind string
	states  map[string]int
	outputs []string
}

func (g gate)String() string {
	return fmt.Sprintf("%s: %s, outs %v, states %v",
		g.name, g.kind, g.outputs, g.states)
}

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {log.Fatalf("cannot open: %v", err)}
	defer fh.Close()
	lr := bufio.NewReader(fh)
	gates := []gate{}
	for {
		str, err := lr.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(strings.Trim(strings.Replace(str, ",", "", -1), "\n"), " ")
		kind, name, outs := string(parts[0][0]), parts[0][1:], parts[2:]
		gate := gate{name: name, kind: kind, outputs: outs, states: map[string]int{}}
		fmt.Println(gate)
		gates = append(gates, gate)
	}
}