File.open(ARGV[0]).each_line do |line|
  s = line.split.map(&:to_i)
  f = false
  s.permutation.each do |i|
    (0..2).each do |o1|
      case o1
      when 0 then r1 = i[0] + i[1]
      when 1 then r1 = i[0] - i[1]
      else r1 = i[0] * i[1]
      end
      (0..2).each do |o2|
        case o2
        when 0 then r2 = r1 + i[2]
        when 1 then r2 = r1 - i[2]
        else r2 = r1 * i[2]
        end
        (0..2).each do |o3|
          case o3
          when 0 then r3 = r2 + i[3]
          when 1 then r3 = r2 - i[3]
          else r3 = r2 * i[3]
          end
          (0..2).each do |o4|
            case o4
            when 0 then r4 = r3 + i[4]
            when 1 then r4 = r3 - i[4]
            else r4 = r3 * i[4]
            end
            if r4 == 42
              f = true
              break
            end
          end
          break if f
        end
        break if f
      end
      break if f
    end
    break if f
  end
  puts f ? 'YES' : 'NO'
end
