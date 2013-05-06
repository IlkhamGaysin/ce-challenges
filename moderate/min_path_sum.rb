def decode(s)
  if s.length < 2
    return 1
  # should never start with zero: let's ignore if it does
  elsif s[0].to_i > 2 || s[0] == "0"
    return decode(s[1..-1])
  elsif s[1] == "0" || (s[0] == "2" && s[1].to_i > 6)
    return s.length == 2 ? 1 : decode(s[2..-1])
  end
  s.length == 2 ? 2 : decode(s[1..-1]) + decode(s[2..-1])
end

File.open(ARGV[0]).each_line do |line|
  puts decode(line.sub( "\n", ""))
end
