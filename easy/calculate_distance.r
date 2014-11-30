caldis <- function(s) {
  x <- as.integer(s[1]) - as.integer(s[3])
  y <- as.integer(s[2]) - as.integer(s[4])
  as.integer(sqrt(x*x + y*y))
}

ns <- strsplit(" -0123456789", "")[[1]]
numf <- function(s) {
  strsplit(paste(Filter(function(c) c %in% ns, strsplit(s, "")[[1]]), collapse=""), " ")[[1]]
}

cat(sapply(lapply(readLines(tail(commandArgs(), n=1)), numf), caldis), sep="\n")
