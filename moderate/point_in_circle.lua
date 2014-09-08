for line in io.lines(arg[1]) do
  _, _, x1, y1, r, x2, y2 = string.find(line, "Center: %((-?%d+.?%d*), (-?%d+.?%d*)%); Radius: (-?%d+.?%d*); Point: %((-?%d+.?%d*), (-?%d+.?%d*)%)")
  local x = x2-x1
  local y = y2-y1
  if x*x + y*y <= r*r then
    print("true")
  else
    print("false")
  end
end
