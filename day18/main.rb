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
puts "row: from #{path.map{|p| p[0]}.min()} to #{path.map{|p| p[0]}.max()}"
puts "col: from #{path.map{|p| p[1]}.min()} to #{path.map{|p| p[1]}.max()}"

max_mr, max_md = sums["R"], sums["D"]
field = (0..max_md*2).map {["."] * max_mr*2}
pos = [max_md-1, max_mr-1]
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
