for line in io.lines(arg[1]) do
  local r = line:match(", (.*)")
  local s = line:match("^(.*), "):gsub("[" .. r .. "]", ""):gsub("^%s+", ""):gsub("%s+$", "")
  print(s)
end
