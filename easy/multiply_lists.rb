File.open(ARGV[0]).each_line do |line|
  s = line.gsub!( /[^\d] /, "").split.map {|s| s.to_i }
  0.upto(s.length/2-1) { |i|
    print " " if i > 0
    print s[i] * s[s.length/2 + i]
  }
  puts
end
