class MinimumScalar

def initialize(va, vb)
  @va = va.sort.reverse
  @vb = vb.sort
end

def va
  @va
end

def vb
  @vb
end

def solve
  total = 0
  va.each_with_index do |a, i|
    total += a * vb[i]
  end
  total
end

end

txt = open(ARGV.first)
lines = txt.readlines
cases = lines[0].to_i
i = 1
index = 0
cases.times do
  index += 1
  va = lines[i+1].split.map {|i| i.to_i}
  vb = lines[i+2].split.map {|i| i.to_i}
  @product = MinimumScalar.new(va, vb)
  puts "Case #{index}: #{@product.solve}"
  i += 3
end


