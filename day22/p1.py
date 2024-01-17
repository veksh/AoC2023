#!/usr/bin/python3

import sys

falling = [l.split("~") for l in open(sys.argv[1]).read().rstrip("\n").split("\n")]
print(falling)