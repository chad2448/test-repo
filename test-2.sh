#!/bin/bash

go run go-info.go
go run cpu.go
go run memory.go

NUM_PROC=$(go run cpu.go | grep "Number" | cut -d ':' -f2 | tr -d ' ')

if [ $NUM_PROC != '2' ]; then
echo "Your total processor number is $NUM_PROC"
else
echo "Your total processor number is not 2. It's $NUM_PROC"
fi