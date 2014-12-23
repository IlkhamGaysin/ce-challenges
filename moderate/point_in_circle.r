ns <- strsplit(" -0123456789.", "")[[1]]
numf <- function(s) {
  strsplit(paste(Filter(function(c) c %in% ns, strsplit(s, "")[[1]]), collapse=""), " ")[[1]]
}

cat(sapply(lapply(readLines(tail(commandArgs(), n=1)), numf), function(s) {
  t <- as.double(s)
  t <- t[!is.na(t)]
  x <- t[4]-t[1]
  y <- t[5]-t[2]
  if (x*x + y*y <= t[3]*t[3]) {
    "true"
  } else {
    "false"
  }
}), sep="\n")
