def choose(n, k)
  return (1..k).inject(1) do |x, i| x*(n-i+1)/i end
end

File.open(ARGV[0]).each_line do |line|
  m, r = line.to_i, 0
  (0...m/2).each do |i|
    r += (-1)**i * choose(m, i) * choose(11*m/2-1-10*i, m-1)
  end
  puts r
end
