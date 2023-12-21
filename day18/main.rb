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
puts "minr: #{path.map{|p| p[0]}.min()}, maxr: #{path.map{|p| p[0]}.max()}"
puts "minc: #{path.map{|p| p[1]}.min()}, maxc: #{path.map{|p| p[1]}.max()}"

field = (0..sums["D"]*2).map {["."] * (sums["R"]*2)}
pos = [sums["R"], sums["D"]]
field[pos[0]][pos[1]] = "#"
puts "#{field.map {|r| r.join()}.join("\n") }"
