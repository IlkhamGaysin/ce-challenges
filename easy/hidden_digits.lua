for line in io.lines(arg[1]) do
  line = line:gsub("[^0-9a-j]", "")
  for i in line:gmatch(".") do
    if i:match("[a-j]") then
      io.write(i:byte()-string.byte("a"))
    else
      io.write(i)
    end
  end
  if #line > 0 then print() else print("NONE") end
end
