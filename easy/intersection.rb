File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split(";")
  a, b = s[0].split(",").map(&:to_i), s[1].split(",").map(&:to_i)
  j, r = 0, Array.new
  a.each do |x|
    while j < b.length and x > b[j] do j += 1 end
    if j < b.length and x == b[j] then
      r << x.to_s
      j += 1
    end
  end
  puts r.join(",")
end
