#!/usr/bin/ruby

maze = ARGF.readlines().map {|l| l.strip().chars()}
puts "#{maze}"