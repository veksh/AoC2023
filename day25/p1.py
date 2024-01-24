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

# lets find N independent paths from s to f (where N <= min(num of neighbours of s or f))
# if N == 3, cutting them (by severing first edges leading from s) will split a graph
def findPathEnds(nodes, s, f):
  print("walking from", s, "to", f)
  ans = set()
  preVisited = set([s])
  for n in nodes[s]:
    visited = preVisited.copy()
    print("trying", n, "excluding", visited)
    queue = deque([n])
    parent = {n: s}
    while len(queue) > 0:
      e = queue.pop()
      visited.add(e)
      if e == f:
        print(" reached", f, "starting from", n)
        ans.add(n)
        while e in parent:
          print("  via", e)
          e = parent[e]
          preVisited.add(e)
        break
      for tip in nodes[e]:
        if tip in visited:
          continue
        parent[tip] = e
        queue.appendleft(tip)
  return ans

print("starts:", findPathEnds(nodes, 'cmg', 'bvb'))