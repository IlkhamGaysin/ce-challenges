import System.Environment (getArgs)
import Data.Char (isDigit, isSpace)
import Data.List (intercalate, sort)

roadTrip      :: Int -> [Int] -> [Int] -> [Int]
roadTrip w x y | null y    = x
               | null x    = roadTrip (head y) [head y] (tail y)
               | otherwise = roadTrip (head y) (x ++ [head y - w]) (tail y)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (intercalate "," . map show . roadTrip 0 []. sort . map read . words) . lines $ [x | x <- input, isDigit x || isSpace x]
