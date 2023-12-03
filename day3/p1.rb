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
  matches = l.gsub(/\d+/).map{ Regexp.last_match }.map {|m| [m[0].to_i, m.begin(0), m.end(0)]}
  puts "#{sprintf('%3d', i)}: #{l}: #{matches}"
  matches.each do |m|
    puts " studying #{m}"
    case
    when m[1] > 0 && l[m[1]-1] != '.'
      puts "  match before"
    when m[2] < l.length()-1 && l[m[2]] != '.'
      puts "  match after"
    else
      puts "  no match"
      next
    end
      puts "    adding #{m[0]}"
      res += m[0]
  end
end