for line in io.lines(arg[1]) do
    line = line:gsub("^%l", string.upper):gsub("%s%l", string.upper)
    print(line)
end
