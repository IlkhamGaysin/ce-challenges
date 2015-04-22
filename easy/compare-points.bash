while read line || [ -n "$line" ]; do
  a=( $line )
  if [ ${a[0]} -eq ${a[2]} ]; then
    if [ ${a[1]} -eq ${a[3]} ]; then
      echo here
    elif [ ${a[1]} -lt ${a[3]} ]; then
      echo N
    else
      echo S
    fi
  elif [ ${a[1]} -eq ${a[3]} ]; then
    if [ ${a[0]} -lt ${a[2]} ]; then
      echo E
    else
      echo W
    fi
  elif [ ${a[0]} -lt ${a[2]} ]; then
    if [ ${a[1]} -lt ${a[3]} ]; then
      echo NE
    else
      echo SE
    fi
  else
    if [ ${a[1]} -lt ${a[3]} ]; then
      echo NW
    else
      echo SW
    fi
  fi
done <$1
