#!/usr/bin/ruby

steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| [a[0], a[1].to_i, a[2][1..-2]]}
puts "#{steps}"