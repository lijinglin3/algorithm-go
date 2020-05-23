#!/usr/bin/env bash

awk '{if(NR==10){print $0}}' file.txt
