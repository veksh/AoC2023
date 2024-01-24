#!/usr/bin/python3

import sys
from collections import defaultdict, deque

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


def findPathEnds(nodes, s, f):
  ans = []
  visited = set([s])
  queue = deque([s])
  while len(queue) > 0:
    e = queue.pop()
    for n in nodes[e]:
      if n in visited:
        continue
      if n == f:
        ans.append(e)
        continue
      visited.add(n)
      queue.appendleft(n)
  return ans

print("paths:", findPathEnds(nodes, 'bvb', 'cmg'))