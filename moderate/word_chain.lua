function lwc(f, c)
  local m = #c/2
  if #f == 0 then return m end
  for ix, i in pairs(f) do
    if #c == 0 or i:sub(-1) == c:sub(1, 1) then
      local g = {}
      for jx, j in pairs(f) do
        if ix ~= jx then g[#g + 1] = j end
      end
      local l = lwc(g, i .. c)
      if l > m then
        m = l
      end
    end
  end
  return m
end

for line in io.lines(arg[1]) do
  local s = {}
  for i in string.gmatch(line, "[^,]+") do
    s[#s + 1] = i:sub(1, 1) .. i:sub(-1)
  end
  local l = lwc(s, "")
  if l == 1 then
    print("None")
  else
    print(l)
  end
end
