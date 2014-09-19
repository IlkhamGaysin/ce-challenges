File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split.map(&:to_i)
  jolly, n = true, s[0]
  if n > 1 then
    u = Array.new(n-1, false)
    (2..n).each do |i|
      x = (s[i] - s[i-1]).abs
      if x >= n or x == 0 or u[x-1] then
        jolly = false
        break
      else
        u[x-1] = true
      end
    end
  end
  puts jolly ? "Jolly" : "Not jolly"
end
