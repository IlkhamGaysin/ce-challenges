File.open(ARGV[0]).each_line do |line|
  puts line.scan(/"id": (\d+),/).inject(0) {|sum, i| sum + i[0].to_i} unless line.chomp.empty?
end
