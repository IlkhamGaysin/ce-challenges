lol = {}
for line in io.lines(arg[1]) do
  if n == nil then
    n = line
  end
  for i = 1, n do
    if lol[i] == nil then
      lol[i] = line
      break
    elseif #line > #lol[i] then
      lol[i], line = line, lol[i]
    end
  end
end
print(table.concat(lol, "\n"))
