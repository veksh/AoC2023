#!/usr/bin/ruby

FILE_NAME = "input.txt"

lines = IO.read(FILE_NAME).lines().map(&:chomp)
res = 0
star2nums = {}
lines.each_with_index do |l, li|
  matches = l.gsub(/\d+/).map{ Regexp.last_match }.map {|m| [m[0].to_i, m.begin(0), m.end(0)]}
  puts "#{sprintf('%3d', li)}: #{l}: #{matches}"
  matches.each do |m|
    puts " studying #{m}"
    if m[1] > 0 && l[m[1]-1] == '*'
      puts "  match before on #{m[1]-1}"
      (star2nums[[li, m[1]-1]] ||= []).push(m[0])
    end
    if m[2] < l.length()-1 && l[m[2]] == '*'
      puts "  match after on #{m[2]}"
      (star2nums[[li, m[2]]] ||= []).push(m[0])
    end
    if li > 0 && lines[li-1][[m[1]-1,0].max..[m[2],l.length()-1].min].match?(/\*/)
      puts "  match above on #{lines[li-1].index('*', [m[1]-1,0].max)}"
      (star2nums[[li-1, lines[li-1].index('*', [m[1]-1,0].max)]] ||= []).push(m[0])
    end
    if li < lines.length()-1 && lines[li+1][[m[1]-1,0].max..[m[2],l.length()-1].min].match?(/\*/)
      puts "  match below on #{lines[li+1].index('*', [m[1]-1,0].max)}"
      (star2nums[[li+1, lines[li+1].index('*', [m[1]-1,0].max)]] ||= []).push(m[0])
    end
  end
  puts "result: #{star2nums.values()}"
  puts star2nums.values.reduce(0) {|s, v| s += v[0]*v[1] if v.length() == 2; s}
end