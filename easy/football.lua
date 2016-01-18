for line in io.lines(arg[1]) do
  ix, r, t, u = 1, {}, {}, {}
  for i in string.gmatch(line, "%S+") do
    if i == "|" then
      ix = ix + 1
    else
      tx = tonumber(i)
      if t[tx] == nil then
        t[tx] = tostring(ix)
        table.insert(u, tx)
      else
        t[tx] = t[tx] .. ',' .. tostring(ix)
      end
    end
  end
  table.sort(u)
  for k, v in ipairs(u) do
    r[#r+1] = tostring(v) .. ":" .. t[u[k]] .. ";"
  end
  print(table.concat(r, " "))
end
