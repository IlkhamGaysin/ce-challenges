#!/bin/bash
{ cat $1; echo; } | while read line; do
    u=1
    for ((i=0; i<${#line}; i++)); do
        a="${line:$i:1}"
        if [ -z "${a//[a-zA-Z]}" ]; then
            if [ $u -eq 1 ]; then
                printf ${a^^}
                u=0
            else
                printf ${a,,}
                u=1
            fi
        else
            printf "$a"
        fi
    done
    echo
done
