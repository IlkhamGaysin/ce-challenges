class String
  def is_upper?
    self == self.upcase
  end
  def is_lower?
    self == self.downcase
  end
end

File.open(ARGV[0]).each_line do |line|
  s = line.split("").each {|x| print (x.is_lower? ? x.upcase : (x.is_upper? ? x.downcase : x))}
end
