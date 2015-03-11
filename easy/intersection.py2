import fileinput

for line in fileinput.input():
    st = line.strip('\n').split(';')
    inter = list(set(st[0].split(',')) & set(st[1].split(',')))
    print ','.join(sorted(inter))
