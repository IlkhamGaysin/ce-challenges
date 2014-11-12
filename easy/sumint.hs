import System.Environment (getArgs)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStrLn . show . sum . map read $ lines input
