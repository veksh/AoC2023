#!/usr/bin/ruby

def predict(s)
  res = 0
  while !s.all? {|e| e == 0} do
    res += s[-1]
    s = s[1..].zip(s).map{|p| p[0] - p[1]}
  end
  return res
end

td = ARGF.readlines().map {|l| l.split().map(&:to_i)}
# td.each{|l| puts predict(l)}
puts td.map {|l| predict(l)}.sum()


# res = 1
# td[0].zip(td[1]).each do |p|
#   dist = (1..p[0]).map {|h| h*(p[0]-h)}
#   cnt   = dist.count {|d| d > p[1]}
#   puts "time #{p[0]} record #{p[1]}: #{cnt} ways (dist: #{dist})"
#   res *= cnt
# end
# puts "ans: #{res}"