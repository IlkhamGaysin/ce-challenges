for line in io.lines(arg[1]) do
    local c = line:sub(-1)
    local l = line:find(c .. "[^" .. c .. "]*,")
    if not l then l = 0 end
    print(l - 1)
end
