#!/bin/sh

seq 1 3 > A
seq 2 4 > B
seq 3 5 > C
seq 4 5 > D
seq 5 5 > E

./matchtable [A-E] | column -t
