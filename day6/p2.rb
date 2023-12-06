#!/usr/bin/ruby

time, record = ARGF.readlines().map {|l| l.gsub(/[^0-9]/, '').to_i}
puts "time: #{time}, record: #{record}"

# a bit inefficient -- but computers are fast :)
dist = (1..time).map {|h| h*(time-h)}
cnt  = dist.count {|d| d > record}
puts "ans: #{cnt}"