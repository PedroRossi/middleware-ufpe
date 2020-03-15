#!/bin/bash
for i in {1..5}
do
  for j in {1..10000}
  do
     ./client $i $1 localhost:8080 package >> $1/$i.txt
  done
done
