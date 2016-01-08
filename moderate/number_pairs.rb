File.open(ARGV[0]).each_line do |line|
  s = line.split(';')
  t, u = s[0].split(',').map(&:to_i), s[1].to_i
  n, i, j = false, 0, t.length - 1
  while i < j && 2 * t[i] < u do
    while j > i && t[i] + t[j] >= u do
      if t[i] + t[j] == u
        if n
          print ';'
        else
          n = true
        end
        printf "%d,%d", t[i], t[j]
      end
      j -= 1
    end
    i += 1
  end
  print "NULL" if !n
  puts
end
