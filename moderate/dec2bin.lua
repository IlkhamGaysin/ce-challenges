for line in io.lines(arg[1]) do
  local a = tonumber(line)
  if a == 0 then
    print(0)
  else
    local b, c = 0, a
    while a > 0 do
      b, a = b * 2 + a % 2, math.floor(a/2)
    end
    while c > 0 do
      io.write(b % 2)
      b, c = math.floor(b/2), math.floor(c/2)
    end
    print()
  end
end
