import fileinput

b = [0, 1, 2, 1, 2]
for line in fileinput.input():
    i = int(line)
    print(i // 5 + b[i % 5])
