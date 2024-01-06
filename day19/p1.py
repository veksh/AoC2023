#!/usr/bin/python3

import sys
import re
import pprint

# class Stage:
#   def __init__(self, line):
#     self.name, self.sink = re.findall(r'^(\w+)\{.*,(\w+)\}$', l)[0]
#     stages.ops = []
#     for m in re.findall(r'([xmas])([<>])(\d+):(\w+)', l):
#       stages[name]["ops"].append((m[0], m[1], int(m[2]), m[3]))

#   def __repr__(self):
#     return "stage %s"

slines, plines = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().rstrip("\n").split("\n\n")

parts = [{k: int(v) for k,v in re.findall(r'([xmas])=(\d+)[,}]', l)} for l in plines.split("\n")]
pprint.pprint(parts)

stages = {}
for l in slines.split("\n"):
  name, sink = re.findall(r'^(\w+)\{.*,(\w+)\}$', l)[0]
  stages[name] = {"sink": sink}
  stages[name]["ops"] = []
  for m in re.findall(r'([xmas])([<>])(\d+):(\w+)', l):
    stages[name]["ops"].append((m[0], m[1], int(m[2]), m[3]))
for n,s in stages.items():
  print("stage %s: %s" % (n, s))

def process(part, ops, sink):
  for op in ops:
    cat, sign, cval, res = op
    pval = part[cat]
    if (sign == "<" and pval < cval) or (sign == ">" and pval > cval):
      return res
  return sink

print(process(parts[0], stages["qqz"]["ops"], "R"))