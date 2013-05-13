require 'prime'
puts Prime.each(1000).find_all{|x| x.to_s == x.to_s.reverse}.last
