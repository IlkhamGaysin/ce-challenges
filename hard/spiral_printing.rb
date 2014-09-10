File.open(ARGV[0]).each_line do |line|
  s = line.split(";")
  row, col, t = s[0].to_i, s[1].to_i, s[2].split(" ")
  i, j = 0, 1
  tn, te, ts, tw = 0, col - 1, row - 1, 0
  print t[0]
  begin
    while j <= te
      print " " + t[i*col + j]
      j += 1
    end
    j, i, tn = j-1, i+1, tn+1
    break if i > ts
    while i <= ts
      print " " + t[i*col + j]
      i += 1
    end
    i, j, te = i-1, j-1, te-1
    break if j < tw
    while j >= tw
      print " " + t[i*col + j]
      j -= 1
    end
    j, i, ts = j+1, i-1, ts-1
    break if i < tn
    while i >= tn
      print " " + t[i*col + j]
      i -= 1
    end
    i, j, tw = i+1, j+1, tw+1
  end until j > te
  puts
end
