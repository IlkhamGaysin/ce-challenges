import System.Environment (getArgs)
import Data.List (sort)
import Numeric (showFFloat)

formatFloat3   :: Double -> String
formatFloat3 f = showFFloat (Just 3) f ""

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map unwords . map (map formatFloat3) . map sort . map (map (read :: String -> Double)) . map words $ lines input
