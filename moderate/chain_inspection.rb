File.open(ARGV[0]).each_line do |line|
  b, s, m, been = "BEGIN", line.chomp.split(";"), Hash.new, Hash.new
  s.each do |x|
    t = x.split("-")
    m[t[0]] = t[1]
  end
  s.length.times do
    if !(m.has_key?(b)) || (been.has_key?(m[b]))
      b = nil
      break
    end
    b = m[b]
    been[b] = true
  end
  puts b == "END" ? "GOOD" : "BAD"
end
