while read line || [ -n "$line" ]; do
  u=0
  l=0
  for ((i=0; i<${#line}; i++)); do
    a="${line:$i:1}"
    if [ -z "${a//[a-z]}" ]; then
      ((l++))
    elif [ -z "${a//[A-Z]}" ]; then
      ((u++))
    fi
  done
  awk 'BEGIN {printf "lowercase: %.2f uppercase: %.2f\n", 100*'$l'/('$l'+'$u'), 100*'$u'/('$l'+'$u')}'
done <$1
