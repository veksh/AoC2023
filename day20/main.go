package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "github.com/samber/lo"
	"github.com/davecgh/go-spew/spew"
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

// 0 none, 1 high, -1 low
func (g Gate)Signal(in string, val int) int {
	switch g.kind {
	case "%":
		if val == 1 {
			return 0
		}
		g.state["all"] *= -1
		return g.state["all"]
	case "&":
		g.state[in] = val
		for _, s := range(g.state) {
			if s == -1 {
				return 1
			}
			return -1
		}
	case "b":
		return val
	default:
		return 0
	}
	return 0
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
				if outg.kind == "&" {
					outg.state[gname] = -1
				} else {
					outg.state["all"] = -1
				}
			} else {
				gates[oname] = Gate{name: oname, kind: "s", state: map[string]int{"dummy": 0}}
			}
		}
	}
	spew.Dump(gates)
}