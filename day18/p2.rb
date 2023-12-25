#!/usr/bin/ruby

def hex2step(h)
  # return [h[0..5].to_i(16)]
  return ["RDLU"[h[7].to_i], h[2..6].to_i(16)]
end

#steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| hex2step(a[2])}
steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| [a[0], a[1].to_i]}
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

# path = [[0,3], [0,6], [2,6], [2,3], [0, 3]] # 4 * 3 = 12 or 2 x 1 = 2 w/o borders
# puts "test path: #{path}"

# 14 points: len 38, area 24, total 62

# stupid: https://www.mathopenref.com/coordpolygonarea.html
# area = abs(sum(x_{n}*y_{n+1} - y_{n}*x_{n+1})/2) (last is x_n*y_1 - y_n*x_1, so init = n)
# area = (path.reduce([0, path[-1]]) {|mem, p| [mem[0] + (p[1]*mem[-1][0] - p[0]*mem[-1][1]), p]}[0]/2).abs
area = 0
pp = path[0]
path.each {|p|
  if p[0] != pp[0]
    w = p[1]         + 1*(p[0] > pp[0] ? 1 : 0)
    h = p[0] - pp[0] + 1*(p[0] > pp[0] ? 1 : -1)
    addArea = w * h
    puts "#{pp} -> #{p}: +area w #{w} h #{h} = #{addArea}"
    area += addArea
  else
    # addx = p[1] <=> pp[1]
    # puts "p #{p}, prev #{pp}: addx #{addx}"
    # area += addx
    puts "#{pp} -> #{p}: skip"
  end
  pp = p
}
puts "area: #{area}"
#puts "ans 2: #{area + pathlen}"