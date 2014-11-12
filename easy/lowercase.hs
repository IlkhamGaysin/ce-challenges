import System.Environment (getArgs)
import Data.Char (toLower)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr $ map toLower input
