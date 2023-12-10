#!/usr/bin/python3

import sys

field = [[c for c in l] for l in open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()]

# row, col == y, x
spos = (-1, -1)
for r in range(0, len(field)):
  if 'S' in field[r]:
    spos = (r, field[r].index('S'))
    break
print("S:", spos)

move = [0, 0]
if   spos[1] > 0               and field[spos[0]][spos[1]-1] in ['-', 'F', 'L']:
  move = [0, -1]
elif spos[1] < len(field[0])-1 and field[spos[0]][spos[1]+1] in ['-', 'J', '7']:
  move = [0,  1]
elif spos[0] > 0               and field[spos[0]-1][spos[1]] in ['|', 'F', '7']:
  move = [-1, 0]
elif spos[0] < len(field) - 1  and field[spos[0]+1][spos[1]] in ['|', 'J', 'L']:
  move = [1,  0]
else:
  print("cannot move")
  sys.exit(1)
print("move1:", move)

pos = [spos[0] + move[0], spos[1] + move[1]]
cnt = 1
while field[pos[0]][pos[1]] != "S":
  sym = field[pos[0]][pos[1]]
  field[pos[0]][pos[1]] = "@"
  if sym == "L" or sym == "7":
    move[0], move[1] = move[1], move[0]
  if sym == "J" or sym == "F":
    move[0], move[1] = -1*move[1], -1*move[0]
  pos = (pos[0] + move[0], pos[1] + move[1])
  cnt += 1
  # print("pos", pos, "move", move)
field[spos[0]][spos[1]] = "@"
print("count:", cnt, "answer 1:", cnt // 2)
print()
for r in range(0, len(field)):
  print("".join(field[r]))