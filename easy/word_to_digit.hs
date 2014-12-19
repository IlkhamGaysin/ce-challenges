import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.Maybe (fromJust)

digits :: [(String, Int)]
digits = [("one", 1), ("two", 2), ("three", 3), ("four", 4), ("five", 5), ("six", 6), ("seven", 7), ("eight", 8), ("nine", 9), ("zero", 0)]

doLookup    :: [(String, Int)] -> String -> String
doLookup d l = show . fromJust $ lookup l d

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (concatMap (doLookup digits) . splitOn ";") $ lines input
