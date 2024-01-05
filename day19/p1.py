#!/usr/bin/python3

import sys
import re

slines, plines = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().rstrip("\n").split("\n\n")

#  re.findall(r'([xmas])([<>])(\d+):(\w+)', s1)
#  re.findall(r'^(\w+)\{.*,(\w+)\}$', s1)
#  re.findall(r'([xmas])=(\d+)[,}]', s2)

parts = [{k: int(v) for k,v in re.findall(r'([xmas])=(\d+)[,}]', l)} for l in plines.split("\n")]
print(parts)
stages = [re.findall(r'^(\w+)\{.*,(\w+)\}$', l) + re.findall(r'([xmas])([<>])(\d+):(\w+)', l) for l in slines.split("\n")]
print(stages)