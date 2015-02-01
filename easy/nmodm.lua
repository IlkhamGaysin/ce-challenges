function div(a, b) return math.floor(a/b) end

for line in io.lines(arg[1]) do
  n, m = line:match("(%d+),(%d+)")
  print(n - div(n, m) * m)
end
