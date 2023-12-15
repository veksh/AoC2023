#!/usr/bin/ruby

def HASH(str)
  str.bytes().reduce(0) {|s, o| (s + o)*17 % 256}
end

steps = ARGF.read().strip().split(',')
ans1 = steps.map {|s| HASH(s)}.sum()
puts "ans1: #{ans1}"

boxes = 255.times.map {[]}
steps.each do |s|
  label, fl = s.split(/[=-]/)
  boxNo = HASH(label)
  if fl != nil
    if (pp = boxes[boxNo].index {|p| p[0] == label}) !=nil
      boxes[boxNo][pp][1] = fl
    else
      boxes[boxNo].append([label, fl])
    end
  else
    if (pp = boxes[boxNo].index {|p| p[0] == label}) !=nil
      boxes[boxNo].delete_at(pp)
    end
  end
end

ans2 = boxes.each_with_index.map { |b, bNum|
  b.each_with_index.map { |l, lPos|
    (bNum + 1)*(lPos + 1)*(l[1].to_i) }.sum
}.sum
puts "ans2: #{ans2}"
