#!/usr/bin/ruby
td = ARGF.readlines().map {|l| l.split()}.map {|a| [a[0], a[1].split(',').map(&:to_i)]}
puts "#{td}"