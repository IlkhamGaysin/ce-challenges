File.open(ARGV[0]).each_line do |line|
  valid, stack = true, Array.new
  line.chomp.each_char do |i|
    case i
    when ")"
      if stack.length > 0 && stack[-1] == "(" then
        stack.delete_at(-1)
      else
        valid = false
        break
      end
    when "]"
      if stack.length > 0 && stack[-1] == "[" then
        stack.delete_at(-1)
      else
        valid = false
        break
      end
    when "}"
      if stack.length > 0 && stack[-1] == "{" then
        stack.delete_at(-1)
      else
        valid = false
        break
      end
    else
      stack << i
    end
  end
  puts valid && stack.length == 0 ? "True" : "False"
end
