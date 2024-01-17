#!/usr/bin/python3

import sys
import re

class Brick:

  def __init__(self, coords):
    self.x = [coords[0], coords[3]]
    self.y = [coords[1], coords[4]]
    self.z = [coords[2], coords[5]]

  def __str__(self):
    return "x %s y %s z %s" % (self.x, self.y, self.z)

  def is_overlap_xy(self, other):
    return ((self.x[0] <= other.x[1] and self.x[1] >= other.x[0]) and
            (self.y[0] <= other.y[1] and self.y[1] >= other.y[0]))

lines = [list(map(int, re.split("[,~]", l))) for l in open(sys.argv[1]).read().rstrip("\n").split("\n")]
# [[x1, x2], [y1, y2], [z1,z2], ...]; x1 <= x2 etc; inclusive,
# x and y from 0, z from 1 (box[2][0] == 1 is on the ground)
# lowest first
falling = [Brick(l) for l in lines]
falling.sort(key = lambda b: b.z[0])

landed = []
supports = set()
for brick in falling:
  print("falling: %s" % brick)
  if brick.z[0] != 1:
    newz = 1
    supp = -1
    for i, brick_l in enumerate(landed):
      if brick_l.z[1] > brick.z[0]:
        break
      if brick.is_overlap_xy(brick_l):
        print("  overlaps with %s" % brick_l)
        if brick_l.z[1] + 1 > newz:
          newz = brick_l.z[1] + 1
          print("   new support %d" % i)
          supp = i
        else:
          if brick_l.z[1] + 1 == newz:
            print("   dup support")
            supp = -1
    brick.z = [newz, brick.z[1] - (brick.z[0] - newz)]
    if supp >= 0:
      print("  supported just by %d" % supp)
      supports.add(supp)
  print(" landed: %s" % brick)
  landed.append(brick)
landed.sort(key = lambda b: b.z[0])
print("supports: %s" % supports)