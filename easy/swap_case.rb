File.open(ARGV[0]).each_line do |line|
  puts line.split('').each { |x|
    if /[[:lower:]]/.match(x)
      x.upcase!
    elsif /[[:upper:]]/.match(x)
      x.downcase!
    end
  }.join
end
