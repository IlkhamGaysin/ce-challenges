function div(a, b) return math.floor(a/b) end

for line in io.lines(arg[1]) do
  s = {}
  for i in line:gmatch("%S+") do
    s[#s+1] = tonumber(i)
  end
  n, s[1], s[#s] = s[1], s[#s], nil
  table.sort(s)
  u, r = s[div(n, 2) + 1], 0
  for _, i in ipairs(s) do
    r = r + math.abs(i - u)
  end
  print(r)
end
