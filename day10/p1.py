#!/usr/bin/python3

import sys

lines = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()
field = [l.split() for l in lines]

print(field)