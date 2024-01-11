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

## init with dummy [step, tile_r, tile_c] to see first match
cnt = (1..maze.length).map {(1..maze[0].length).map {[0, 0, 0]}}

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
        cnt[n[0] % rows][n[1] % cols].push([i, n[0] / rows, n[1] / cols])
      end
    end
  end
  q = qnew
end
# took 4GB and 2:45 to reach 5K steps on test input :)
puts "ans2: #{seen[0].length()}"
cnt.each_with_index do |r, ri|
  r.each_with_index do |c, ci|
    # +1 event: step, grid r, grid c
    plus1s = cnt[ri][ci].each_cons(2).select {|p| p[0][0] != p[1][0]}
    printf("%2d:%2d: ", ri, ci)
    print(plus1s.map {|p| sprintf("s %2d (+%2d) %3d:%3d",
      p[1][0], p[1][0] - p[0][0], p[1][1], p[1][2])}.join("|"))
    printf("\n")
  end
end