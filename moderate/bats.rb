File.open(ARGV[0]).each_line do |line|
  s = line.split.map(&:to_i)
  l, d, n, count, t, tx, i = s[0], s[1], s[2], 0, 0, 6-s[1], 6
  while i <= l-6 do
    if i > tx-d then
      i = tx
      if t == n then
        tx = l-6+d
      else
        tx = s[t+3]
        t = t + 1
      end
    else
      count = count + 1
    end
    i = i + d
  end
  puts count
end
