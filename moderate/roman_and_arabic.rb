File.open(ARGV[0]).each_line do |line|
  a, num, ar, rr = 0, 0, 0, 0
  line.chomp.each_char do |x|
    if x =~ /\d/
      a = x.to_i
    else
      case x
      when 'M'
        r = 1000
      when 'D'
        r = 500
      when 'C'
        r = 100
      when 'L'
        r = 50
      when 'X'
        r = 10
      when 'V'
        r = 5
      when 'I'
        r = 1
      end
      if r > rr
        num -= ar * rr
      else
        num += ar * rr
      end
      ar, rr = a, r
    end
  end
  puts num + ar * rr
end
