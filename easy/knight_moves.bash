chr() {
  printf \\$(printf '%03o' $1)
}

ord() {
  printf '%d' "'$1"
}

while read line || [ -n "$line" ]; do
  l=$((`ord ${line:0:1}`-97))
  r=$((`ord ${line:1:1}`-49))
  declare -a m
  if [ $l -gt 1 ] && [ $r -gt 0 ]; then
    m+=(`chr $((97+$l-2))``chr $((49+$r-1))`)
  fi
  if [ $l -gt 1 ] && [ $r -lt 7 ]; then
    m+=(`chr $((97+$l-2))``chr $((49+$r+1))`)
  fi
  if [ $l -gt 0 ] && [ $r -gt 1 ]; then
    m+=(`chr $((97+$l-1))``chr $((49+$r-2))`)
  fi
  if [ $l -gt 0 ] && [ $r -lt 6 ]; then
    m+=(`chr $((97+$l-1))``chr $((49+$r+2))`)
  fi
  if [ $l -lt 7 ] && [ $r -gt 1 ]; then
    m+=(`chr $((97+$l+1))``chr $((49+$r-2))`)
  fi
  if [ $l -lt 7 ] && [ $r -lt 6 ]; then
    m+=(`chr $((97+$l+1))``chr $((49+$r+2))`)
  fi
  if [ $l -lt 6 ] && [ $r -gt 0 ]; then
    m+=(`chr $((97+$l+2))``chr $((49+$r-1))`)
  fi
  if [ $l -lt 6 ] && [ $r -lt 7 ]; then
    m+=(`chr $((97+$l+2))``chr $((49+$r+1))`)
  fi
  echo ${m[*]}
  unset m
done <$1
