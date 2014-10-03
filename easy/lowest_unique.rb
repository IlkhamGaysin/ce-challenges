File.open(ARGV[0]).each_line do |line|
  t = line.split.map(&:to_i)
  s = t.sort
  while s.length > 1 do
    if s[0] == s[1]
      s.delete(s[0])
    else
      break
    end
  end
  puts s.length == 0 ? 0 : t.find_index(s[0])+1
end
