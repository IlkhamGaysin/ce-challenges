import System.Environment (getArgs)

juggle     :: Integer -> [Int] -> Integer
juggle i xs | null xs      = i
            | head xs == 1 = juggle (i * 2^(head (tail xs))) (tail (tail xs))
            | otherwise    = juggle (((succ i) * 2^(head (tail xs))) - 1) (tail (tail xs))

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map (juggle 0) . map (map length) . map words $ lines input
