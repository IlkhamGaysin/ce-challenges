sq, n = Hash.new, 0
(0..46340).each do |i| sq[i*i] = true end

File.open(ARGV[0]).each_line do |line|
  if n == 0
    n = line.to_i
  else
    x = line.to_i
    num, l, bot, top = 0, Hash.new, Math.sqrt(x/2).to_i, Math.sqrt(x).to_i
    (bot..top).each do |i|
      t = x - i*i
      l[i*i], num = true, num+1 if sq[t] && !l[t]
    end
    puts num
  end
end
