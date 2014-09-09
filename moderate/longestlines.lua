lol = {}
for line in io.lines(arg[1]) do
  if n == nil then
    n = line
  end
  for i = 1, n do
    if lol[i] == nil or #line > #lol[i] then
      lol[i-1] = lol[i]
      lol[i] = line
    end
  end
end
for i = n, 1, -1 do
  print(lol[i])
end
