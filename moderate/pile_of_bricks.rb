File.open(ARGV[0]).each_line do |line|
  r, s = Array.new, line.scan(/\[([-0-9.]+),([-0-9.]+)\] \[([-0-9.]+),([-0-9.]+)\]/)[0].map(&:to_i)
  h = [(s[0].to_i - s[2].to_i).abs, (s[1].to_i - s[3].to_i).abs].sort
  t = line.scan(/\(([0-9]+) \[([-0-9.]+),([-0-9.]+),([-0-9.]+)\] \[([-0-9.]+),([-0-9.]+),([-0-9.]+)\]\)/) {|i,s1,s2,s3,s4,s5,s6|
    b = [(s1.to_i - s4.to_i).abs, (s2.to_i - s5.to_i).abs, (s3.to_i - s6.to_i).abs].sort
    r << i if b[0] <= h[0] and b[1] <= h[1] # or diagonal(b, h)
  }
  puts r.empty? ? "-" : r.map(&:to_i).sort.join(",")
end
