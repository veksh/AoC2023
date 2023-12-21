#!/usr/bin/ruby

# "dir" steps "#color"
steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| [a[0], a[1].to_i, a[2][1..-2]]}
# puts "#{steps}"

sums = steps.reduce(Hash.new(0)) {|msum, m| msum[m[0]] += m[1]; msum }
puts "#{sums}"

RLUD = {"R" => [0, 1], "L" => [0, -1], "U" => [-1, 0], "D" => [1, 0]}

path = steps.reduce([[0,0]]) {|sofar, step|
  p = sofar[-1].clone()
  p[0] += step[1]*RLUD[step[0]][0]
  p[1] += step[1]*RLUD[step[0]][1]
  sofar.push(p)
}
# puts "#{path}"
min_r, max_r = path.map{|p| p[0]}.min(), path.map{|p| p[0]}.max()
min_c, max_c = path.map{|p| p[1]}.min(), path.map{|p| p[1]}.max()
puts "row: from #{min_r} to #{max_r}"
puts "col: from #{min_c} to #{max_c}"

depth_r, width_c = max_r - min_r, max_c - min_c
puts "rows #{depth_r} cols #{width_c}"

field = (0..depth_r).map {["."] * width_c}
pos = [[0, -1*min_r].max(), [0, -1*min_c].max()]
puts "start #{pos}"
field[pos[0]][pos[1]] = "#"

steps.each do |s|
  dir, len = RLUD[s[0]], s[1]
  r0, r1 = [pos[0], pos[0]+dir[0]*len].sort()
  c0, c1 = [pos[1], pos[1]+dir[1]*len].sort()
  (r0..r1).each do |r|
    (c0..c1).each do |c|
      field[r][c] = "#"
    end
  end
  pos = [pos[0]+dir[0]*len, pos[1]+dir[1]*len]
end
puts "#{field.map {|r| r.join()}.join("\n") }"
