import Text.Printf (printf)

table       :: [[String]] -> [String] -> [[String]]
table xs ys | null ys   = xs
            | otherwise = table (xs ++ [take 12 ys]) (drop 12 ys)

main = do
    putStr . unlines . map concat . table [] $ map (printf "%4d" :: Int -> String) [x * y | x <- [1 .. 12], y <- [1 .. 12]]
