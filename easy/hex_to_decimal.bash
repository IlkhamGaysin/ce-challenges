#!/bin/bash
while read line; do
    echo $((0X$line))
done < $1
