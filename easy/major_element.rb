File.open(ARGV[0]).each_line do |line|
  el, found = {}, false
  s = line.split(",").map(&:to_i)
  for i in s.uniq
    if s.count(i) > s.length/2
      puts i
      found = true
      break
    end
  end
  puts "None" if not found
end
