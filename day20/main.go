package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	// "github.com/samber/lo"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
)

type Gate struct {
	name string
	kind string
	state map[string]int
	outputs []string
}

type Signal struct {
	gateName string
	value    int
}

func (g Gate) String() string {
	return fmt.Sprintf("%s: %s, outs %v, states %v",
		g.name, g.kind, g.outputs, g.state)
}

// 0 none, 1 high, -1 low
func (g Gate) Process(in string, val int) int {
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
	queue := []Signal{{"roadcaster", -1}}
	lows, highs := 0, 0
	for len(queue) > 0 {
		h := lo.Reduce(queue, func(agg int, item Signal, _ int) int {return agg + (item.value + 1)/2}, 0)
		l := len(queue) - h
		spew.Printf("q: %v, highs +%d, lows +%d\n", queue, h, l)
		highs, lows = highs + h, lows + l
		newq := []Signal{}
	 	for _, signal := range(queue) {
	 		gate := gates[signal.gateName]
	 		fmt.Printf(" processing signal %d from %s to %v\n", signal.value, signal.gateName, gate.outputs)
	 		for _, oname := range(gate.outputs) {
	 			res := gates[oname].Process(signal.gateName, signal.value)
				if res != 0 {
					fmt.Printf("  %s outputs %d\n", oname, res)
					newq = append(newq, Signal{oname, res})
				}
			}
		}
		queue = newq
	}
	fmt.Printf("higs %d, lows %d, ans1 %d\n", highs, lows, highs*lows)
}