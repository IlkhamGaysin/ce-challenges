import System.Environment (getArgs)

juggle     :: Integer -> [Int] -> Integer
juggle i xs | null xs      = i
            | head xs == 1 = juggle (i * 2 ^ head (tail xs)) (tail (tail xs))
            | otherwise    = juggle ((succ i * 2 ^ head (tail xs)) - 1) (tail (tail xs))

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . juggle 0 . map length . words) $ lines input
