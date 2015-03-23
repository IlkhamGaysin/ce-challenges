def pos(l, r)
  (l + ?a.ord).chr + (r + ?1.ord).chr
end

File.open(ARGV[0]).each_line do |line|
  l, r, m = line[0].ord - ?a.ord, line[1].ord - ?1.ord, Array.new
  m << pos(l-2, r-1) if l > 1 and r > 0
  m << pos(l-2, r+1) if l > 1 and r < 7
  m << pos(l-1, r-2) if l > 0 and r > 1
  m << pos(l-1, r+2) if l > 0 and r < 6
  m << pos(l+1, r-2) if l < 7 and r > 1
  m << pos(l+1, r+2) if l < 7 and r < 6
  m << pos(l+2, r-1) if l < 6 and r > 0
  m << pos(l+2, r+1) if l < 6 and r < 7
  puts m.join(" ")
end
