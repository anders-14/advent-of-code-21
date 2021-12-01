#!/bin/env python3

input = []

with open("input", "r") as f:
    input = [int(i) for i in f.read().split("\n")[:-1]]

times_increased = 0

for i in range(len(input)):
    if i+4 <= len(input):
        s = sum(input[i:i+4])
        if s - input[i+3] < s - input[i]:
            times_increased += 1

print(times_increased)
