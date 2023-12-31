#/usr/bin/python3

# Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
FILE_NAME = "input.txt"

res = 0
for line in open(FILE_NAME).read().splitlines():
  cardNo = int(line[5:line.index(':')])
  winning, ours = [set(map(int, n.split())) for n in line[line.index(':')+1:].split('|')]
  print("line %3d: winning %s ours %s" % (cardNo, winning, ours))
  wins = len(winning & ours)
  if wins > 0:
    res += 2**(wins-1)
print("result:", res)
