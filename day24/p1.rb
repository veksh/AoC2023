#!/usr/bin/ruby

stones = ARGF.readlines().map do |l|
  Hash[%i(px py pz vx vy vz).zip(l.split(/[,@ ]+/).map(&:to_f))]
end
puts "#{stones}"