File.open(ARGV[0]).each_line do |s|
  r, n = 1, s.length / 2
  for i in 0..n
    if (s[i] == "A" && s[i+n] == "B") || (s[i] == "B" && s[i+n] == "A")
      r = 0
      break
    elsif s[i] == "*" && s[i+n] == "*"
      r *= 2
    end
  end
  puts r
end
