#!/bin/bash
while read line; do
    a=( $line )
    l=$((${#a[@]}-1))
    printf "${a[$l]}"
    for ((i=$l-1; i>=0; i--)); do
        printf " ${a[$i]}"
    done
    echo
done < $1
