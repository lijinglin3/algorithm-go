#!/usr/bin/env bash

cat words.txt | tr ' ' '\n' | grep -v ^$ | sort | uniq -c | sort -rn | awk '{print $2 " " $1}'
