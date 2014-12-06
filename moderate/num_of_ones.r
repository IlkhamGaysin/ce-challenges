noo <- function(s) {
  r <- 0
  while (s > 0) {
    o <- s %% 2
    r <- r + o
    s <- (s - o) / 2
  }
  r
}

cat(sapply(as.integer(readLines(tail(commandArgs(), n=1))), noo), sep="\n")
