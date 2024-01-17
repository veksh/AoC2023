#!/usr/bin/python3

import sys
import re

lines = [list(map(int, re.split("[,~]", l))) for l in open(sys.argv[1]).read().rstrip("\n").split("\n")]
falling = []
for l in lines:
  # [[x1, x2], [y1, y2], [z1,z2], ...]; x1 <= x2 etc; inclusive,
  # x and y from 0, z from 1 (box[2][0] == 1 is on the ground)
  falling.append([[l[0], l[3]], [l[1], l[4]], [l[2], l[5]]])
# topmost first
falling.sort(key = lambda box: box[2][1])
print(falling)
