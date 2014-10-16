import System.Environment (getArgs)

isEven   :: Integer -> Integer
isEven i = mod ((+) i 1) 2

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map isEven . map read $ lines input
