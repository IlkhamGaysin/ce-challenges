import System.Environment (getArgs)
import Data.Char (isDigit)

toDigits  :: Integer -> [Integer]
toDigits x | x < 0     = [-1]
           | x < 10    = [x]
           | otherwise = (toDigits (div x 10)) ++ [mod x 10]

toDigitsRev  :: Integer -> [Integer]
toDigitsRev x | x < 0  = [-1]
              | otherwise = reverse (toDigits x)

dobSec :: [Integer] -> [Integer] -> [Integer]
dobSec xs [] = xs
dobSec xs [x] = xs ++ [x]
dobSec xs ys = dobSec (xs ++ [y, 2*z]) (tail (tail ys))
             where y = head ys
                   z = head (tail ys)

doubleSecond :: [Integer] -> [Integer]
doubleSecond xs = dobSec [] xs

sumDigits :: [Integer] -> Integer
sumDigits [] = 0
sumDigits (x:xs) = (div x 10) + (mod x 10) + sumDigits xs

isValid  :: Integer -> String
isValid x | mod (sumDigits (doubleSecond (toDigitsRev x))) 10 == 0 = "1"
          | otherwise                                              = "0"

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (isValid . read . filter isDigit) $ lines input
