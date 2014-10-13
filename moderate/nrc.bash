#!/bin/bash
ord() {
  printf '%d' "'$1"
}

while read line; do
  declare -a ch
  for ((c=0; c<${#line}; c++)); do
    ((ch[`ord ${line:$c:1}`]++))
  done
  for ((c=0; c<${#line}; c++)); do
    if [ ${ch[`ord ${line:$c:1}`]} -eq 1 ]; then
      printf ${line:$c:1}
      break
    fi
  done
  echo
  unset ch
done <$1
