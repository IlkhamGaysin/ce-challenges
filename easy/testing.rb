m = ["Done", "Low", "Medium", "High", "Critical"]

def prio(a)
  case a
  when 0
    0
  when 1..2
    1
  when 3..4
    2
  when 5..6
    3
  else
    4
  end
end

File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split(" | ")
  r = 0
  for ix in (0...s[0].length)
    r += 1 if s[0][ix] != s[1][ix]
  end
  puts m[prio(r)]
end
