import System.Environment (getArgs)
import Data.List (intercalate)

tele     :: String -> String -> [String]
tele t s | length s == 0 = [t]
         | h == '2'      = tele (t ++ "a") (tail s) ++ tele (t ++ "b") (tail s) ++ tele (t ++ "c") (tail s)
         | h == '3'      = tele (t ++ "d") (tail s) ++ tele (t ++ "e") (tail s) ++ tele (t ++ "f") (tail s)
         | h == '4'      = tele (t ++ "g") (tail s) ++ tele (t ++ "h") (tail s) ++ tele (t ++ "i") (tail s)
         | h == '5'      = tele (t ++ "j") (tail s) ++ tele (t ++ "k") (tail s) ++ tele (t ++ "l") (tail s)
         | h == '6'      = tele (t ++ "m") (tail s) ++ tele (t ++ "n") (tail s) ++ tele (t ++ "o") (tail s)
         | h == '7'      = tele (t ++ "p") (tail s) ++ tele (t ++ "q") (tail s) ++ tele (t ++ "r") (tail s) ++ tele (t ++ "s") (tail s)
         | h == '8'      = tele (t ++ "t") (tail s) ++ tele (t ++ "u") (tail s) ++ tele (t ++ "v") (tail s)
         | h == '9'      = tele (t ++ "w") (tail s) ++ tele (t ++ "x") (tail s) ++ tele (t ++ "y") (tail s) ++ tele (t ++ "z") (tail s)
         | otherwise     = tele (t ++ [h]) (tail s)
         where h = head s

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (intercalate ",") . map (tele "") $ lines input
