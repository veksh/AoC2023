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

def findMirr(maze: list[str]) -> int:
  for lineno in range(0, len(maze)-1):
    if all(maze[lineno-offset] == maze[lineno+offset+1] for offset in range(0, min(lineno+1, len(maze)-lineno-1))):
      return lineno
  return -1

# print("\n".join(mazes[0]))
# print(findMirr(mazes[0]))
# print()
# print("\n".join(rotateMaze(mazes[0])))
# print(findMirr(rotateMaze(mazes[0])))

ans1 = sum(100*(findMirr(m) + 1) + (findMirr(rotateMaze(m))+1) for m in mazes)
print("ans1:", ans1)