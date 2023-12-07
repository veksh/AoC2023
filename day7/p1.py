#!/usr/bin/python3

import sys
from collections import Counter

# "0" for stdin
lines = open(sys.argv[1] if len(sys.argv) > 1 else "input_test.txt").read().splitlines()
bets = {k: int(v) for k,v in [p.split() for p in lines]}
label_val = {l: i for i, l in enumerate(reversed("A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2".split(", ")))}
# tuple(sorted(Counter([l for l in card]).values()))
combo_val = {
  (1, 1, 1, 1, 1): 0,  # high card
  (1, 1, 1, 2):    1,  # one pair
  (1, 2, 2):       2,  # two pair
  (1, 1, 3):       3,  # three of a kind
  (2, 3):          4,  # full house
  (1, 4):          5,  # four of a kind
  (5, ):           6   # five of a kind
}

card2val = lambda card: \
  chr(ord('A') + combo_val[tuple(sorted(Counter([l for l in card]).values()))]) + \
  "".join([chr(ord('A') + label_val[c]) for c in card])

print(bets)
print(label_val)
print(combo_val)
res = 0
for i, c in enumerate(sorted(bets.keys(), key=card2val)):
  res += (i+1)*bets[c]
print("answer:", res)