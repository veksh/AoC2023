#!/usr/bin/python3

import sys
from collections import defaultdict, deque

# find a subset of nodes reachable from start while avoiding edges in taboo
# proper min cut is a bit harder, but our weights are 1
def reachable(edges, start, taboo):
  q = deque([start])
  visited = set()
  while len(q) > 0:
    e = q.pop()
    if e in visited:
      continue
    visited.add(e)
    for nn in edges[e]:
      if (e, nn) in taboo or (nn, e) in taboo or nn in visited:
        continue
      q.appendleft(nn)
  return visited

# lets find N independent paths from s to f (where N <= min(num of neighbours of s or f))
# if N == 3, cutting them (by severing first edges leading from s) will split a graph
def findPathStarts(edges, s, f):
  # print("walking from %s to %s" % (s, f))
  ans = []
  usedEdges = set()
  for n in edges[s]:
    visited = set([s])
    # print("trying", n, ", seen edges", usedEdges)
    queue = deque([n])
    parent = {n: s, s: None}
    while len(queue) > 0:
      e = queue.pop()
      # print(" looking at", e)
      visited.add(e)
      if e == f:
        # print("  reached", f, "starting from", n)
        ans.append([])
        while parent[e] != None:
          # print("   via", parent[e])
          usedEdges.add((parent[e], e))
          ans[-1].append((e, parent[e]))
          e = parent[e]
        break
      for tip in edges[e]:
        if (e, tip) in usedEdges:
          continue
        if tip in parent:
          continue
        parent[tip] = e
        # print("  adding to queue:", tip)
        queue.appendleft(tip)
  return ans, usedEdges

## main part
lines = [l.replace(":", "").split() for l in open(sys.argv[1]).read().rstrip("\n").splitlines()]
edges = defaultdict(set)
for l in lines:
  n = l[0]
  for p in l[1:]:
    edges[n].add(p)
    edges[p].add(n)
#for n, edges in nodes.items():
#  print("%s: vals %s" % (n, edges))

nodes = edges.keys()
for i, s in enumerate(nodes):
  for j, d in enumerate(list(nodes)[i+1:]):
    ps, usedEdges = findPathStarts(edges, s, d)
    if len(ps) == 3:
      break
  if len(ps) == 3:
    break
print("found a pair: %s to %s via %s" % (s, d, ps))

# print("total:", len(edges), "==", len(reachable(edges, s, set())))
# print("taboo:", usedEdges)
fs = len(reachable(edges, s, usedEdges))
fd = len(reachable(edges, d, usedEdges))
print("reachable from %s: %d, from %s: %d" % (s, fs, d, fd))
print("ans1:", fs*fd)

