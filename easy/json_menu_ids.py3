import fileinput, re

json = re.compile('"id": (\d+),')
for line in fileinput.input():
    if line != "\n":
        print(sum(int(i) for i in json.findall(line)))
