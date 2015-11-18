Point = Struct.new(:x, :y)

def in_line(a, b, c)
  return (a.x - b.x) * (a.y - c.y) == (a.x - c.x) * (a.y - b.y)
end

def line(ax, bx, t)
  t.each_with_index do |i, ix|
    return ix >= bx if ix != ax and ix != bx and in_line(i, t[ax], t[bx])
  end
  false
end

File.open(ARGV[0]).each_line do |line|
  s = line.split(" | ").map { |i|
    t = i.split().map(&:to_i)
    Point.new(t[0], t[1])
  }
  r = 0
  for ix in (0..s.length-2)
    for jx in (ix+1..s.length-1)
      r += 1 if line(ix, jx, s)
    end
  end
  puts r
end
