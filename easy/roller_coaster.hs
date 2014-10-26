import System.Environment (getArgs)
import Data.Char (isAlpha, toLower, toUpper)

rollerCoaster         :: Bool -> String -> String -> String
rollerCoaster up xs ys | null ys                 = xs
                       | isAlpha (head ys) && up = rollerCoaster False (xs ++ [toUpper (head ys)]) (tail ys)
                       | isAlpha (head ys)       = rollerCoaster True (xs ++ [toLower (head ys)]) (tail ys)
                       | otherwise               = rollerCoaster up (xs ++ [head ys]) (tail ys)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (rollerCoaster True "") $ lines input
