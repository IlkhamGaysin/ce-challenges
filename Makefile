# Formatting:
#   C:	clang-format -style='{ColumnLimit: 0, IndentWidth: 8, TabWidth: 8, UseTab: "Always"}'	is a good start
#   Go:	go fmt

clean:
	find {easy,moderate,hard} -name "*.hi" -o -name "*.o" -o -name "Main*" -o -name "test.*.txt" -o -type f ! -name "*.*"|xargs rm -f
