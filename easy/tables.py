from __future__ import print_function

for i in xrange(1, 13):
	for j in xrange(1, 13):
		print("{0:4}".format(i * j), sep = '', end = '')
	print()
