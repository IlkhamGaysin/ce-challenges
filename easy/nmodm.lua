function div(a, b) return math.floor(a/b) end

for line in io.lines(arg[1]) do
  _, _, n, m = string.find(line, "(%d+),(%d+)")
  print(n - div(n, m) * m)
end
