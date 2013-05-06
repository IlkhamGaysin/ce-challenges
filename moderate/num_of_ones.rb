File.open(ARGV[0]).each_line do |line|
  s = line.to_i
  r = 0
  while s > 0 do
    r += s & 1
    s /= 2
  end
  puts r
end
