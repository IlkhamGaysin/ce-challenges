#!/bin/bash
tr "," " " <$1 | while read line; do
    a=( $line )
    printf ${a[0]}
    for ((i=1; i<${#a[*]}; i++)); do
        if [ ${a[$i-1]} -ne ${a[$i]} ]; then
            printf ",%d" ${a[$i]}
        fi
    done
    echo
done
