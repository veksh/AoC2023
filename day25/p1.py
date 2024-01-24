#!/usr/bin/python3

import sys
from collections import defaultdict, deque

lines = [l.replace(":", "").split() for l in open(sys.argv[1]).read().rstrip("\n").splitlines()]
edges = defaultdict(set)
for l in lines:
  n = l[0]
  for p in l[1:]:
    edges[n].add(p)
    edges[p].add(n)
#for n, edges in nodes.items():
#  print("%s: vals %s" % (n, edges))

# lets find N independent paths from s to f (where N <= min(num of neighbours of s or f))
# if N == 3, cutting them (by severing first edges leading from s) will split a graph
def findPathStarts(edges, s, f):
  # print("walking from %s to %s" % (s, f))
  ans = set()
  usedEdges = set()
  for n in edges[s]:
    visited = set([s])
    # print("trying", n, ", seen edges", usedEdges)
    usedEdges.add((s, n))
    queue = deque([n])
    parent = {n: s, s: None}
    while len(queue) > 0:
      e = queue.pop()
      # print(" looking at", e)
      visited.add(e)
      if e == f:
        # print("  reached", f, "starting from", n, "parent", parent[e])
        ans.add(n)
        while parent[e] != s:
          # print("   via", parent[e])
          usedEdges.add((parent[e], e))
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
  return ans

# print("*** starts for cmg->bvb:", findPathStarts(edges, 'cmg', 'bvb'))
# print("*** starts for cmg->frs:", findPathStarts(edges, 'cmg', 'frs'))

nodes = edges.keys()
for i, s in enumerate(nodes):
  for j, d in enumerate(list(nodes)[i+1:]):
    ps = findPathStarts(edges, s, d)
    if len(ps) == 3:
      break
  if len(ps) == 3:
    break
print("found: %s to %s via %s" % (s, d, ps))

def subsetSize(edges, n, taboo):
  q = deque([s])
  visited = set()
  while len(q) > 0:
    e = q.pop()
    if e in visited:
      continue
    visited.add(e)
    for n in edges[e]:
      if (n, e) in taboo or n in visited:
        continue
      q.appendleft(n)
  return len(visited)

print("total:", len(edges), "or", subsetSize(edges, s, set()))
taboo_edges = set([(s, n) for n in ps])
print("taboo:", taboo_edges)
fs = subsetSize(edges, s, taboo_edges)
fd = subsetSize(edges, d, taboo_edges)
print("from %s: %d, from %s: %d" % (s, fs, d, fd))
print("ans1:", fs*fd)