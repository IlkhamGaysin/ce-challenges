function add(x, y)
  if #x > #y then
    a, b = x, y
  else
    a, b = y, x
  end
  local r, c = "", 0
  local i, j = #a, #b
  while i > 0 do
    local s, t = tonumber(a:sub(i, i)), 0
    if j > 0 then
      t = tonumber(b:sub(j, j))
    end
    if c+s+t < 10 then
      r = tostring(c+s+t) .. r
      c = 0
    else
      r = tostring((c+s+t)%10) .. r
      c = math.floor((c+s+t)/10)
    end
    i, j = i-1, j-1
  end
  if c > 0 then r = tostring(c) .. r end
  return r
end

for line in io.lines(arg[1]) do
  local a = tonumber(line)
  if a < 2 then
    print(a)
  else
    local b, c = "1", "1"
    while a > 1 do
      a, b, c = a-1, add(b, c), b
    end
    print(b)
  end
end
