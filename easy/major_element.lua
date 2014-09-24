for line in io.lines(arg[1]) do
  local el, found, n = {}, false, 0
  for i in line:gmatch("%d+") do
    i = tonumber(i)
    if el[i] == nil then
      el[i] = 1
    else
      el[i] = el[i] + 1
    end
    n = n + 1
  end
  for k, v in pairs(el) do
    if v > n/2 then
      found = true
      print(k)
      break
    end
  end
  if not found then print("None") end
end
