def balanced(s, c)
  loop do
    if c < 0
      return false
    elsif s == ""
      return c == 0
    elsif s[0].match(/[a-z ]/)
      s = s[1..-1]
    elsif s[-1].match(/[a-z :]/)
      s = s[0..-2]
    elsif s[0] == '('
      c += 1
      s = s[1..-1]
    elsif s[0] == ')'
      c -= 1
      s = s[1..-1]
    elsif s[0] == ':'
      if s[1] == '('
        return balanced(s[2..-1], c) || balanced(s[2..-1], c+1)
      elsif s[1] == ')'
        return balanced(s[2..-1], c) || balanced(s[2..-1], c-1)
      else
        s = s[1..-1]
      end
    else
      return false # illegal input
    end
  end
end

File.open(ARGV[0]).each_line do |line|
  puts balanced(line.sub("\n", ""), 0) ? "YES" : "NO"
end
