cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), ","), function(s) {
  a <- tail(gregexpr(s[2], s[1])[[1]], n=1)
  if (a == -1) {
    -1
  } else {
    a-1
  }
}), sep="\n")
