for line in io.lines(arg[1]) do
  s = ""
  for i in string.gmatch(line, "%S+") do
    s = i .. " " .. s
  end
  print(s:sub(1, -2))
end
