for line in io.lines(arg[1]) do
  _, _, n, m = string.find(line, "(%d+),(%d+)")
  print(n - math.floor(n / m) * m)
end
