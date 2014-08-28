#!/bin/bash
while read -r line; do
    echo $line|tr "[:lower:][:upper:]" "[:upper:][:lower:]"
done < $1
