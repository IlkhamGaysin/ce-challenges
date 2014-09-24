for line in io.lines(arg[1]) do
  local a, n = {}, 0
  for i in line:gmatch("%d+") do
    a[#a+1] = tonumber(i)
  end
  n, a[#a] = math.min(a[#a], #a-2), nil
  for i = 1, n do
    for j = 2, #a do
      if a[j] < a[j-1] then a[j], a[j-1] = a[j-1], a[j] end
    end
  end
  for i = 1, #a do
    if i > 1 then io.write(" ") end
    io.write(a[i])
  end
  print()
end
