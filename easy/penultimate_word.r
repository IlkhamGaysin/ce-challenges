penultimate <- function(xs) {
  if (length(xs) < 2) {
    return("")
  }
  xs[[length(xs)-1]]
}

cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), " "), penultimate), sep="\n")
