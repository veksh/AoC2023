#!/usr/bin/python3

import sys

field = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()

ID_DIR = ((1,0),(0,1))
DIR = {
  "|": ID_DIR,
  "-": ID_DIR,
  "L": ((0,1),(1,0)),
  "J": ((0,-1),(-1,0)),
  "7": ((0,1),(1,0)),
  "F": ((0,-1),(-1,0)),
}

# t = DIR[sym]; d = (d[0]*t[0,0] + d[1]*d[0,1], d[0]*t[1,0] + d[1]*t[1,1])
print(field)

# row, col == y, x
spos = (-1, -1)
for r in range(0, len(field)):
  if field[r].find('S') >= 0:
    spos = (r, field[r].index('S'))
    break
print("S:", spos)

move = (0, 0)
if   spos[1] > 0               and field[spos[0]][spos[1]-1] in ['-', 'F', 'L']:
  move = (0, -1)
elif spos[1] < len(field[0])-1 and field[spos[0]][spos[1]+1] in ['-', 'J', '7']:
  move = (0,  1)
elif spos[0] > 0               and field[spos[0]-1][spos[1]] in ['|', 'F', '7']:
  move = (-1, 0)
elif spos[0] < len(field) - 1  and field[spos[0]+1][spos[1]] in ['|', 'J', 'L']:
  move = (1,  0)
else:
  print("cannot move")
  sys.exit(1)
print("move1:", move)