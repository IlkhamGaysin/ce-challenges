for line in io.lines(arg[1]) do
    local a = tonumber(line)
    local b = math.floor(a)
    io.write(b .. ".")
    a = (a - b) * 60
    b = math.floor(a)
    if b <= 9 then
        io.write("0")
    end
    io.write(b .. "'")
    a = (a - b) * 60
    b = math.floor(a)
    if b <= 9 then
        io.write("0")
    end
    print(b .. "\"")
end
