#!/usr/bin/ruby

stones = ARGF.readlines().map do |l|
  Hash[%i(px py pz vx vy vz).zip(l.split(/[,@ ]+/).map(&:to_f))]
end
puts "#{stones}"
a, b = stones[0], stones[1]
t = ((b[:px] - a[:px])/b[:vx] + (a[:py] - b[:py])/b[:vy])/(a[:vx]/b[:vx] - a[:vy]/b[:vy])
cx, cy = a[:vx]*t + a[:px], a[:vy]*t + a[:py]
puts "#{cx}, #{cy}"