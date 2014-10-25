import System.Environment (getArgs)

roman    :: String -> Int -> Int -> String
roman s j i | i == 0       = s
            | ronum !! j > i = roman s (succ j) i
            | otherwise    = roman (s ++ rostr !! j) j (i - ronum !! j)
            where ronum = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
                  rostr = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (roman "" 0) . map read $ lines input
