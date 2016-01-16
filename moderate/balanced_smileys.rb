def balanced(s, c)
  loop do
    return false if c < 0
    return c == 0 if s.empty?
    return balanced(s[2..-1], c) || balanced(s[2..-1], c + 1) if s[0..1] == ':('
    return balanced(s[2..-1], c) || balanced(s[2..-1], c - 1) if s[0..1] == ':)'
    case s[0]
    when '(' then c += 1
    when ')' then c -= 1
    end
    s = s[1..-1]
  end
end

File.open(ARGV[0]).each_line do |line|
  puts balanced(line.chomp, 0) ? 'YES' : 'NO'
end
