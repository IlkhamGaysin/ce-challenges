longest <- function(xs) {
  Reduce(function(x,y) {if (nchar(y) > nchar(x)) y else x}, xs, "")
}

cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), " "), longest), sep="\n")
