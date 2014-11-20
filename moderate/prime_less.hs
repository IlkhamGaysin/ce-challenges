import System.Environment (getArgs)
import Data.List (intercalate)

factors  :: Int -> [Int]
factors n = [x | x <- [1..n], mod n x == 0]

prime  :: Int -> Bool
prime x = factors x == [1, x]

primesTo       :: [Int] -> Int -> Int -> [Int]
primesTo xs i n | i >= n        = xs
                | i == 2        = primesTo [2] 3 n
                | not (prime i) = primesTo xs (i + 2) n
                | otherwise     = primesTo (xs ++ [i]) (i + 2) n

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (intercalate "," . map show . (primesTo [] 2) . read) $ lines input
