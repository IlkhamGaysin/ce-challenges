File.open(ARGV[0]).each_line do |line|
  l, u = line.count("a-z").to_f, line.count("A-Z").to_f
  if l+u > 0
    printf("lowercase: %.2f uppercase: %.2f\n", 100*l/(l+u), 100*u/(l+u))
  else
    puts "lowercase: 0.00 uppercase: 0.00"
  end
end
