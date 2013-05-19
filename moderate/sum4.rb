def numzero(s, c, z)
  if c == 0
    return z == 0 ? 1 : 0
  end
  return 0 if s.length < c
  numzero(s[1..-1], c-1, z+s[0]) + numzero(s[1..-1], c, z)
end

File.open(ARGV[0]).each_line do |line|
  s = line.sub("\n", "").split(',').map(&:to_i)
  puts numzero(s, 4, 0)
end
