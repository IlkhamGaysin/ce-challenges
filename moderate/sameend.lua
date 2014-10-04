for line in io.lines(arg[1]) do
  local sep = line:find(",")
  if sep then
    if line:sub(1, sep-1):find(line:sub(sep+1) .. "$") then
      print(1)
    else
      print(0)
    end
  end
end
