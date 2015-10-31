File.open(ARGV[0]).each_line do |line|
  s = line.scan(/Vampires: (\d+), Zombies: (\d+), Witches: (\d+), Houses: (\d+)/)[0].map(&:to_i)
  puts (s[0] * 3 + s[1] * 4 + s[2] * 5) * s[3] / (s[0] + s[1] + s[2])
end
