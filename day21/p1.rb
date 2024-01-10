#!/usr/bin/ruby

require "set"

steps = ARGV.pop().to_i || 6
maze = ARGF.readlines().map {|l| l.strip().chars()}

sr = maze.find_index {|r| r.include? "S"}
sc = maze[sr].find_index("S")
puts "start: #{sr} #{sc}, moves: #{steps}"

moves = [[1, 0], [-1, 0], [0, 1], [0, -1]]

seen = [Set.new(), Set.new()] # even and odd

q = [[sr, sc]]
(1..steps).each do |i|
  qnew = []
  curr_seen = seen[i % 2]
  q.each do |p|
    moves.each do |move|
      n = [p[0] + move[0], p[1] + move[1]]
      if !curr_seen.include?(n) && maze[n[0]][n[1]] == "."
        curr_seen.add(n)
        qnew.push(n)
      end
    end
  end
  q = qnew
end
# puts "seen: #{seen}, sizes: #{seen.map(&:length)}"
puts "ans1: #{seen[0].length() + 1}"