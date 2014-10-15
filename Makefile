clean:
	find {easy,moderate,hard} -name "*.hi" -o -name "*.o" -o -type f ! -name "*.*"|xargs rm -f
