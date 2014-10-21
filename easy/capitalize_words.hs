import System.Environment (getArgs)
import Data.Char (toUpper)

capitalize          :: String -> String
capitalize (x : xs) = toUpper x : xs

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map unwords . map (map capitalize) . map words $ lines input
