def push(stack, int)
  stack << int
end

def pop(stack)
  int = stack[-1]
  stack.delete_at(-1)
  int
end

File.open(ARGV[0]).each_line do |line|
  stack = Array.new
  line.split.map(&:to_i).each do |i|
    push(stack, i)
  end
  out = Array.new
  while stack.length > 0 do
    out << pop(stack)
    pop(stack)
  end
  puts out.join(" ")
end
