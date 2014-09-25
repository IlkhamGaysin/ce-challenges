function log2(a)
  return math.floor(math.log(a)*1.44269504088896341 + 0.5)
end

for line in io.lines(arg[1]) do
  local a, b = line:match("(%d+),(%d+)")
  a, b = tonumber(a), tonumber(b)
  c = a - math.floor(a * 2^(-log2(b))) * 2^log2(b)
  if c > 0 then
    print(string.format("%.0f", a - c + b))
  else
    print(string.format("%.0f", a))
  end
end
