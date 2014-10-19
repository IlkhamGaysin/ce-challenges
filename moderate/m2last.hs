import System.Environment (getArgs)

result   :: [String] -> [String]
result s = [x | x <- s, x /= ""]

m2last   :: [String] -> String
m2last s | l > length s = ""
         | otherwise    = s !! ((length s) - l)
         where l = succ . read $ last s

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . result . map m2last . map words $ lines input
