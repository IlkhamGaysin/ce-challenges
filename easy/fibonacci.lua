for line in io.lines(arg[1]) do
  local a = tonumber(line)
  if a < 2 then
    print(a)
  else
    local b = {1, 1, 0}
    local c = 1
    while a > c do
      c = c + 1
      b[1] = b[2] + b[3]
      b[3] = b[2]
      b[2] = b[1]
    end
    print(b[1])
  end
end
