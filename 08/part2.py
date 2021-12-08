#!/bin/env python3

input = []

with open("input", "r") as f:
    input = f.read().split("\n")[:-2]


def count_same_symbols(a, b):
    count = 0
    for x in a:
        if x in b:
            count += 1
    return count


def sort_string(string):
    return "".join(sorted(list(string)))


def determine_digits(patterns):
    digits = {}
    patterns.sort(key=lambda x: len(x))
    for i in patterns:
        if len(i) == 2:
            digits[1] = sort_string(i)
        elif len(i) == 3:
            digits[7] = sort_string(i)
        elif len(i) == 4:
            digits[4] = sort_string(i)
        elif len(i) == 5:
            if count_same_symbols(i, digits[1]) == 1 and count_same_symbols(i, digits[4]) == 3:
                digits[5] = sort_string(i)
            elif count_same_symbols(i, digits[1]) == 2 and count_same_symbols(i, digits[4]) == 3:
                digits[3] = sort_string(i)
            else:
                digits[2] = sort_string(i)
        elif len(i) == 6:
            if count_same_symbols(i, digits[1]) == 1:
                digits[6] = sort_string(i)
            else:
                if count_same_symbols(i, digits[3]) == 4:
                    digits[0] = sort_string(i)
                else:
                    digits[9] = sort_string(i)
        elif len(i) == 7:
            digits[8] = sort_string(i)

    return digits

def compare_output(output, digits):
    num = ""
    for i in output:
        num += str(list(digits.keys())[list(digits.values()).index(sort_string(i))])

    return int(num)


s = 0
for a in input:
    a = a.split(" | ")
    s += compare_output(a[1].split(" "), determine_digits(a[0].split(" ")))

print("Sum of 7 segment display numbers:", s)
