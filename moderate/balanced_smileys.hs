import System.Environment (getArgs)

isBalanced     :: Int -> String -> Bool
isBalanced c xs | null xs && c == 0                 = True
                | null xs || c < 0                  = False
                | not (elem x ":()")                = isBalanced c (tail xs)
                | not (elem y "()")                 = isBalanced c (init xs)
                | x == '('                          = isBalanced (succ c) (tail xs)
                | x == ')'                          = isBalanced (pred c) (tail xs)
                | x == ':' && head (tail xs) == '(' = isBalanced c (tail (tail xs)) || isBalanced (succ c) (tail (tail xs))
                | x == ':' && head (tail xs) == ')' = isBalanced c (tail (tail xs)) || isBalanced (pred c) (tail (tail xs))
                | x == ':'                          = isBalanced c (tail xs)
                | otherwise                         = False
                where x = head xs
                      y = last xs

balance   :: String -> String
balance xs | isBalanced 0 xs = "YES"
           | otherwise       = "NO"

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map balance $ lines input
