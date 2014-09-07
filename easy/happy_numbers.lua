function happy (a)
  local ret = 0
  while a > 0 do
    local b = math.fmod(a, 10)
    ret = ret +  b * b
    a = math.floor(a/10)
  end
  return ret
end

for line in io.lines(arg[1]) do
  local a = tonumber(line)
  local b = {a}
  for _ = 1, 127 do
    if a <= 1 then
      break
    end
    a = happy(a)
    for _, j in ipairs(b) do
      if j == a then
        a = 0
        break
      end
    end
  end
  if a == 1 then
    print(1)
  else
    print(0)
  end
end
