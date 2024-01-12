#!/usr/bin/ruby

require "set"

steps = ARGV.pop().to_i || 26501365 # 5*11*481843 steps
maze = ARGF.readlines().map {|l| l.strip().chars()}

sr = maze.find_index {|r| r.include? "S"}
sc = maze[sr].find_index("S")
puts "start: #{sr} #{sc}, moves: #{steps}"
maze[sr][sc] = "."

moves = [[1, 0], [-1, 0], [0, 1], [0, -1]]
rows, cols = maze.length(), maze[0].length()

seen = [Set.new(), Set.new()] # even and odd

## init with dummy [tile_r, tile_c] to see first match
cnt = (1..maze.length).map {(1..maze[0].length).map {[]}}

q = [[sr, sc]]
(1..steps).each do |i|
  puts "step #{i}, sizes: #{seen.map(&:length)}" if i % 1000 == 0
  qnew = []
  curr_seen = seen[i % 2]
  q.each do |p|
    moves.each do |move|
      n = [p[0] + move[0], p[1] + move[1]]
      if !curr_seen.include?(n) && maze[n[0] % rows][n[1] % cols] == "."
        curr_seen.add(n)
        qnew.push(n)
        cnt[n[0] % rows][n[1] % cols].push([n[0] / rows, n[1] / cols, i])
      end
    end
  end
  q = qnew
end
# took 4GB and 2:45 to reach 5K steps on test input :)
puts "ans2: #{seen[0].length()}"
cnt.each_with_index do |r, ri|
  r.each_with_index do |c, ci|
    next if c.length() == 0
    firststep = c[c.find_index {|p| p[0] == 0 && p[1] == 0}][2]
    printf("r %2d c %2d: first %d\n", ri, ci, firststep)
    c.sort {|p1, p2| (p1[0].abs + p1[1].abs) <=> (p2[0].abs + p2[1].abs)}.each do |tr, tc, step|
      ctr = tr < 0? -1*tr + 1: tr
      ctc = tc < 0? -1*tc + 1: tc
      printf(" %2d %2d: %d (%d)\n", tr, tc, step, step - (ctr + ctc)*131)
    end
  end
end