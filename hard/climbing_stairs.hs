import System.Environment (getArgs)

main = do
    let stairs = 1 : 1 : zipWith (+) stairs (tail stairs)
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map (stairs!!) . map read $ lines input
