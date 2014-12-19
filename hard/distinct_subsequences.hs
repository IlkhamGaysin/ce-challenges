import System.Environment (getArgs)
import Data.List.Split (splitOn)

subs   :: [String] -> Int
subs s | null (last s)                  = 1
       | null (head s)                  = 0
       | head (head s) == head (last s) = subs [tail (head s), tail (last s)] + subs [tail (head s), last s]
       | otherwise                      = subs [tail (head s), last s]

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . subs . splitOn ",") $ lines input
