import System.Environment (getArgs)
import Data.List (sortBy)

sortLength    :: [a] -> [b] -> Ordering
sortLength a b | length a > length b = LT
               | otherwise           = GT

merge     :: Int -> [String] -> [String]
merge i xs | length xs <= i = sortBy sortLength xs
           | otherwise      = take i (sortBy sortLength xs)

longest        :: Int -> [String] -> [String] -> [String]
longest i xs ys | null ys   = xs
                | i == 0    = longest (read $ head ys) [] (tail ys)
                | otherwise = longest i (merge i ((head ys) : xs)) (tail ys)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . longest 0 [] $ lines input
