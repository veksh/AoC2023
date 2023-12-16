#!/usr/bin/python3

import sys
import typing

maze = open(sys.argv[1] if len(sys.argv) > 1 else 0).read().splitlines()
print("maze:\n", "\n".join(maze))