#!/bin/sh

seq 1 26 | xargs -n1 -I{} bash -c "cut -b{} --complement day2.txt | sort | uniq -d"
