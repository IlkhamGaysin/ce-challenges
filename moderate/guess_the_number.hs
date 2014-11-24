import System.Environment (getArgs)

guess         :: Int -> Int -> [String] -> String
guess lo hi xs | hi < 0        = guess lo (read x) ys
               | x == "Lower"  = guess lo (d - 1) ys
               | x == "Higher" = guess (d + 1) hi ys
               | otherwise     = show d
               where x  = head xs
                     ys = tail xs
                     c  = mod (lo+hi) 2
                     d  = div (lo+hi) 2 + c

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (guess 0 (-1) . words) $ lines input
