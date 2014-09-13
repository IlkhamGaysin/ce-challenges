for line in io.lines(arg[1]) do
  local n, i = tonumber(line), 0
  while n > 0 do
    if n%2 > 0 then i = i + 1 end
    n = math.floor(n/2)
  end
  print(i%3)
end
