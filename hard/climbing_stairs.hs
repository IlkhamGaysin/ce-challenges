import System.Environment (getArgs)

main = do
    let stairs = 1 : 1 : zipWith (+) stairs (tail stairs)
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . (stairs!!) . read) $ lines input
