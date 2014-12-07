File.open(ARGV[0]).each_line do |line|
  s = line.split
  lo, hi = 0, s[0].to_i
  (1..s.length-1).each do |x|
    c = (lo+hi)%2
    case s[x]
    when "Lower"
      hi = (lo+hi)/2 + c - 1
    when "Higher"
      lo = (lo+hi)/2 + c + 1
    else
      puts (lo+hi)/2 + c
    end
  end
end
