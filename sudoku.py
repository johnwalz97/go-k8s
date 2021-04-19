import sys


def print_puzzle(puzzle):
    for row in puzzle:
        print(" ".join([str(num) for num in row]))


def find_empty_position(puzzle):
    for row in range(9):
        for col in range(9):
            if puzzle[row][col] == 0:
                return row, col

    return -1, -1


def can_place_number(number, puzzle, row, col):
    for i in range(9):
        if puzzle[row][i] == number:
            return False

    for i in range(9):
        if puzzle[i][col] == number:
            return False

    for i in range(3):
        for j in range(3):
            if puzzle[i + (row - row % 3)][j + (col - col % 3)] == number:
                return False

    return True


def solve_puzzle(puzzle):
    row, col = find_empty_position(puzzle)

    if row == -1:
        print_puzzle(puzzle)
        return True

    for number in range(1, 10):
        if can_place_number(number, puzzle, row, col):
            puzzle[row][col] = number

            if solve_puzzle(puzzle):
                return True

            puzzle[row][col] = 0

    return False


if __name__ == "__main__":
    lines = [line.strip('\n') for line in sys.stdin]
    num_puzzles = int(lines[0])

    puzzles = [
        [
            [int(c) for c in lines[j*i + 1].replace(" ", "")]
            for j in range(9)
        ]
        for i in range(1, num_puzzles + 1)
    ]

    for sudoku_puzzle in puzzles:
        solve_puzzle(sudoku_puzzle)
