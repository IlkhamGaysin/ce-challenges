import System.Environment (getArgs)
import Data.List.Split (splitOn)

numzero       :: Int -> Int -> [Int] -> Int
numzero c z s | c == 0 && z == 0 = 1
              | c == 0           = 0
              | length s < c     = 0
              | otherwise        = numzero (c - 1) (z + head s) (tail s) + numzero c z (tail s)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map (numzero 4 0) . map (map read) . map (splitOn ",") $ lines input
