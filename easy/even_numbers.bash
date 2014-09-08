#!/bin/bash
while read line; do
    if [ $(($line % 2)) -ne 0 ]; then
        echo 0
    else
        echo 1
    fi
done < $1
