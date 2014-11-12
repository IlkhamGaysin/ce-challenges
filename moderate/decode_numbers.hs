import System.Environment (getArgs)

decode  :: String -> Int
decode s | length s < 2                                     = 1
         | head s /= '1' && head s /= '2'                   = decode t
         | head t == '0' || (head s == '2' && head t > '6') = decode (tail t)
         | length s == 2                                    = 2
         | otherwise                                        = decode t + decode (tail t)
         where t = tail s

main :: IO ()
main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (show . decode) $ lines input
