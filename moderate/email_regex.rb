email_regex = /((^([\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~]+[\.\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~])*[\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~]+)|(^([\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~]+\.)*("[@\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~]+"(\.[\w!#\$%&'\*\+-\/=\?\^_`\{\|\}~]+)*)+))@(\w+\.)+\w{2,4}$/

File.open(ARGV[0]).each_line do |line|
  puts line =~ email_regex ? 'true' : 'false'
end
