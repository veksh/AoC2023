#!/usr/bin/python3

import sys

# "0" for stdin
lines = open(sys.argv[1] if len(sys.argv) > 1 else "input_test.txt").read().splitlines()
bets = {k: int(v) for k,v in [p.split() for p in lines]}
print(bets)