import System.Environment (getArgs)

main = do
    let fibs = 0 : 1 : zipWith (+) fibs (tail fibs)
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map (fibs!!) . map read $ lines input
