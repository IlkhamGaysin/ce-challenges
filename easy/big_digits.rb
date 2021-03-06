digits = ['-**----*--***--***---*---****--**--****--**---**--',
          '*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-',
          '*--*---*---**---**--****-***--***----*---**---***-',
          '*--*---*--*-------*----*----*-*--*--*---*--*----*-',
          '-**---***-****-***-----*-***---**---*----**---**--',
          '--------------------------------------------------']
w = 5
h = 6
File.open(ARGV[0]).each_line do |line|
  (0...h).each do |i|
    line.gsub(/\D/, '').each_char do |j|
      print digits[i][j.to_i * w, w]
    end
    puts
  end
end
