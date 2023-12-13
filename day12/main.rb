#!/usr/bin/ruby

def numVariants(str, strPos, pattern, patPos)
  puts "#{' ' * patPos}checking '#{pattern[patPos..]}' on '#{str[strPos..]}' (pos: #{strPos}, #{patPos})"
  if patPos >= pattern.length()
    # pattern ends: must be only ok left
    if str[strPos..].count("#") == 0
      puts " #{' ' * patPos}hit"
      return 1
    else
      puts " #{' ' * patPos}miss"
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
        puts " #{' ' * patPos}wrong dot"
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
      puts " #{' ' * patPos}wrong hash"
      return 0
    when "."
      return numVariants(str, strPos + 1, pattern, patPos + 1)
    when "?"
      return numVariants(str, strPos + 1, pattern, patPos + 1)
    end
  end
end

td = ARGF.readlines().map {|l| l.split()}.map {|a| [a[0].gsub(%r{\.*$|^\.*}, ''), a[1].split(',').map(&:to_i)]}
puts "#{td}"
tdt = td.map {|p| [p[0].gsub(%r{\.*$|^\.*}, '').gsub(%r{\.+}, '.'), p[1].map {|n| "#" * n}.join(".")]}
tdt.each {|p| puts "#{p[0]} vs #{p[1]}"}
puts
# puts numVariants("???.###", 0, "#.#.###", 0)
# puts numVariants("??.??.?##", 0, "#.#.###", 0)
# puts numVariants("???", 0, "###", 0)
puts numVariants("??#", 0, "##", 0)

