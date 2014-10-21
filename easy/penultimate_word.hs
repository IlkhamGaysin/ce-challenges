import System.Environment (getArgs)

penulti   :: [String] -> String
penulti s | length s < 2 = ""
          | otherwise    = s !! (length s - 2)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map penulti . map words $ lines input
