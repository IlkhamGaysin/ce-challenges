for line in io.lines(arg[1]) do
  f, line = false, line:gsub("%u", string.lower)
  for i = string.byte('a'), string.byte('z') do
    if not line:find(string.char(i)) then
      io.write(string.char(i))
      f = true
    end
  end
  if f then io.write('\n') else io.write("NULL\n") end
end
