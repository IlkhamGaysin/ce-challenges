caldis <- function(s) {
  x <- s[1] - s[3]
  y <- s[2] - s[4]
  as.integer(sqrt(x*x + y*y))
}

numf <- function(s) {
  sapply(strsplit(gsub("[^- 0123456789]", "", s), " ")[[1]], as.integer)
}

cat(sapply(lapply(readLines(tail(commandArgs(), n=1)), numf), caldis), sep="\n")
