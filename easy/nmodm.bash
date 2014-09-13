#!/bin/bash
sed -e "s/,/ /g" $1 | while read line; do
    a=( $line )
    echo $((${a[0]}-( ${a[0]}/${a[1]} )*${a[1]}))
done
