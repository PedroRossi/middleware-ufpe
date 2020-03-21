#!/bin/bash
for i in {1..5}
do
  for j in {1..10000}
  do
     ./client $i localhost:8080 package >> results/$i.txt
  done
done
