M = [50, 25, 10, 5, 1]

def alter(n, p)
  return 1 if n == 0
  while M[p] > n
    p += 1
  end
  return 1 if M[p] == 1
  alter(n - M[p], p) + alter(n, p + 1)
end

File.open(ARGV[0]).each_line do |line|
  puts alter(line.to_i, 0)
end