#!/bin/bash
let s=0
while read line; do
    let "s += $line"
done < $1
echo $s
