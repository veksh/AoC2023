#!/usr/bin/ruby

DEBUG = false

def numVariants(str, strPos, pattern, patPos, memo)
  if memo.has_key?([strPos, patPos])
    return memo[[strPos, patPos]]
  end
  puts "#{' ' * patPos}checking '#{pattern[patPos..]}' on '#{str[strPos..]}' (pos: #{strPos}, #{patPos})" if DEBUG
  ret = -1
  if patPos >= pattern.length()
    # pattern ends: must be only ok left
    if str[strPos..].count("#") == 0
      puts " #{' ' * patPos}hit" if DEBUG
      ret = 1
    else
      puts " #{' ' * patPos}miss" if DEBUG
      ret = 0
    end
  elsif strPos >= str.length()
    # string ends, but pattern does not (as above)
    ret = 0
  elsif pattern[patPos] == "#"
    case str[strPos]
    when "#"
      return numVariants(str, strPos + 1, pattern, patPos + 1, memo)
    when "."
      # return 0
      if patPos == 0 || pattern[patPos-1] != "#"
        ret = numVariants(str, strPos + 1, pattern, patPos, memo)
      else
        puts " #{' ' * patPos}wrong dot" if DEBUG
        ret = 0
      end
    when "?"
      res = numVariants(str, strPos + 1, pattern, patPos + 1, memo)
      if patPos == 0 || pattern[patPos-1] != "#"
        res += numVariants(str, strPos + 1, pattern, patPos, memo)
      end
      ret = res
    end
  else
    # . between groups
    case str[strPos]
    when "#"
      puts " #{' ' * patPos}wrong hash" if DEBUG
      ret = 0
    when "."
      ret = numVariants(str, strPos + 1, pattern, patPos + 1, memo)
    when "?"
      ret = numVariants(str, strPos + 1, pattern, patPos + 1, memo)
    end
  end
  memo[[strPos, patPos]] = ret
  return ret
end

td = ARGF.readlines().map {|l| l.split()}.map {|a| [a[0], a[1].split(',').map(&:to_i)]}
# pattern: to keep only pos: [3,2,1] -> "###.##.#"
# string: clean up leading/trailing dots, compress runs of dots to one
td1 = td.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}

# part1
# td1.each {|p| puts "#{p[0]} vs #{p[1]}: #{numVariants(p[0], 0, p[1], 0, {})}"}
puts td1.map {|p| numVariants(p[0], 0, p[1], 0, {}) }.sum()

# part2: 5 times more, needs memoization
td2 = td.map {|p| [([p[0]]*5).join('?'), p[1]*5] }.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}
# td2.each {|p| puts "#{p[0]} vs #{p[1]}: #{numVariants(p[0], 0, p[1], 0, {})}"}
puts td2.map {|p| numVariants(p[0], 0, p[1], 0, {}) }.sum()
