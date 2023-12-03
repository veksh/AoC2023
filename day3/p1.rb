#!/usr/bin/ruby

FILE_NAME = "input_test.txt"

def l2numpos(l)
  return [[0, 1]]
end

# File.open(FILE_NAME).each_line {|l| puts l}
# IO.foreach(FILE_NAME) {|line|
# lines = IO.readlines(FILE_NAME)

lines = IO.read(FILE_NAME).lines().map(&:chomp)
res = 0
lines.each_with_index do |l, i|
  m = l.gsub(/\d+/).map{ Regexp.last_match }.map {|m| [m[0].to_i, m.begin(0), m.end(0)]}
  puts "#{sprintf('%3d', i)}: #{l}: #{m}"
end