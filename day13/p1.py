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

print("\n".join(mazes[0]))
print()
print("\n".join(rotateMaze(mazes[0])))