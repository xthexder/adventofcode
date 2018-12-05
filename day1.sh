#!/bin/sh

awk '{s+=$1} END {print s}' day1.txt

