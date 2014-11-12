import System.Environment (getArgs)

morse        :: Int -> String -> String -> String
morse i xs ys | null ys        = xs ++ [m !! (i - 2)]
              | head ys == '.' = morse (i * 2) xs (tail ys)
              | head ys == '-' = morse (i * 2 + 1) xs (tail ys)
              | i == 1         = morse i (xs ++ " ") (tail ys)
              | otherwise      = morse 1 (xs ++ [m !! (i - 2)]) (tail ys)
              where m = "ETIANMSURWDKGOHVF L PJBXCYZQ  54 3   2       16       7   8 90"

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (morse 1 "") $ lines input
