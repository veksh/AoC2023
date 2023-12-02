#!/usr/bin/python3

import re
from collections import defaultdict
from functools import reduce

FILE_NAME = "input.txt"

res = 0
for line in open(FILE_NAME).read().splitlines():
  print("got", line)
  id = int(re.match('Game (\d+):', line).group(1))
  print("  game id: %d" % id)
  stillOK = True
  mc = defaultdict(int)
  for game_set in line[line.index(": ")+2:].split('; '):
    print("    game set: %s" % game_set)
    for pair in game_set.split(", "):
      qty, color = pair.split(" ")
      qty = int(qty)
      print("      pair: %s -> %d" %(color, qty))
      if mc[color] < qty:
        mc[color] = qty
  res += reduce(lambda s, x: s*x, mc.values())

print("result: %d" % res)