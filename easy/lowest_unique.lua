for line in io.lines(arg[1]) do
  local n, c, f = {}, 1, false
  for i in line:gmatch("%S+") do
    local ix = tonumber(i)
    if n[ix] ~= nil then
      n[ix] = 0
    else
      n[ix] = c
    end
    c = c + 1
  end
  for i = 1, 9 do
    if n[i] ~= nil and n[i] > 0 then
      print(n[i])
      f = true
      break
    end
  end
  if not f then
    print(0)
  end
end