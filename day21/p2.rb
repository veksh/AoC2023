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

seen = [Set[[sr, sc]]] # enough to keep only last 2
max_seen = 2
ans2 = 1
(1..steps).each do |i|
  puts "step #{i}, sizes: #{seen.map(&:length)}" if i % 100 == 0
  curr_seen = Set.new()
  seen.unshift(curr_seen)
  seen[1].each do |p|
    moves.each do |move|
      n = [p[0] + move[0], p[1] + move[1]]
      if !seen.any? {|s| s.include?(n)} && maze[n[0] % rows][n[1] % cols] == "."
        curr_seen.add(n)
      end
    end
  end
  ans2 += curr_seen.length() if i % 2 == 0
  seen.pop() if seen.length > max_seen
end
# took 4GB and 2:45 to reach 5K steps on test input :)
#puts "ans2: #{seen[0].length()}"
puts "ans2: #{ans2}"