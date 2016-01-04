File.open(ARGV[0]).each_line do |line|
  s, r = line.to_i, 0
  while s > 0 do
    s &= s - 1
    r += 1
  end
  puts r % 3
end
