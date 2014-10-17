import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.Bits (testBit)

compareBits   :: [Int] -> String
compareBits s = if testBit i a == testBit i b
                  then "true"
                  else "false"
                where i = head s
                      a = (s !! 1) - 1
                      b = (last s) - 1

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map compareBits . map (map read) . map (splitOn ",") $ lines input
