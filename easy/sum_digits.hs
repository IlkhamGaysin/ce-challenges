import System.Environment (getArgs)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map sum . map (map read) . map (map (:[])) $ lines input
