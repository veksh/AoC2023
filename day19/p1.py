#!/usr/bin/python3

import sys
import re

# class Stage:
#  def __init__(self, line):

slines, plines = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().rstrip("\n").split("\n\n")

parts = [{k: int(v) for k,v in re.findall(r'([xmas])=(\d+)[,}]', l)} for l in plines.split("\n")]
print(parts)

stage0 = re.search(r'^\w+', slines[0])[0]
stages = {}
for l in slines.split("\n"):
  name, sink = re.findall(r'^(\w+)\{.*,(\w+)\}$', l)[0]
  stages[name] = {"sink": sink}
  stages[name]["ops"] = []
  for m in re.findall(r'([xmas])([<>])(\d+):(\w+)', l):
    stages[name]["ops"].append((m[0], m[1], int(m[2]), m[3]))
print(stages)
