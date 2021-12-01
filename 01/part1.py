#!/bin/env python3

input = []

with open("input", "r") as f:
    input = [int(i) for i in f.read().split("\n")[:-1]]

prev = 0
times_increased = -1

for i in input:
    if prev < i:
        times_increased += 1

    prev = i

print(times_increased)
