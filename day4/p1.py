#/usr/bin/python3

import re

# Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
FILE_NAME = "input_test.txt"

for line in open(FILE_NAME).read().splitlines():
  cardNo = int(line[5:line.index(':')])
  winning, ours = [set(map(int, n.split())) for n in line[line.index(':')+1:].split('|')]
  print("line %3d: winning %s ours %s" % (cardNo, winning, ours))
