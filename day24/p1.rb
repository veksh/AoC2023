#!/usr/bin/ruby

stones = ARGF.readlines().map do |l|
  Hash[%i(px py pz vx vy vz).zip(l.split(/[,@ ]+/).map(&:to_f))]
end

# test
#minx, maxx, miny, maxy = 7.0, 27.0, 7.0, 27.0
# main
minx, maxx, miny, maxy = 200000000000000.0, 400000000000000.0, 200000000000000.0, 400000000000000.0
ans1 = 0
(0..stones.length()-1).each do |i|
  a = stones[i]
  (i+1..stones.length()-1).each do |j|
    b = stones[j]
    t1 = ((b[:px] - a[:px])/b[:vx] + (a[:py] - b[:py])/b[:vy])/(a[:vx]/b[:vx] - a[:vy]/b[:vy])
    t2 = (a[:vy]*t1 + a[:py] - b[:py])/b[:vy]
    next if t1 < 0 || t2 < 0
    cx, cy = a[:vx]*t1 + a[:px], a[:vy]*t1 + a[:py]
    if cx >= minx && cx <= maxx && cy >= miny && cy <= maxy
      puts "#{i} and #{j} cross at #{cx}:#{cy}"
      ans1 += 1
    end
  end
end
puts "ans1: #{ans1}"