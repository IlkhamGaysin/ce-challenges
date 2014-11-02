import System.Environment (getArgs)
import Data.List (sort)

jolly   :: [Int] -> String
jolly xs | null xs              = "Jolly"
         | head xs /= length xs = "Not jolly"
         | otherwise            = jolly (tail xs)

delta      :: [Int] -> [Int] -> [Int]
delta xs ys | length ys == 1 = xs
            | otherwise      = delta ((abs (head ys - ys!!1)) : xs) (tail ys)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (jolly . reverse . sort . delta [] . map read . tail . words) $ lines input
