for line in io.lines(arg[1]) do
  a = {}
  for i in line:gmatch("%d+") do a[#a + 1] = tonumber(i) end
  table.sort(a)
  io.write(a[1])
  for i = #a, 2, -1 do a[i] = a[i] - a[i-1] end
  for i = 2, #a do io.write("," .. a[i]) end
  print()
end
