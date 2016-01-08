File.open(ARGV[0]).each_line do |line|
  s, r = line.split.map(&:to_i), 0
  (1 << s[0]).upto(s[1]) do |i|
    n, a = s[0], i
    while a > 0 && n >= 0 do
      n -= 1 if a % 2 == 0
      a /= 2
    end
    r += 1 if n == 0
  end
  puts r
end
