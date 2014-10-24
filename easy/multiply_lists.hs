import System.Environment (getArgs)
import Data.List.Split (splitOn)

multiply   :: [[Int]] -> [Int]
multiply xs = [x * y | (x, y) <- zip (head xs) (last xs)]

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map unwords . map (map show) . map multiply . map (map (map read)) . map (map words) . map (splitOn "|") $ lines input
