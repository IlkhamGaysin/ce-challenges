cat(sapply(lapply(strsplit(readLines(tail(commandArgs(), n=1)), " "), rev), paste, collapse=" "), sep="\n")
