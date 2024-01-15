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
puts "maze: #{rows} rows x #{cols} cols (#{rows * cols} total)"

seen = [Set[[sr, sc]]] # enough to keep only last 2
max_seen = 2
ans2 = [1, 0]
# for all steps
cnt = [0]
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
  cnt.push(curr_seen.length())
  ans2[i % 2] += curr_seen.length()
  seen.pop() if seen.length > max_seen
end
# took 4GB and 2:45 to reach 5K steps on test input :)
#puts "ans2: #{seen[0].length()}"
puts "ans2: #{ans2[steps % 2]}"

tgt = 5000
# try to find a period
# step width
(2..cnt.length()/4).each do |period|
  # end point of search: cover whole step width, bot to top
  (steps-period..steps).each do |endpoint|
    # we are interested only in same evennes as a target (26501365 for final)
    next if (endpoint % 2) != (tgt % 2)
    # now step over them, watching diff
    diff = cnt[endpoint] - cnt[endpoint-period]
    depth = 1
    (2..cnt.length/period).each do |attempt|
      att_end = endpoint-attempt*period
      break if cnt[att_end+period] - cnt[att_end] != diff
      depth += 1
    end
    if depth > 1
      # temp kill > 11
      next if period > 11
      puts "period #{period} for #{depth} times from #{endpoint - period*depth} (val #{cnt[endpoint - period*depth]}) upto #{endpoint} (val #{cnt[endpoint]}), diff #{diff}"
    end
  end
end

# cnt[0..-12].each_with_index {|c, i| puts "#{i}: cnt #{c}, vs +11: #{cnt[i+11]-c}"}