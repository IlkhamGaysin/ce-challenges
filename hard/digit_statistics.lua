stats = {{2, 4, 8, 6}, {3, 9, 7, 1}, {4, 6}, {5}, {6},
         {7, 9, 3, 1}, {8, 4, 2, 6}, {9, 1}}
for line in io.lines(arg[1]) do
  a, n = line:match("(%d) (%d+)")
  a = a - 1
  m, res = math.floor(n / #stats[a]), {[0]=0,0,0,0,0,0,0,0,0,0}
  for _, i in ipairs(stats[a]) do res[i] = res[i] + m end
  for i = 1, n % #stats[a] do res[stats[a][i]] = res[stats[a][i]] + 1 end
  for i = 0, 9 do
    if i > 0 then io.write(", ") end
    io.write(tostring(i) .. ": " .. tostring(res[i]))
  end
  print()
end
