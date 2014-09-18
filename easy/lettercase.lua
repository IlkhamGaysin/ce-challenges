for line in io.lines(arg[1]) do
  local l = select(2, line:gsub("[a-z]", "."))
  local u = select(2, line:gsub("[A-Z]", "."))
  print(string.format("lowercase: %.2f uppercase: %.2f", 100*l/(l+u), 100*u/(l+u)))
end
