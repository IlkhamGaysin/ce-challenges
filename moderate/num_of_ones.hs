import System.Environment (getArgs)
import Data.Bits (popCount)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map popCount . map (read :: String -> Int) $ lines input
