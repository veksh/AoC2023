#!/usr/bin/ruby

DEBUG = false

def numVariants(str, strPos, pattern, patPos)
  puts "#{' ' * patPos}checking '#{pattern[patPos..]}' on '#{str[strPos..]}' (pos: #{strPos}, #{patPos})" if DEBUG
  if patPos >= pattern.length()
    # pattern ends: must be only ok left
    if str[strPos..].count("#") == 0
      puts " #{' ' * patPos}hit" if DEBUG
      return 1
    else
      puts " #{' ' * patPos}miss" if DEBUG
      return 0
    end
  end
  if strPos >= str.length()
    # string ends, but pattern does not (as above)
    return 0
  end
  if pattern[patPos] == "#"
    case str[strPos]
    when "#"
      return numVariants(str, strPos + 1, pattern, patPos + 1)
    when "."
      # return 0
      if patPos == 0 || pattern[patPos-1] != "#"
        return numVariants(str, strPos + 1, pattern, patPos)
      else
        puts " #{' ' * patPos}wrong dot" if DEBUG
        return 0
      end
    when "?"
      res = numVariants(str, strPos + 1, pattern, patPos + 1)
      if patPos == 0 || pattern[patPos-1] != "#"
        res += numVariants(str, strPos + 1, pattern, patPos)
      end
      return res
    end
  else
    # . between groups
    case str[strPos]
    when "#"
      puts " #{' ' * patPos}wrong hash" if DEBUG
      return 0
    when "."
      return numVariants(str, strPos + 1, pattern, patPos + 1)
    when "?"
      return numVariants(str, strPos + 1, pattern, patPos + 1)
    end
  end
end

td = ARGF.readlines().map {|l| l.split()}.map {|a| [a[0], a[1].split(',').map(&:to_i)]}
# pattern: to keep only pos: [3,2,1] -> "###.##.#"
# string: clean up leading/trailing dots, compress runs of dots to one
td1 = td.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}
# td1.each {|p| puts "#{p[0]} vs #{p[1]}"}

# puts numVariants("?.?.?.#", 0, "#.#", 0)
# puts numVariants("???.###", 0, "#.#.###", 0)
# puts numVariants("??.??.?##", 0, "#.#.###", 0)
# puts numVariants("?###????????", 0, "###.##.#", 0)
puts td1.map {|p| numVariants(p[0], 0, p[1], 0) }.sum()

# 5 times more
td2 = td.map {|p| [([p[0]]*5).join('?'), p[1]*5] }.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}
td2.each {|p| puts "#{p[0]} vs #{p[1]}: #{numVariants(p[0], 0, p[1], 0)}"}
puts td2.map {|p| numVariants(p[0], 0, p[1], 0) }.sum()
