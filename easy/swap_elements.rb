File.open(ARGV[0]).each_line do |line|
  s, swap, r = line.chomp.split, false, Array.new
  s.each do |i|
    if swap then
      t = i.chomp(",").split("-").map(&:to_i)
      r[t[0]], r[t[1]] = r[t[1]], r[t[0]]
    elsif i == ":" then
      swap = true
    else
      r << i
    end
  end
  puts r.join(" ")
end
