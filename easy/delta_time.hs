import System.Environment (getArgs)
import Data.List.Split (splitOn)
import Data.List (intercalate)

fmt0  :: Integer -> String
fmt0 i | i < 10    = "0" ++ (show i)
       | otherwise = show i

nice  :: Integer -> String
nice i = intercalate ":" $ map fmt0 t
       where t = (div i 3600) : (mod (div i 60) 60) : [mod i 60]

delta  :: [[Integer]] -> Integer
delta s = abs (h1*3600 + m1*60 + s1 - h2*3600 - m2*60 - s2)
        where h1 = head (head s)
              m1 = head s !! 1
              s1 = last (head s)
              h2 = head (last s)
              m2 = last s !! 1
              s2 = last (last s)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (nice . delta . map (map read . splitOn ":") . words) $ lines input
