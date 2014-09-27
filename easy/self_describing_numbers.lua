for line in io.lines(arg[1]) do
  n = tonumber(line)
  local m, r = n, 0
  a = {}
  for i = 0, #line-1 do
    a[i] = 0
  end
  self = true
  while m > 0 do
    local v = m % 10
    if v <= #a then
      a[v] = a[v] + 1
    else
      self = false
      break
    end
    m = math.floor(m/10)
  end
  if self then
    for i = 0, #a do
      r = r*10 + a[i]
    end
    if n ~= r then
      self = false
    end
  end
  if self then print(1) else print(0) end
end
