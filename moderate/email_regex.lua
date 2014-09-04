email_regex1 = "^[%w.+-=]+@%w+[%w.]*%w*%.%w%w%w?%w?$"
email_regex2 = "^\"[%w@.+-=]+\"@%w+[%w.]*%w*%.%w%w%w?%w?$"

for line in io.lines(arg[1]) do
    f = string.find(line, email_regex1) or string.find(line, email_regex2)
    if f then print("true") else print("false") end
end
