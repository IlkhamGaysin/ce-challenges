happy <- function(s) {
  toString(Reduce("+", lapply(strsplit(s, "")[[1]], function(a) as.integer(a)*as.integer(a))))
}

ishappy <- function(a) {
  b <- c()
  for (i in 1:127) {
    if (a == "1" || a == "0") {
      return(a)
    }
    a <- happy(a)
    if (a %in% b) {
      return("0")
    }
    b <- append(b, a)
  }
  if (a == "1") {
    "1"
  } else {
    "0"
  }
}

cat(sapply(readLines(tail(commandArgs(), n=1)), ishappy), sep="\n")
