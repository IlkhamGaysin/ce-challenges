import System.Environment (getArgs)
import Data.Char (toLower)

pang          :: String -> String -> String
pang []     _  = ""
pang (x:xs) ys | elem x ys = pang xs ys
               | otherwise = x : pang xs ys

pangram  :: String -> String
pangram s | null ls   = "NULL"
          | otherwise = ls
          where ls = pang ['a'..'z'] s

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map pangram $ lines [toLower x | x <- input, x == '\n' || elem x ['a'..'z'] || elem x ['A'..'Z']]
