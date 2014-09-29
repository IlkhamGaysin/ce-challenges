File.open(ARGV[0]).each_line do |line|
  s = line.split(";")
  if s.length > 1
    m = Hash.new
    s[1].split(",").map(&:to_i).each do |i|
      if m[i]
        puts i
        break
      else
        m[i] = true
      end
    end
  end
end
