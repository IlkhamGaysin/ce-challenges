import Data.Binary (decode)
import Data.Binary.Put (putWord16host, runPut)
import Data.Word (Word8)

littleEndian = (decode . runPut $ putWord16host 42 :: Word8) == 42

main | littleEndian = putStrLn "LittleEndian"
     | otherwise    = putStrLn "BigEndian"
