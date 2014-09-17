File.open(ARGV[0]).each_line do |line|
  s, m, n = line.chomp, -1, -1
  begin
    m, n = n, s.index(s[-1], n+1)
  end until n == s.length-1
  puts m
end
