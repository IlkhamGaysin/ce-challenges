import System.Environment (getArgs)
import Numeric (readHex)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map show . map fst . map head . map readHex $ lines input
