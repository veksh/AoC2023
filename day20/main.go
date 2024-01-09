package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "github.com/samber/lo"
)

type Gate struct {
	name string
	kind string
	state map[string]int
	outputs []string
}

func (g Gate)String() string {
	return fmt.Sprintf("%s: %s, outs %v, states %v",
		g.name, g.kind, g.outputs, g.state)
}

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {log.Fatalf("cannot open: %v", err)}
	defer fh.Close()
	lr := bufio.NewReader(fh)
	gates := map[string]Gate{}
	for {
		str, err := lr.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(strings.Trim(strings.Replace(str, ",", "", -1), "\n"), " ")
		kind, name, outs := string(parts[0][0]), parts[0][1:], parts[2:]
		gate := Gate{name: name, kind: kind, outputs: outs, state: map[string]int{}}
		fmt.Println(gate)
		gates[name] = gate
	}
	for gname, gate := range(gates) {
		for _, oname := range(gate.outputs) {
			if outg, ok := gates[oname]; ok {
				outg.state[gname] = -1
			} else {
				gates[oname] = Gate{name: oname, kind: "s", state: map[string]int{}}
			}
		}
	}
	fmt.Println(gates)
}