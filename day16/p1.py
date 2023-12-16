#!/usr/bin/python3

import sys
import typing

maze = [[c for c in l] for l in open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()]
print("maze:\n", "\n".join([''.join(l) for l in maze]))

DNUM = "RDLU"
DYDX  = [[0, 1], [1, 0], [0, -1], [-1, 0]] # DYDX[DNUM.index(d)]

beams = [[0, 0, 0]] # y, x, dir

# while len(beams) > 0:
