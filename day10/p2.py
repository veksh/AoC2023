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

# pos = [spos[0] + move[0], spos[1] + move[1]]
pos = [spos[0], spos[1]]
cnt = 0
m2s = {
  (1,0):  '↓',
  (0,1):  '→',
  (-1,0): '↑',
  (0,-1): '←'
}

while not (pos[0] == spos[0] and pos[1] == spos[1] and cnt > 0):
  field[pos[0]][pos[1]] = m2s[tuple(move)]
  pos = (pos[0] + move[0], pos[1] + move[1])
  sym = field[pos[0]][pos[1]]
  if sym in "L7":
    move[0], move[1] = move[1], move[0]
  if sym in "JF":
    move[0], move[1] = -1*move[1], -1*move[0]
  cnt += 1

print("count:", cnt, "answer 1:", cnt // 2)
print()
for r in range(0, len(field)):
  print("".join(field[r]))