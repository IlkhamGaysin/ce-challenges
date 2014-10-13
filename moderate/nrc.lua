for line in io.lines(arg[1]) do
  local a = {}
  for i in line:gmatch(".") do
    if a[i] then
      a[i] = a[i] + 1
    else
      a[i] = 1
    end
  end
  for i in line:gmatch(".") do
    if a[i] == 1 then
      io.write(i)
      break
    end
  end
  io.write("\n")
end
