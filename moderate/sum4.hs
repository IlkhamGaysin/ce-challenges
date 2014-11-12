import System.Environment (getArgs)
import Data.List.Split (splitOn)

numzero      :: Int -> Int -> [Int] -> Int
numzero c z s | c == 0 && z == 0 = 1
              | c == 0           = 0
              | length s < c     = 0
              | otherwise        = numzero (c - 1) (z + head s) (tail s) + numzero c z (tail s)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . numzero 4 0 . map read . splitOn ",") $ lines input
