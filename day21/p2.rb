#!/usr/bin/ruby

require "set"

steps = ARGV.pop().to_i || 26501365 # 5*11*481843 steps = 202300*131 + 65
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

period = 131
tries = 5
sumB, sumD = 0, 0
(steps-(period-1)*2..steps).each do |stepno|
  next if stepno % 2 != 1
  diff = cnt[stepno] - cnt[stepno - period]
  (1..tries).each do |try|
    diffN = cnt[stepno-period*try] - cnt[stepno-period*(try+1)]
    if diffN != diff
      puts "#{stepno}: mismatch on try #{try} (want #{diff} got #{diffN})"
    end
  end
  puts "#{stepno} (val #{cnt[stepno]}): diff #{diff}"
  sumB += cnt[stepno]
  sumD += diff
end
puts "sumP #{sumB}, sumD #{sumD}"
# test input: period = 11
# - 600 steps: 240644, sumP 17338, sumD 324
# => 622  steps (+1 cycle):    240644+2*(0+17338+324*2)/2 = 258630
# => 5000 steps (+200 cycles): 240644+200*(17338+324*2+17986+324*200*2)/2

# prod input: period = 131, 26501365 = 65+20*131 (=2685) + 202280*131
# - 2685 steps: 6388728, sumP 1185578, sumD 60788
# => + 101140 cycles: 6388728+101140*(1185578+60788*2 + 1185578+60788*2*101140)/2 = 621944727930768
puts "ans2 = 621944727930768 :)"
