largest <- function(s) {
  m <- as.integer(s[1])
  l <- 0
  for (i in 1:length(s)) {
    a <- as.integer(s[i])
    if (a > m) {
      m <- a
    }
    if (a + l > m) {
      m <- a + l
    }
    if (a + l > 0) {
      l <- l + a
    } else {
      l <- 0
    }
  }
  m
}

cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), ","), largest), sep="\n")
