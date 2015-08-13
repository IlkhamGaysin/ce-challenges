File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split(" | ").map { |x| x.split.map(&:to_i) }
  for i in 1...s.length
    s[i].each_with_index do |x, j|
      s[0][j] = x if s[0][j] < x
    end
  end
  puts s[0].join(" ")
end
