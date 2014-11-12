import System.Environment (getArgs)
import Data.List.Split (splitOn)

modulo  :: [Integer] -> Integer
modulo a = h - div h l * l
       where h = head a
             l = last a

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . modulo . map read . splitOn ",") $ lines input
