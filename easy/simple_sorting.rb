File.open(ARGV[0]).each_line do |line|
  s = line.split.sort! { |x,y| x.to_f <=> y.to_f }
  puts s.join(" ")
end
