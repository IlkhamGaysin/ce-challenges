function div(a, b) return math.floor(a/b) end

s = {[0]="", "One", "Two", "Three", "Four",
     "Five", "Six", "Seven", "Eight", "Nine",
     "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
     "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
     "Twenty", "Thirty", "Forty", "Fifty",
     "Sixty", "Seventy", "Eighty", "Ninety",
     "Hundred", "Thousand", "Million"}

function wrd(a1, a2, a3)
  if a1 + a2 + a3 == 0 then return false end
  if a1 > 0 then
    io.write(s[a1] .. s[28])
  end
  if a2 > 1 then
    io.write(s[18 + a2] .. s[a3])
  else
    io.write(s[10*a2 + a3])
  end
  return true
end

for line in io.lines(arg[1]) do
  local b, c = tonumber(line), {}
  if b == 0 then
    io.write("Zero")
  else
    for i = 9, 1, -1 do
      c[i], b = b%10, div(b, 10)
    end
    if wrd(c[1], c[2], c[3]) then io.write(s[30]) end
    if wrd(c[4], c[5], c[6]) then io.write(s[29]) end
    _ = wrd(c[7], c[8], c[9])
  end
  io.write("Dollars\n")
end
