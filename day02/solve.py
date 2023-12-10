#!/usr/bin/python3

import re

FILE_NAME = "input.txt"

LIMITS = {"red": 12, "green": 13, "blue": 14}

res = 0
for line in open(FILE_NAME).read().splitlines():
  print("got", line)
  id = int(re.match('Game (\d+):', line).group(1))
  print("  game id: %d" % id)
  stillOK = True
  for game_set in line[line.index(": ")+2:].split('; '):
    print("    game set: %s" % game_set)
    for pair in game_set.split(", "):
      qty, color = pair.split(" ")
      print("      pair: %s -> %d" %(color, int(qty)))
      if int(qty) > LIMITS[color]:
        print("        over limit!")
        stillOK = False
        break
    if not stillOK:
      break
  if stillOK:
    res += id
print("result: %d" % res)