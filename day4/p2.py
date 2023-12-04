#/usr/bin/python3

# Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
FILE_NAME = "input.txt"

qty = []
for line in open(FILE_NAME).read().splitlines():
  cardNo = int(line[5:line.index(':')])
  winning, ours = [set(map(int, n.split())) for n in line[line.index(':')+1:].split('|')]
  wins = len(winning & ours)
  print("card %3d: winning %s ours %s wins %d" % (cardNo, winning, ours, wins))
  if len(qty) < cardNo + wins:
    qty += [1] * (cardNo + wins - len(qty))
  for i in range(0, wins):
    qty[cardNo + i] += qty[cardNo-1]
  print(" qty", qty)
print("result: ", sum(qty))
