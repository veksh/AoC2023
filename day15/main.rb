#!/usr/bin/ruby

def HASH(str)
  str.bytes().reduce(0) {|s, o| (s + o)*17 % 256}
end

steps = ARGF.read().strip().split(',')
ans1 = steps.map {|s| HASH(s)}.sum()
puts "ans1: #{ans1}"