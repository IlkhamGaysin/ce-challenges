import System.Environment (getArgs)

main :: IO ()
main = do
    let fibs = 0 : 1 : zipWith (+) fibs (tail fibs)
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . (fibs!!) . read) $ lines input
