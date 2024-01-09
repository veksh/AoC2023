package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	// "github.com/samber/lo"
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

const NUM_PRESSES = 1000

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
		}
		return -1
	case "b":
		return val
	default:
		return 0
	}
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
	cnt := map[int]int{-1: 0, 1: 0}
	for i := 0; i < NUM_PRESSES; i++ {
		fmt.Printf("*** run %d\n", i)
		queue := []Signal{{"roadcaster", -1}}
		cnt[-1] += 1
		for len(queue) > 0 {
			spew.Printf("q: %v\n", queue)
			newq := []Signal{}
		 	for _, signal := range(queue) {
		 		gate := gates[signal.gateName]
		 		fmt.Printf(" processing signal %d from %s to %v\n", signal.value, signal.gateName, gate.outputs)
		 		for _, oname := range(gate.outputs) {
		 			cnt[signal.value] += 1
		 			res := gates[oname].Process(signal.gateName, signal.value)
					if res != 0 {
						fmt.Printf("  %s outputs %d\n", oname, res)
						newq = append(newq, Signal{oname, res})
					}
				}
			}
			fmt.Printf(" end of queue: cnt %v\n", cnt)
			queue = newq
		}
	}
	fmt.Printf("highs %d, lows %d, ans1 %d\n", cnt[1], cnt[-1], cnt[1]*cnt[-1])

	res2 := 0
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
	for i := 0; i < 1_000_000_000; i++ {
		if i % 1_000_000 == 0 {
			fmt.Printf("*** run %d\n", i)
		}
		queue := []Signal{{"roadcaster", -1}}
		cnt[-1] += 1
		for len(queue) > 0 {
			newq := []Signal{}
		 	for _, signal := range(queue) {
		 		gate := gates[signal.gateName]
		 		for _, oname := range(gate.outputs) {
		 			if signal.value == -1 && oname == "rx" {
		 				fmt.Printf("*** found answer at step %d\n", i)
		 				res2 = i + 1
		 			}
		 			cnt[signal.value] += 1
		 			res := gates[oname].Process(signal.gateName, signal.value)
					if res != 0 {
						newq = append(newq, Signal{oname, res})
					}
				}
			}
			queue = newq
			if res2 != 0 {
				queue = []Signal{}
			}
		}
	}
	fmt.Printf("res2: %d\n", res2)
}