# Formatting:
#   C:		clang-format -style='{ColumnLimit: 0, IndentWidth: 8, TabWidth: 8, UseTab: "Always"}'	is a good start
#   Go:		go fmt
#   Perl:	perltidy
#   Ruby:	rubocop	is a good start

clean:
	find {easy,moderate,hard} -name "*.hi" -o -name "*.o" -o -name "*.pl.tdy" -o -name "Main*" -o -name "test.*.txt" -o -type f ! -name "*.*"|xargs rm -f
