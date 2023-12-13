#!/usr/bin/python3

import sys

mazes = [p.splitlines() for p in open(sys.argv[1] if len(sys.argv) > 1 else 0).read().split("\n\n")]
print(mazes)

