import System.Environment (getArgs)

compress      :: [Int] -> [Int] -> [Int]
compress xs ys | null ys                       = xs
               | null xs || last xs /= head ys = compress (xs ++ [1, head ys]) (tail ys)
               | otherwise                     = compress (init (init xs) ++ [succ . last $ init xs] ++ [last xs]) (tail ys)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (unwords . map show . compress [] . map read . words) $ lines input
