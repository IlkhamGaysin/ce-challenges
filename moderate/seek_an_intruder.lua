function o2d(a)
  return 64*tonumber(a:sub(1, 1)) + 8*tonumber(a:sub(2,2)) + tonumber(a:sub(3,3))
end

function b2d(a)
  local ret = 0
  for i = 1, 8 do
    ret = ret + tonumber(a:sub(-i, -i)) * 2^(i-1)
  end
  return ret
end

ips = {}
for line in io.lines(arg[1]) do
  for i in line:gmatch("[1-2]?%d?%d%.[1-2]?%d?%d%.[1-2]?%d?%d%.[1-2]?%d?%d") do
    _, _, a1, a2, a3, a4 = i:find("(%d+)%.(%d+)%.(%d+)%.(%d+)")
    local b1, b2, b3, b4 = tonumber(a1), tonumber(a2), tonumber(a3), tonumber(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("0x%x?%x%.0x%x?%x%.0x%x?%x%.0x%x?%x") do
    _, _, a1, a2, a3, a4 = i:find("(0x%x%x?)%.(0x%x%x?)%.(0x%x%x?)%.(0x%x%x?)")
    local b1, b2, b3, b4 = tonumber(a1), tonumber(a2), tonumber(a3), tonumber(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("0[0-3][0-7][0-7]%.0[0-3][0-7][0-7]%.0[0-3][0-7][0-7]%.0[0-3][0-7][0-7]") do
    _, _, a1, a2, a3, a4 = i:find("0([0-3][0-7][0-7])%.0([0-3][0-7][0-7])%.0([0-3][0-7][0-7])%.0([0-3][0-7][0-7])")
    local b1, b2, b3, b4 = o2d(a1), o2d(a2), o2d(a3), o2d(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("[01][01][01][01][01][01][01][01]%.[01][01][01][01][01][01][01][01]%.[01][01][01][01][01][01][01][01]%.[01][01][01][01][01][01][01][01]") do
    _, _, a1, a2, a3, a4 = i:find("([01][01][01][01][01][01][01][01])%.([01][01][01][01][01][01][01][01])%.([01][01][01][01][01][01][01][01])%.([01][01][01][01][01][01][01][01])")
    local b1, b2, b3, b4 = b2d(a1), b2d(a2), b2d(a3), b2d(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("[01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01][01]") do
    _, _, a1, a2, a3, a4 = i:find("([01][01][01][01][01][01][01][01])([01][01][01][01][01][01][01][01])([01][01][01][01][01][01][01][01])([01][01][01][01][01][01][01][01])")
    local b1, b2, b3, b4 = b2d(a1), b2d(a2), b2d(a3), b2d(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("0[0-3][0-7][0-7][0-3][0-7][0-7][0-3][0-7][0-7][0-3][0-7][0-7]") do
    _, _, a1, a2, a3, a4 = i:find("0([0-3][0-7][0-7])([0-3][0-7][0-7])([0-3][0-7][0-7])([0-3][0-7][0-7])")
    local b1, b2, b3, b4 = o2d(a1), o2d(a2), o2d(a3), o2d(a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("0x%x%x%x%x%x%x%x%x") do
    _, _, a1, a2, a3, a4 = i:find("(0x%x%x)(%x%x)(%x%x)(%x%x)")
    local b1, b2, b3, b4 = tonumber(a1), tonumber("0x" .. a2), tonumber("0x" .. a3), tonumber("0x" .. a4)
    if b1 > 0 and b1 < 256 and b2 < 256 and b3 < 256 and b4 < 255 then
      ips[#ips + 1] = b1 .. "." .. b2 .. "." .. b3 .. "." .. b4
    end
  end
  for i in line:gmatch("[1-4]?%d?%d?%d?%d?%d?%d?%d?%d?%d") do
    a1 = tonumber(i)
    if a1 >= 2^24 and a1 < 2^32-1 then
      ips[#ips + 1] = math.floor(a1/2^24)%256 .. "." .. math.floor(a1/2^16)%256 .. "." .. math.floor(a1/2^8)%256 .. "." .. a1%256
    end
  end
end

m, s = 0, ""
for i = 1, #ips-1 do
  local c = 0
  for j = i+1, #ips do
    if ips[i] == ips[j] then c = c + 1 end
  end
  if c > m then
    m, s = c, ips[i]
  end
end
print(s)
