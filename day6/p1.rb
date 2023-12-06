#!/usr/bin/ruby

td = ARGF.readlines().map {|l| l.split()[1..].map(&:to_i)}
puts "times: #{td[0]}, records: #{td[1]}"

res = 1
td[0].zip(td[1]).each do |p|
  dist = (1..p[0]).map {|h| h*(p[0]-h)}
  cnt   = dist.count {|d| d > p[1]}
  puts "time #{p[0]} record #{p[1]}: #{cnt} ways (dist: #{dist})"
  res *= cnt
end
puts "ans: #{res}"