a = {}
for line in io.lines(arg[1]) do
  ix, cm, m, n = 0, 0, 0, 0
  for i in line:gmatch("%-?%d+") do
    c = tonumber(i)
    if ix == 0 then
      n = c
    else
      if ix <= n then
        cm = cm + c
        a[ix-1] = c
      else
        cm = cm + c - a[(ix-1) % n]
        a[(ix-1) % n] = c
      end
      if ix >= n and cm > m then
        m = cm
      end
    end
    ix = ix + 1
  end
  print(m)
end
