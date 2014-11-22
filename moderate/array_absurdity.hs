import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.List (sort)

absurd   :: [String] -> String
absurd xs | length xs < 2 = ""
          | x == head ys  = x
          | otherwise     = absurd ys
          where x  = head xs
                ys = tail xs

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (absurd . sort . splitOn "," . last . splitOn ";") $ lines input
