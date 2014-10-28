import System.Environment (getArgs)
import Data.List.Split (splitOn)

largest       :: Int -> Int -> [Int] -> String
largest i j xs | null xs     = show i
               | otherwise   = largest k l (tail xs)
               where k = maximum [head xs, (head xs) + j, i]
                     l = max ((head xs) + j) 0

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (largest (minBound::Int) 0) . map (map read) . map (splitOn ",") $ lines input
