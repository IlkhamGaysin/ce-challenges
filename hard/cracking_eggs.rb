def floors(e, s)
  return 1 if s == 1
  return s if e == 1
  return 0 if e == 0 || s == 0
  return floors(s, s) if e > s
  floors(e - 1, s - 1) + floors(e, s - 1) + 1
end

File.open(ARGV[0]).each_line do |line|
  s, n = line.split.map(&:to_i), 0
  while floors(s[0], n) < s[1] do n += 1 end
  puts n
end
