File.open(ARGV[0]).each_line do |line|
  s = line.split(";")
  n, t, m, c = s[0].to_i, s[1].split.map(&:to_i), 0, 0
  for i in 0...n
    c += t[i]
  end
  m = c if c > m
  for i in n...t.length
    c = c - t[i-n] + t[i]
    m = c if c > m
  end
  puts(m)
end
