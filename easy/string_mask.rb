File.open(ARGV[0]).each_line do |line|
  s = line.split
  for i in 0..s[0].length
    print s[1][i] == "1" ? s[0][i].upcase : s[0][i]
  end
  puts
end
