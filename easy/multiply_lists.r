multiply <- function(s) {
  a <- s[1:(length(s)/2)]
  b <- s[(length(s)/2+2):(length(s)+1)]
  paste(sapply(a, as.integer) * sapply(b, as.integer), collapse=" ")
}

cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), " "), multiply), sep="\n")
