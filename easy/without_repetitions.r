wr <- function(xs) {
  Reduce(function(x,y) {if (y == substr(x, nchar(x), nchar(x))) x else paste(x, y, sep="")}, strsplit(xs, "")[[1]], "")
}

cat(sapply(readLines(tail(commandArgs(), n=1)), wr), sep="\n")
