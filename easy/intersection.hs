import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.List (intercalate, intersect)

doIntersect   :: Eq a => [[a]] -> [a]
doIntersect l = intersect (head l) (last l)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (intercalate ",") . map doIntersect . map (map (splitOn ",")) . map (splitOn ";") $ lines input
