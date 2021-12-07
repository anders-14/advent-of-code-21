#!/bin/env python3

input = []

with open("input", "r") as f:
    input = [int(i) for i in f.readline().split(",")]


def calc_fuel_usage(dist):
    return sum(range(dist+1))


most_left = min(input)
most_right = max(input)

fuel_usage = []

for i in range(most_left, most_right+1):
    fuel_used = sum(map(lambda x: calc_fuel_usage(abs(x - i)), input))
    fuel_usage.append(fuel_used)

print("Required fuel usage:", min(fuel_usage))
