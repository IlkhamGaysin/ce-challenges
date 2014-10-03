File.open(ARGV[0]).each_line do |line|
  a, r, i = Array.new(26, 0), 0, 26
  line.chomp.downcase.each_char do |x|
    a[x.ord - 97] += 1 if x =~ /[a-zA-Z]/
  end
  a.sort!
  while i > 0 and a[i-1] > 0
    r += i * a[i-1]
    i -= 1
  end
  puts r
end
