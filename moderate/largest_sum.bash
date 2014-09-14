#!/bin/bash
sed -e "s/,[:space:]*/ /g" $1 | while read line; do
    a=( $line )
    l=0
    m=${a[0]}
    for i in ${a[*]}; do
        if [ $i -gt $m ]; then
            m=$i
        fi
        if [ $(($i+$l)) -gt $m ]; then
            m=$(($i+$l))
        fi
        if [ $(($i+$l)) -gt 0 ]; then
            l=$(($i+$l))
        else
            l=0
        fi
    done
    echo $m
done
