p = nil
q = nil
File.open(ARGV[0]).each_line do |line|
  if p.nil?
    p = line.index('C')
    p = line.index('_') if p.nil?
    line[p] = '|'
  else
    loop do
      q = line.index('C')
      break if q.nil? || (p - q).abs < 2
    end
    if q.nil?
      loop do
        q = line.index('_')
        break if (p - q).abs < 2
      end
    end
    if q < p
      sym = '/'
    elsif q == p
      sym = '|'
    else
      sym = '\\'
    end
    line[q] = sym
    p = q
  end
  puts line
end
