for line in io.lines(arg[1]) do
    line = line:gsub("%u", string.lower)
    print(line)
end
