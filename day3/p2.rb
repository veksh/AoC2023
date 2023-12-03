#!/usr/bin/ruby

FILE_NAME = "input_test.txt"

lines = IO.read(FILE_NAME).lines().map(&:chomp)
res = 0
lines.each_with_index do |l, li|
  matches = l.gsub(/\d+/).map{ Regexp.last_match }.map {|m| [m[0].to_i, m.begin(0), m.end(0)]}
  puts "#{sprintf('%3d', li)}: #{l}: #{matches}"
  matches.each do |m|
    puts " studying #{m}"
    case
    when m[1] > 0 && l[m[1]-1] == '*'
      puts "  match before on #{m[1]-1}"
    when m[2] < l.length()-1 && l[m[2]] == '*'
      puts "  match after on #{m[2]}"
    when li > 0 && lines[li-1][[m[1]-1,0].max..[m[2],l.length()-1].min].match?(/\*/)
      puts "  match above on #{lines[li-1].index('*', [m[1]-1,0].max)}"
    when li < lines.length()-1 && lines[li+1][[m[1]-1,0].max..[m[2],l.length()-1].min].match?(/\*/)
      puts "  match below on #{lines[li+1].index('*', [m[1]-1,0].max)}"
    else
      puts "  no match"
      next
    end
    puts "    adding #{m[0]}"
    res += m[0]
  end
  puts "result: #{res}"
end