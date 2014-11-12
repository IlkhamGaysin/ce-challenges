import System.Environment (getArgs)
import Data.Char (isLower, isUpper)
import Text.Printf (printf)

ratios  :: String -> String
ratios s = printf "lowercase: %.2f uppercase: %.2f" (100 * l / t) (100 * u / t)
         where l = fromIntegral (length [x | x <- s, isLower x]) :: Double
               u = fromIntegral (length [x | x <- s, isUpper x]) :: Double
               t = l + u

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map ratios $ lines input
