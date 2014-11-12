import System.Environment (getArgs)
import Data.List.Split (splitOn)

remove  :: [String] -> String
remove s = [x | x <- head s, not . elem x $ last s]

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (remove . splitOn ", ") $ lines input
