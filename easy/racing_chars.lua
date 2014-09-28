for line in io.lines(arg[1]) do
  if #line < 1 then
  elseif curr == nil then
    curr = line:find("C")
    if curr == nil then
      curr = line:find("_")
    end
    print(line:sub(1, curr-1) .. "|" .. line:sub(curr+1))
  else
    px = curr
    local a, b = px-1, px+2
    if a < 1 then a = 1 end
    if b > #line then b = #line end
    curr = line:sub(a, b):find("C")
    if curr == nil then
      curr = line:sub(a, b):find("_")
    end
    curr = curr + a - 1
    io.write(line:sub(1, curr-1))
    if curr < px then
      io.write("/")
    elseif curr == px then
      io.write("|")
    else
      io.write("\\")
    end
    print(line:sub(curr+1))
  end
end
