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
  print("walking from %s to %s" % (s, f))
  ans = set()
  usedEdges = set()
  for n in nodes[s]:
    visited = set([s])
    print("trying", n, ", seen edges", usedEdges)
    usedEdges.add((s, n))
    queue = deque([n])
    parent = {n: s, s: None}
    while len(queue) > 0:
      e = queue.pop()
      print(" looking at", e)
      visited.add(e)
      if e == f:
        print("  reached", f, "starting from", n, "parent", parent[e])
        ans.add(n)
        while parent[e] != s:
          print("   via", parent[e])
          usedEdges.add((parent[e], e))
          e = parent[e]
        break
      for tip in nodes[e]:
        if (e, tip) in usedEdges:
          continue
        if tip in parent:
          continue
        parent[tip] = e
        print("  adding to queue:", tip)
        queue.appendleft(tip)
  return ans

print("starts:", findPathEnds(nodes, 'cmg', 'bvb'))