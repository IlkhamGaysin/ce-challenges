#!/bin/bash
while read line; do
    a=( $line )
    i=1
    while [ $i -le ${a[2]} ]; do
        if [ $i -gt 1 ]; then
            printf " "
        fi
        if [ $(($i % ${a[0]})) -gt 0 ] && [ $(($i % ${a[1]})) -gt 0 ]; then
            printf $i
        else
            if [ $(($i % ${a[0]})) -eq 0 ]; then
                printf "F"
            fi
            if [ $(($i % ${a[1]})) -eq 0 ]; then
                printf "B"
            fi
        fi
        ((i++))
    done
    echo
done < $1
