#!/usr/bin/python3

import sys
import typing

mazes = [p.splitlines() for p in open(sys.argv[1] if len(sys.argv) > 1 else 0).read().split("\n\n")]

def rotateMaze(maze: list[str]) -> list[str]:
  rot = ['' for i in range(0, len(maze[0]))]
  for lnum,line in enumerate(maze):
    for cnum,char in enumerate(line):
      rot[cnum] += char
  return rot

def lineDiff(l1, l2: str) -> int:
  return sum(0 if l1[i] == l2[i] else 1 for i in range(0, len(l1)))

def findSmudge(maze: list[str]) -> int:
  for lineno in range(0, len(maze)-1):
    if sum(lineDiff(maze[lineno-offset], maze[lineno+offset+1]) for offset in range(0, min(lineno+1, len(maze)-lineno-1))) == 1:
      return lineno
  return -1

ans2 = sum(100*(findSmudge(m) + 1) + (findSmudge(rotateMaze(m))+1) for m in mazes)
print("ans2:", ans2)