@sq = { 4 => 2, 9 => 3 }
def region(a, n)
  x, y, z = a % n, a / n, @sq[n]
  return x / z + z * (y / z)
end

File.open(ARGV[0]).each_line do |line|
  s = line.chomp.split(";")
  g = s[1].split(',').map(&:to_i)
  n = s[0].to_i
  valid = true
  row, col, reg = Array.new(9, 0), Array.new(9, 0), Array.new(9, 0)
  g.each_with_index { |i, ix|
    if row[ix / n] & 1 << i > 0
      valid = false
      break
    else
      row[ix / n] |= 1 << i
    end
    if col[ix % n] & 1 << i > 0
      valid = false
      break
    else
      col[ix % n] |= 1 << i
    end
    if reg[region(ix, n)] & 1 << i > 0
      valid = false
      break
    else
      reg[region(ix, n)] |= 1 << i
    end
  }
  puts valid ? "True" : "False"
end
