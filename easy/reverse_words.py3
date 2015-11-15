import fileinput

for line in fileinput.input():
    if not line == '\n':
        print(*[i for i in reversed(line.split())], sep=' ')
