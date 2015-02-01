#!/bin/bash
tr "," " " <$1 | while read line || [ -n "$line" ]; do
    a=( $line )
    echo $((${a[0]}-( ${a[0]}/${a[1]} )*${a[1]}))
done
