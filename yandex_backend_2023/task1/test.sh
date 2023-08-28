#!/bin/bash

for ((i=0; ; i++))
do
  ./generate
  ./main_slice > result1.txt
  ./main_map   > result2.txt
  diff -w result1.txt result2.txt

  echo "finish test " + "$i"
  sleep 1
done