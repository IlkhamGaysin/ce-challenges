lol, n = Array.new(), 0
File.open(ARGV[0]).each_line do |line|
  if n == 0
    n = line.to_i
  else
    for i in 0...n
      if i == lol.length
        lol[i] = line
        break
      elsif line.length > lol[i].length
        lol[i], line = line, lol[i]
      end
    end
  end
end
puts lol.join()
