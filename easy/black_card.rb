File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split(" | ")
  t, m = s[0].split(), s[1].to_i
  while t.length > 1 do
    n = m % t.length - 1
    t.delete_at((n >= 0) ? n : t.length - 1)
  end
  puts t[0]
end
