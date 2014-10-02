File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split
  if s.length > 1
    c, o, r, v = 0, 1, 0, 0
    s[1].each_char do |i|
      if i =~ /[a-z]/
        v, c = v*10 + s[0][c].to_i, c+1
      elsif i == '+'
        o, r, v = 1, r + v*o, 0
      elsif i == '-'
        o, r, v = -1, r + v*o, 0
      end
    end
    puts r + v*o
  end
end
