import System.Environment (getArgs)
import Data.List (intercalate, nub)
import Data.List.Split (splitOn)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (intercalate ",") . map nub . map (splitOn ",") $ lines input
