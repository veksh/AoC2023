#!/usr/bin/ruby

def hex2step(h)
  # return [h[0..5].to_i(16)]
  return ["RDLU"[h[7].to_i], h[2..6].to_i(16)]
end

steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| hex2step(a[2])}
#steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| [a[0], a[1].to_i]}
puts "#{steps}"

pathlen = steps.map {|s| s[1]}.sum
puts "pathlen: #{pathlen}"

# directions
RLUD = {"R" => [0, 1], "L" => [0, -1], "U" => [-1, 0], "D" => [1, 0]}

# steps -> positions [[row, col]]
path = steps.reduce([[0,0]]) {|sofar, step|
  p = sofar[-1].clone()
  p[0] += step[1]*RLUD[step[0]][0]
  p[1] += step[1]*RLUD[step[0]][1]
  sofar.push(p)
}
puts "#{path}"

# path  = [[0,3], [0,6], [2,6], [2,3], [0, 3]] # 4 * 3 = 12 or 2 x 1 = 2 w/o borders
# steps = [["R", 3], ["D", 2], ["L", 3], ["U", 2]]
# puts "test steps #{steps}, path #{path}" # 14 points: len 38, area 24, total 62

# idea: like https://www.mathopenref.com/coordpolygonarea.html
# - lets add area to the left when moving down, and subtract it when moving up
#   - for down movement, lets count a border too (+1 to width)
#   - for up movement, border is already counted, lets only subtract width (== X position)
# - top borders:
#   - when moving right, lets count them (next down move will not count the top border)
#   - when moving left, nothing neet to be done (they are already included in the area)
# - doing this accurately was pretty hard
area = 1
steps.each_with_index {|step, i|
  dir, len = step
  p = path[i+1]
  addl = 0
  case dir
  when "R"
    addl = len
  #when "L"
  #  addl = len
  when "D"
    addl =    len*(p[1]+1)
  when "U"
    addl = -1*len*(p[1])
  end
  area += addl
  puts "#{i}: #{steps[i]} to #{p}, area + #{addl} = #{area}"
}
puts "area: #{area}"