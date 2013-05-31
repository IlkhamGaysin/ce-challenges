def dist(a, b)
  Math.sqrt(a*a + b*b).to_i
end

File.open(ARGV[0]).each_line do |line|
  s = line.scan(/\(([-0-9]+), ([-0-9]+)\) \(([-0-9]+), ([-0-9]+)\)/)[0].map(&:to_i)
  puts dist(s[0]-s[2], s[1]-s[3])
end
