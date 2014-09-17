#!/bin/bash
while read line; do
    a=( `echo $line | tr ' ' '\n' | sort -g` )
    echo ${a[*]}
done <$1
