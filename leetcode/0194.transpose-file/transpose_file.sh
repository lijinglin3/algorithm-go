#!/usr/bin/env bash

column=$(awk '{print NF}' file.txt | uniq)
for ((i = 1; i <= column; i++)); do
  cut -d' ' -f$i file.txt | xargs
done
