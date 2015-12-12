while read -r line || [ -n "$line" ]; do
  echo "$line"|tr "abcdefghijklmnopqrstuvwxyz" "uvwxyznopqrstghijklmabcdef"
done <$1
