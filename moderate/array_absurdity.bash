sed -e "s/^.*;//" -e "s/,/ /g" $1|while read line; do
  a=( $line )
  echo "${a[*]}"|tr " " "\n"|sort|uniq -d
done
