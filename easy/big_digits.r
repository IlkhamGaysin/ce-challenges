digits <- c("-**----*--***--***---*---****--**--****--**---**--",
            "*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-",
            "*--*---*---**---**--****-***--***----*---**---***-",
            "*--*---*--*-------*----*----*-*--*--*---*--*----*-",
            "-**---***-****-***-----*-***---**---*----**---**--",
            "--------------------------------------------------")
ns <- strsplit("0123456789", "")[[1]]
numf <- function(s) {
  strsplit(paste(Filter(function(c) c %in% ns, strsplit(s, "")[[1]]), collapse=""), " ")[[1]]
}

cat(sapply(lapply(readLines(tail(commandArgs(), n=1)), numf), function(s) {
  r <- c("", "", "", "", "", "")
  for (i in strsplit(s, "")[[1]]) {
    for (j in 1:6) {
      p <- as.integer(i)*5 + 1
      r[j] <- paste(r[j], substr(digits[j], p, p+4), sep="")
    }
  }
  paste(r, collapse="\n")
}), sep="\n")
