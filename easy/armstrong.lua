for line in io.lines(arg[1]) do
  local a = tonumber(line)
  local e = math.floor(math.log10(a)+1)
  local r = a
  while a > 0 do
    r = r - math.pow(math.fmod(a, 10), e)
    a = math.floor(a/10)
  end
  if r == 0 then
    print("True")
  else
    print("False")
  end
end
