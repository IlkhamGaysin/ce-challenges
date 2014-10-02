function rev(a)
  local r = 0
  while a >= 1 do
    r, a = 10*r + a%10, math.floor(a/10)
  end
  return r
end

for line in io.lines(arg[1]) do
  local a, i = tonumber(line), 0
  while i < 100 do
    local r = rev(a)
    if r == a then break end
    a, i = a+r, i+1
  end
  if i < 100 then print(i .. " " .. a) else print("not found") end
end
