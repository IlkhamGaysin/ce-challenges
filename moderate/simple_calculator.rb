$p = { "^" => 3, "*" => 2, "/" => 2, "+" => 1, "-" => 1 }
$error = -999999

def operate(a, b, c)
  r = 0
  case b
  when '^'
    r = a ** c
  when '*'
    r = a * c
  when '/'
    r = a / c
  when '+'
    r = a + c
  when '-'
    r = a - c
  else
    print "Unexpected operator: " + b
    return $error
  end
  r
end

def parse(s)
  r = []
  state = 0
  while s.length > 0
    case state
    when 0
      if s[0] == '('
        r << '('
        s = s[1..-1]
      else
        n = s[/^[-]?[0-9]+([.][0-9]+)?/]
        if s[0] == '-' and s[1] == '('
          r << '['
          s = s[2..-1]
        elsif not n
          print "Input error: expected number, opening bracket, or unary -, got " + s
          return $error
        else
          r << n.to_f
          s = s[n.length..-1]
          state = 1
        end
      end
    when 1
      if '^*/+-'.include? s[0]
        r << s[0]
        s = s[1..-1]
        state = 0
        if r.length > 3 and r[-2] != ')' and r[-3] != '(' and r[-3] != '[' and r[-4] != ')'
          unless '^*/+-'.include? r[-3]
            print "Expected operator, got " + r[-3]
            return $error
          end
          if $p[r[-3]] >= $p[r[-1]]
            r[-4] = operate(r[-4], r[-3], r[-2])
            return $error if r[-4] == $error
            r = r[0..-4] << r[-1]
          end
        end
      elsif s[0] == ')'
        r << ')'
        s = s[1..-1]
        while r.length > 3 and r[-2] != ')' and r[-3] != '(' and r[-3] != '[' and r[-4] != ')'
          r[-4] = operate(r[-4], r[-3], r[-2])
          return $error if r[-4] == $error
          r = r[0..-4] << r[-1]
        end
        if r.length > 3 and r[-3] == '('
          r = r[0...-3] << r[-2]
        elsif r.length > 3 and r[-3] == '['
          r = r[0...-3] << -r[-2]
        elsif r.length == 3  and r[0] == '('
          r = [r[1]]
        elsif r.length == 3  and r[0] == '['
          r = [-r[1]]
        end
      else
        print "Input error: expected operator or closing bracket, got " + s
        return $error
      end
    end
  end
  while r.length > 1
    r[-3] = operate(r[-3], r[-2], r[-1])
    return $error if r[-3] == $error
    r = r[0..-3]
  end
  r
end

File.open(ARGV[0]).each_line do |line|
  s = line.sub("\n", "").gsub("\t", '').gsub(' ', '')
  p = parse(s)
  if p == $error
    puts " on input " + s
  elsif p[0] and p[0].class == Float
    t = "%.5f" % p[0]
    puts t.sub(/[.]?0+$/, "")
  end
end
