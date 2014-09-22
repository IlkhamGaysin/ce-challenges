File.open(ARGV[0]).each_line do |line|
  if not line.include?(';') then next end
  s = line.split(';')
  s0, s1 = s[0].split, s[1].split.map(&:to_i)
  (1..s0.length).each do |i|
    if not s1.include?(i) then s1 << i end
  end
  puts s0.sort_by.with_index { |_, i| s1[i] }.join(' ')
end
