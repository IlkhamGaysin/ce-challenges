capiword <- function(s) {
  paste(toupper(substr(s, 1, 1)), substr(s, 2, nchar(s)), sep="")
}

capita <- function(s) {
  paste(sapply(s, capiword), collapse=" ")
}

cat(sapply(strsplit(readLines(tail(commandArgs(), n=1)), " "), capita), sep="\n")
