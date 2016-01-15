File.open(ARGV[0]).each_line do |line|
  s = line.split('|')
  t = s[0].split.map(&:to_i)
  n = [s[1].to_i, t.length / 2].min
  n.times do
    (1..t.length - 1).each do |i|
      t[i - 1], t[i] = t[i], t[i - 1] if t[i - 1] > t[i]
    end
    (t.length - 2).downto(1) do |i|
      t[i - 1], t[i] = t[i], t[i - 1] if t[i - 1] > t[i]
    end
  end
  puts t.join(' ')
end
