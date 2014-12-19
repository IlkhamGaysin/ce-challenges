import System.Environment (getArgs)

stripChars :: String -> String -> String
stripChars = filter . flip notElem

distance   :: [Double] -> String
distance xs = show . floor $ sqrt (x*x + y*y)
            where x = head xs - xs!!2
                  y = xs!!1 - last xs

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (distance . map read . words) . lines $ stripChars "()," input
