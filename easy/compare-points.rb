File.open(ARGV[0]).each_line do |line|
  s = line.split.map(&:to_i)
  puts (s[0] == s[2]) ? ((s[1] == s[3]) ? "here" :
                         (s[1] < s[3]) ? "N" : "S") :
       (s[1] == s[3]) ? ((s[0] < s[2]) ? "E" : "W") :
       ((s[0] < s[2]) ? ((s[1] < s[3]) ? "NE" : "SE") :
                        ((s[1] < s[3]) ? "NW" : "SW"))
end
