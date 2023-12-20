#!/usr/bin/ruby

# "dir" steps "#color"
steps = ARGF.readlines().map {|l| l.split(' ')}.map {|a| [a[0], a[1].to_i, a[2][1..-2]]}
puts "#{steps}"
sums = steps.reduce(Hash.new(0)) {|msum, m| msum[m[0]] += m[1]; msum }
puts "#{sums}"