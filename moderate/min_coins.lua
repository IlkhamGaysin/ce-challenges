b = {[0]=0, 1, 2, 1, 2}
for line in io.lines(arg[1]) do
  print(math.floor(line/5) + b[line%5])
end
