import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.List (elemIndices)

lastIndex   :: [String] -> String
lastIndex s | null e    = "-1"
            | otherwise = show . last $ e
            where e = elemIndices (head (last s)) (head s)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map lastIndex . map (splitOn ",") $ lines input
