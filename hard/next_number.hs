import System.Environment (getArgs)

dsig    :: Int -> Int -> Int
dsig x y | y == 0       = x
         | mod y 10 > 0 = dsig (x + 2^(3 * mod y 10)) (div y 10)
         | otherwise    = dsig x (div y 10)

nextNu    :: Int -> Int -> Int
nextNu x y | x == dsig 0 y = y
           | otherwise     = nextNu x (y + 9)

nextNum  :: Int -> Int
nextNum x = nextNu (dsig 0 x) (x + 9)

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . nextNum . read) $ lines input
