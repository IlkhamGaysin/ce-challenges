import System.Environment (getArgs)
import Data.Char (toUpper)

capitalize   :: String -> String
capitalize s = toUpper (head s) : tail s

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map unwords . map (map capitalize) . map words $ lines input
