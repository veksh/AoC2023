#!/usr/bin/python3

import sys

maze = [[c for c in l] for l in open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()]

DYDX = {'R': [0, 1], 'D': [1, 0], 'L': [0, -1], 'U': [-1, 0]}
REFL = {
  "\\": {'R': 'D', 'L': 'U', 'D': 'R', 'U': 'L'},
  "/" : {'R': 'U', 'L': 'D', 'D': 'L', 'U': 'R'}
}

res = 0
starts = [(r, -1,           'R') for r in range(0, len(maze))]    + \
         [(r, len(maze[0]), 'L') for r in range(0, len(maze))]    + \
         [(-1, c,           'D') for c in range(0, len(maze[0]))] + \
         [(len(maze), c,    'U') for c in range(0, len(maze[0]))]
for beam in starts:
  energized = set()
  beams = [beam]
  while len(beams) > 0:
    newbeams = []
    for r, c, d in beams:
      r, c = r + DYDX[d][0], c + DYDX[d][1]
      if r < 0 or r >= len(maze) or c < 0 or c >= len(maze[0]):
        continue
      if tuple([r, c, d]) in energized:
        continue
      energized.add(tuple([r, c, d]))
      sym = maze[r][c]
      if sym in '/\\':
        newbeams.append((r, c, REFL[sym][d]))
      elif sym == '|' and d in 'RL':
        newbeams.append((r, c, 'U'))
        newbeams.append((r, c, 'D'))
      elif sym == '-' and d in 'UD':
        newbeams.append((r, c, 'L'))
        newbeams.append((r, c, 'R'))
      else:
        newbeams.append((r, c, d))
    beams = newbeams
  newres = len(set((p[0], p[1]) for p in energized))
  res = max(res, newres)
print(res)
