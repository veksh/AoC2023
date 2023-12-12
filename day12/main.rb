#!/usr/bin/ruby

def numVariants(str, strPos, pattern, patterPos)
  return 1
end

td = ARGF.readlines().map {|l| l.split()}.map {|a| [a[0].gsub(%r{\.*$|^\.*}, ''), a[1].split(',').map(&:to_i)]}
puts "#{td}"
tdt = td.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}
puts "#{tdt}"

puts numVariants("???.###", 0, "#.#.###", 0)