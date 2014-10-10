$s = ["", "One", "Two", "Three", "Four",
     "Five", "Six", "Seven", "Eight", "Nine",
     "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
     "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
     "Twenty", "Thirty", "Forty", "Fifty",
     "Sixty", "Seventy", "Eighty", "Ninety",
     "Hundred", "Thousand", "Million"]

def wrd(a, b, c)
  return false if a + b + c == 0
  print $s[a] + $s[28] if a > 0
  if b > 1
    print $s[18 + b] + $s[c]
  else
    print $s[10*b + c]
  end
  true
end

File.open(ARGV[0]).each_line do |line|
  if line.to_i == 0
    print "Zero"
  else
    l = line.chomp.split("").map(&:to_i)
    while l.length < 9 do l.insert(0, 0) end
    print $s[30] if wrd(l[0], l[1], l[2])
    print $s[29] if wrd(l[3], l[4], l[5])
    wrd(l[6], l[7], l[8])
  end
  puts "Dollars"
end
