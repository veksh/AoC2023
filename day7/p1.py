#!/usr/bin/python3

import sys

# "0" for stdin
lines = open(sys.argv[1] if len(sys.argv) > 1 else "input_test.txt").read().splitlines()
bets = {k: int(v) for k,v in [p.split() for p in lines]}
cards = {l: i for i, l in enumerate(reversed("A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2".split(", ")))}
print(bets)
print(cards)