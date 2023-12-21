#!/usr/bin/ruby

def hex2step(h)
  # return [h[0..5].to_i(16)]
  return ["RDLU"[h[7].to_i], h[2..6].to_i(16)]
end

# input into [["dir" steps "#color"]]
steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| hex2step(a[2])}
puts "#{steps}"
# first check: move range; unused
sums = steps.reduce(Hash.new(0)) {|msum, m| msum[m[0]] += m[1]; msum }
puts "#{sums}"

# pretty hopeless at this point :)
exit(0)

# directions
RLUD = {"R" => [0, 1], "L" => [0, -1], "U" => [-1, 0], "D" => [1, 0]}

# steps -> positions [[row, col]]
path = steps.reduce([[0,0]]) {|sofar, step|
  p = sofar[-1].clone()
  p[0] += step[1]*RLUD[step[0]][0]
  p[1] += step[1]*RLUD[step[0]][1]
  sofar.push(p)
}

# range of rows and cols (actually y and x, will offset later)
min_r, max_r = path.map{|p| p[0]}.min(), path.map{|p| p[0]}.max()
min_c, max_c = path.map{|p| p[1]}.min(), path.map{|p| p[1]}.max()
puts "row: from #{min_r} to #{max_r}"
puts "col: from #{min_c} to #{max_c}"

# construct field from range
depth_r, width_c = max_r - min_r, max_c - min_c
puts "rows #{depth_r} cols #{width_c}"
field = (0..depth_r).map {["."] * width_c}

# start pos: offset to the middle if min_x/y < 0
start_pos = [[0, -1*min_r].max(), [0, -1*min_c].max()]
puts "start #{start_pos}"
field[start_pos[0]][start_pos[1]] = "#"

# paint path over the field
num_border = 0
pos = start_pos.clone()
steps.each do |s|
  dir, len = RLUD[s[0]], s[1]
  r0, r1 = [pos[0], pos[0]+dir[0]*len].sort()
  c0, c1 = [pos[1], pos[1]+dir[1]*len].sort()
  (r0..r1).each do |r|
    (c0..c1).each do |c|
      field[r][c] = "#"
      num_border += 1
    end
  end
  num_border -= 1
  pos = [pos[0]+dir[0]*len, pos[1]+dir[1]*len]
end
puts "borders: #{num_border}"
# field[start_pos[0]][start_pos[1]] = "X"
# puts "#{field.map {|r| r.join()}.join("\n") }"

# now flood fill
# choice of start point is a bit hard: lets cheat, it is +1 for both of cases
fill_start = [start_pos[0]+1, start_pos[1]+1]
num_intl = 0
queue = [fill_start]
while queue.length() > 0 do
  p = queue.pop()
  next if field[p[0]][p[1]] == "#" || field[p[0]][p[1]] == "@"
  field[p[0]][p[1]] = "@"
  num_intl += 1
  RLUD.values().each {|dr, dc| queue.unshift([p[0]+dr, p[1]+dc])}
end
puts "inside: #{num_intl}"
# puts "#{field.map {|r| r.join()}.join("\n") }"
puts "ans 1: #{num_border + num_intl}"