import System.Environment (getArgs)
import Data.List (isSuffixOf)
import Data.List.Split (splitOn)

nonempty   :: [[String]] -> [[String]]
nonempty xs = [x | x <- xs, length x == 2]

isSuff   :: [String] -> String
isSuff xs | isSuffixOf (last xs) (head xs) = "1"
          | otherwise                      = "0"

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map isSuff . nonempty . map (splitOn ",") $ lines input
