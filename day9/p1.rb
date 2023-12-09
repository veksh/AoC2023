#!/usr/bin/ruby

def predict(s)
  res = 0
  while !s.all? {|e| e == 0} do
    res += s[-1]
    s = s[1..].zip(s).map{|p| p[0] - p[1]}
  end
  return res
end

td = ARGF.readlines().map {|l| l.split().map(&:to_i)}
puts td.map {|l| predict(l)}.sum()
