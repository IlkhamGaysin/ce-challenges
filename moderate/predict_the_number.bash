while read line || [ -n "$line" ]; do
    i=0
    while [ $line -gt 0 ]; do
        i=$(($i + ($line & 1)))
        line=$(($line >> 1))
    done
    echo $(($i % 3))
done <$1
