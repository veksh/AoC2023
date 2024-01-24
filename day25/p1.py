#!/usr/bin/python3

import sys
from collections import defaultdict

lines = [l.replace(":", "").split() for l in open(sys.argv[1]).read().rstrip("\n").splitlines()]
print(lines)
nodes = defaultdict(set)
for l in lines:
  n = l[0]
  for p in l[1:]:
    nodes[n].add(p)
    nodes[p].add(n)
for n, edges in nodes.items():
  print("%s: vals %s" % (n, edges))