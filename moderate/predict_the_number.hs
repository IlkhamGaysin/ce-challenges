import System.Environment (getArgs)
import Data.Bits (popCount)

count   :: Int -> String
count i = show $ mod (popCount i) 3

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map count . map (read :: String -> Int) $ lines input
