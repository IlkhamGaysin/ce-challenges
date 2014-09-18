for line in io.lines(arg[1]) do
  for i in line:gmatch(".") do
    if i:find("[a-z]") then
      local c, _ = i:gsub("%l", string.upper)
      io.write(c)
    elseif i:find("[A-Z]") then
      local c, _ = i:gsub("%u", string.lower)
      io.write(c)
    else
      io.write(i)
    end
  end
  print()
end
