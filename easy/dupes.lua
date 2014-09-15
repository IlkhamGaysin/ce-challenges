for line in io.lines(arg[1]) do
  first = true
  for i in line:gmatch("%d+") do
    if first then
      io.write(i)
      first = false
    elseif p ~= i then
      io.write("," .. i)
    end
    p = i
  end
  print()
end
