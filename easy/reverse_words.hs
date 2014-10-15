import System.Environment (getArgs)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map unwords . map reverse . map words $ lines input
