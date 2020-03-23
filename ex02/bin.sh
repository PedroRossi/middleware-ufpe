#!/bin/bash
for i in {1..5}
do
  echo $i
  range=$[10000 / $i]
  j=0
  while [ $j -lt $range ]
  do
    j=$[$j+1]
    ./client $i $1 localhost:8080 20 >> $1/$i.txt
  done
done
