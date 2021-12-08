#!/bin/env python3

input = []

with open("input", "r") as f:
    input = list(map(lambda x: x.split(" | ")[1].split(" "), f.read().split("\n")[:-2]))

count = 0
for i in input:
    for j in i:
        l = len(j)
        if l == 2 or l == 3 or l == 4 or l == 7:
            count += 1    

print("Count 1, 4, 7, 8 in output:", count)
