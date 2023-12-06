#!/usr/bin/ruby

td = ARGF.readlines().map {|l| l.split()[1..]}
puts "times: #{td[0]}, distances: #{td[1]}"

