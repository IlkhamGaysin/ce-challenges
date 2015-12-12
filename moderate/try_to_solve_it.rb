File.open(ARGV[0]).each_line do |line|
  puts line.chars.map {|x|
    case x
    when 'a'..'f'
      (x.ord + 20).chr
    when 'g'..'m'
      (x.ord + 7).chr
    when 'n'..'t'
      (x.ord - 7).chr
    when 'u'..'z'
      (x.ord - 20).chr
    else
      x
    end
  }.join
end
