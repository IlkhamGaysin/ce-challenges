phrases = [", yeah!",
	   ", this is crazy, I tell ya.",
	   ", can U believe this?",
	   ", eh?",
	   ", aw yea.",
	   ", yo.",
	   "? No way!",
	   ". Awesome!"]
i, l = 0, false
File.open(ARGV[0]).each_line do |line|
  s = line.split("").each { |x|
    if x == "." || x == "!" || x == "?"
      if l
        print phrases[i]
        i, l = (i+1) % 8, false
      else
        print x
        l = true
      end
    else
      print x
    end
  }
  puts if s.last != "\n"
end
