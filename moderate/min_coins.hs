import System.Environment (getArgs)

coin   :: Integer -> String
coin i = show (fromInteger (div i 5) + b !! fromInteger (mod i 5))
       where b = [0, 1, 2, 1, 2]

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map coin . map (read :: String -> Integer) $ lines input
