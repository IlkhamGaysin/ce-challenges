import System.Environment (getArgs)
import Data.List (sortBy)

jolly   :: [Int] -> String
jolly xs | null xs              = "Jolly"
         | head xs /= length xs = "Not jolly"
         | otherwise            = jolly (tail xs)

delta          :: [Int] -> [Int] -> [Int]
delta _ []      = []
delta xs (y:ys) | null ys   = xs
                | otherwise = delta (abs (y - head ys) : xs) ys

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (jolly . sortBy (flip compare) . delta [] . map read . tail . words) $ lines input
