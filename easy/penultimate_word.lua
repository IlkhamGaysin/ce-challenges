for line in io.lines(arg[1]) do
  local c, p = nil, nil
  for i in line:gmatch("%S+") do
    c, p = i, c
  end
  if p ~= nil then print(p) else print() end
end
