while read line || [ -n "$line" ]; do
  if [ -n "$line" ]; then
    a=( $line )
    l=$((${#a[@]}-1))
    printf "${a[$l]}"
    for ((i=$l-1; i>=0; i--)); do
        printf " ${a[$i]}"
    done
    echo
  fi
done <$1
