tr -cd "[:digit:] \n" <$1 | while read line || [ -n "$line" ]; do
  a=( $line )
  echo $(( (a[0] * 3 + a[1] * 4 + a[2] * 5) * a[3] / (a[0] + a[1] + a[2]) ))
done
