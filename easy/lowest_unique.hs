import System.Environment (getArgs)
import Data.Bits (clearBit, setBit, testBit)
import Data.List (elemIndex)
import Data.Maybe (fromJust)

lowes  :: Int -> Int
lowes i | testBit i 1 = 1
        | otherwise   = succ (lowes (div i 2))

lowest          :: Int -> Int -> [[Int]] -> Int
lowest i j [ws, xs] | null xs && i == 0 = 0
                    | null xs           = succ . fromJust $ elemIndex (lowes i) ws
                    | testBit j x = lowest i j [ws, ys]
                    | testBit i x = lowest (clearBit i x) (setBit j x) [ws, ys]
                    | otherwise   = lowest (setBit i x) j [ws, ys]
                    where x = head xs
                          ys = tail xs

twice   :: [Int] -> [[Int]]
twice xs = [xs, xs]

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . (lowest 0 0 . twice) . (map read) . words) $ lines input
