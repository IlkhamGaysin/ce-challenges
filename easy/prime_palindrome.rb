require 'prime'
puts 999.downto(3).find { |i| i.to_s == i.to_s.reverse && Prime.prime?(i) }
