#!/usr/bin/python3

import sys
import re
import pprint
import copy
from functools import reduce

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
# pprint.pprint(parts)

stages = {}
for l in slines.split("\n"):
  name, sink = re.findall(r'^(\w+)\{.*,(\w+)\}$', l)[0]
  stages[name] = {"sink": sink}
  stages[name]["ops"] = []
  for m in re.findall(r'([xmas])([<>])(\d+):(\w+)', l):
    stages[name]["ops"].append((m[0], m[1], int(m[2]), m[3]))
#for n,s in stages.items():
#  print("stage %s: %s" % (n, s))

def process(part, stage):
  for op in stage["ops"]:
    cat, sign, cval, res = op
    pval = part[cat]
    if (sign == "<" and pval < cval) or (sign == ">" and pval > cval):
      return res
  return stage["sink"]

res = 0
for p in parts:
  stage = "in"
  while stage != "A" and stage != "R":
    stage = process(p, stages[stage])
  if stage == "A":
    res += sum(p.values())
print("ans1:", res)

# ranges of possible values for coords + produce new by applying codition
class Bounds:
  def __init__(self, bmap):
    self.bmap = bmap

  def __str__(self):
    return "%s" % self.bmap

  def withCond(self, cat, sign, val):
    # newb = self.bmap.copy()
    newb = copy.deepcopy(self.bmap)
    if (sign == "<"):
      newb[cat][1] = min(newb[cat][1], val - 1)
    else:
      newb[cat][0] = max(newb[cat][0], val + 1)
    return Bounds(newb)

  def withNegCond(self, cat, sign, val):
    newb = copy.deepcopy(self.bmap)
    if (sign == "<"):
      newb[cat][0] = max(newb[cat][0], val)
    else:
      newb[cat][1] = min(newb[cat][1], val)
    return Bounds(newb)

  def area(self):
    rs = [p[1]-p[0]+1 for p in self.bmap.values()]
    return reduce(lambda s, x: s*x, rs)

# step.res == "A":   rec(bounds + !new cond, step+1) + rsum(bounds + new cond)
# step.res == "R":   rec(bounds + !new cond, step+1)
# step.res == stage: rec(bounds + !new cond, step+1) + rec(bounds + new cond, stage)
# step == sink:
# - sink == "A": rsum(bounds)
# - sink == "R": 0
# - sink == stage: rec(bounds, stage)
def vars(stageName, step, bounds):
  stage = stages[stageName]
  if step == len(stage["ops"]):
    sink = stage["sink"]
    if sink == "A":
      return bounds.area()
    elif sink == "R":
      return 0
    else:
      return vars(sink, 0, bounds)
  cat, sign, cval, res = stage["ops"][step]
  rest = vars(stageName, step + 1, bounds.withNegCond(cat, sign, cval))
  if res == "R":
    return rest
  elif res == "A":
    return rest + bounds.withCond(cat, sign, cval).area()
  else:
    return rest + vars(res, 0, bounds.withCond(cat, sign, cval))

b = Bounds({m: [1, 4000] for m in "xmas"})
res2 = vars("in", 0, b)
print("ans2:", res2)