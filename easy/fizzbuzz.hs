import System.Environment (getArgs)

fizz        :: Int -> [String] -> [Int] -> [String]
fizz i xs ys | i > last ys                  = xs
             | mod i f == 0 && mod i b == 0 = fizz (succ i) (xs ++ ["FB"]) ys
             | mod i f == 0                 = fizz (succ i) (xs ++ ["F"]) ys
             | mod i b == 0                 = fizz (succ i) (xs ++ ["B"]) ys
             | otherwise                    = fizz (succ i) (xs ++ [show i]) ys
             where f = head ys
                   b = ys!!1

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (unwords . fizz 1 [] . map read . words) $ lines input
