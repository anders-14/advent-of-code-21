#!/bin/env python3

import re

FILENAME = "input"
BOARDS = 100

numbers = []
drawn_numbers = []
boards = []

with open(FILENAME, "r") as f:
    numbers = [int(i) for i in f.readline().split(",")]
    f.readline()
    
    for i in range(BOARDS):
        board = []
        for j in range(5):
            line = f.readline().strip()
            board.append([int(a) for a in re.split('\s+', line)])
        boards.append(board)
        f.readline()
         

def col(board, idx):
    col = []
    for row in board:
        col.append(row[idx])
    return col


def check_board(board):
    for i in range(5):
        nums_found = 0
        for num in drawn_numbers:
            if num in col(board, i):
                nums_found += 1
        if nums_found == 5:
            return True
    
    for row in board:
        nums_found = 0
        for num in drawn_numbers:
            if num in row:
                nums_found += 1
        if nums_found == 5:
            return True

    return False


def find_unmarked(board):
    unmarked = []
    all_nums_on_board = sum(board, [])
    for num in all_nums_on_board:
        if num not in drawn_numbers:
            unmarked.append(num) 

    return unmarked


def draw_number():
    global numbers
    drawn_numbers.append(numbers[0])
    numbers = numbers[1:]


draw_number()
draw_number()
draw_number()
draw_number()
while len(numbers) != 0:
    if len(boards) == 1:
        draw_number()
        while not check_board(boards[0]):
            draw_number()

        print("Final score: " + str(sum(find_unmarked(boards[0])) * drawn_numbers[-1]))
        break

    draw_number()
    for board in boards:
        if check_board(board):
            boards.remove(board)
