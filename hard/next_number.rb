def dsig(a)
  d = 0
  while a > 0 do
    r = a % 10
    d += 1 << (3 * r) if r > 0
    a /= 10
  end
  d
end

File.open(ARGV[0]).each_line do |line|
  a = line.to_i
  b, d = a + 9, dsig(a)
  b += 9 while d != dsig(b)
  puts b
end
