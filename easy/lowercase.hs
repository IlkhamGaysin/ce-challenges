import System.Environment (getArgs)
import Data.Char

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr $ map toLower input
