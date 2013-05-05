from sys import argv

test_cases = open(argv[1], 'r')
for test in test_cases:
    st = test.split()
    print ' '.join(i for i in reversed(st))
test_cases.close()
