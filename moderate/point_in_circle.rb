def in_circle(a, b, c)
  a*a + b*b <= c*c
end

File.open(ARGV[0]).each_line do |line|
  s = line.scan(/\(([-0-9.]+), ([-0-9.]+)\); Radius: ([-0-9.]+); Point: \(([-0-9.]+), ([-0-9.]+)/)[0].map(&:to_f)
  puts in_circle(s[0]-s[3], s[1]-s[4], s[2])
end
