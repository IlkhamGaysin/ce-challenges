def contains(s, p)
  until s.empty? or p.empty?
    if s[0] < p[0]
      s = s[1..-1]
    elsif s[0] == p[0]
      s, p = s[1..-1], p[1..-1]
    else
      return false
    end
  end
  p.empty?
end

File.open(ARGV[0]).each_line do |line|
  r, s = Array.new, line.chomp.split(" | ")
  t, p = s[0].split, s[1].chars.sort.join
  u = t.map { |i| i.chars.sort.join }
  t.each_with_index { |i, ix| if contains(u[ix], p) then r << i end }
  puts r.empty? ? "False" : r.join(" ")
end
