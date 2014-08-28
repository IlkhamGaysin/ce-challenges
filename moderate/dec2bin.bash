#!/bin/bash
while read line; do
    echo "obase=2;$line"|bc
done < $1
