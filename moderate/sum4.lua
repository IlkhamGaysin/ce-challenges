function numzero(s, c, z)
  if c == 0 then
    if z == 0 then
      return 1
    else
      return 0
    end
  end
  if #s < c then return 0 end
  local t = {}
  for i = 1, #s-1 do t[#t+1] = s[i+1] end
  return numzero(t, c-1, z + s[1]) + numzero(t, c, z)
end

for line in io.lines(arg[1]) do
  local s = {}
  for i in line:gmatch("[-]?%d+") do
    s[#s+1] = tonumber(i)
  end
  print(numzero(s, 4, 0))
end