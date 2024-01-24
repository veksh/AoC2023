#!/usr/bin/ruby

require "z3"

stones = ARGF.readlines().map do |l|
  Hash[%i(px py pz vx vy vz).zip(l.split(/[,@ ]+/).map(&:to_f))]
end

solver = Z3::Solver.new

r = %i[px py pz vx vy vz].to_h{|l| [l, Z3.Real(l.to_s)]}

(0..2).each do |i|
  t = Z3.Real("t#{i}")
  s = stones[i]
  solver.assert r[:px] + t*r[:vx] == s[:px] + t*s[:vx]    
  solver.assert r[:py] + t*r[:vy] == s[:py] + t*s[:vy]
  solver.assert r[:pz] + t*r[:vz] == s[:pz] + t*s[:vz]
end

if solver.satisfiable?
  model = solver.model
  model.each do |var, value|
    puts "#{var}=#{value}"
  end
  puts "ans2: #{model[r[:px] + r[:py] + r[:pz]]}"
else
  puts "There are no solutions"
end

