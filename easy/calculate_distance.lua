for line in io.lines(arg[1]) do
    _, _, x1, y1, x2, y2 = string.find(line, "%((-?%d+), (-?%d+)%) %((-?%d+), (-?%d+)%)")
    local x = tonumber(x1)-tonumber(x2)
    local y = tonumber(y1)-tonumber(y2)
    print(math.floor(math.sqrt(x*x+y*y)))
end
