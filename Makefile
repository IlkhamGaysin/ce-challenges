clean:
	find {easy,moderate,hard} -name "*.hi" -o -name "*.o" -o -name "Main*" -o -name "test*.txt" -o -type f ! -name "*.*"|xargs rm -f
