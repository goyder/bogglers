from typing import List
from bogglers import dictionary_filepath
import random

"""
data.py
In this script, capture static data structures that are common to the Boggle structure.
"""

dice = [
    ["P", "S", "A", "F", "K", "F"],
    ["H", "I", "U", "N", "QU", "M"],
    ["I", "M", "T", "O", "U", "C"],
    ["I", "T", "S", "T", "D", "Y"],
    ["L", "R", "E", "I", "X", "D"],
    ["E", "E", "U", "S", "I", "N"],
    ["E", "R", "W", "T", "H", "V"],
    ["T", "Y", "E", "L", "R", "T"],
    ["N", "G", "E", "E", "A", "A"],
    ["T", "O", "E", "S", "S", "I"],
    ["B", "B", "A", "O", "O", "S"],
    ["H", "E", "E", "N", "H", "E"],
    ["W", "T", "O", "O", "T", "A"],
    ["S", "O", "A", "C", "H", "P"],
    ["R", "N", "Z", "N", "H", "L"],
    ["D", "E", "Y", "L", "V", "R"],
]

column_names = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


def generate_dictionary(filepath: str = dictionary_filepath) -> List[str]:
    """
    Return a dictionary of words.
    :param filepath: Filepath to the dictionary.
    :return:
    """
    with open(filepath, 'r') as f:
        dictionary = [line.rstrip() for line in f]

    return dictionary


def generate_network(rows: int, columns: int) -> dict:
    """
    Generate a network that we will explore.
    :param rows: Number of rows to include in network.
    :param columns: Number of columns to include in network.
    :return: Dictionary of network.
    """

    if rows < 1:
        raise ValueError("Need at least one row.")
    if columns < 1:
        raise ValueError("Need at least one column.")
    if rows > 10 or columns > 10:
        raise ValueError("Calm down.")

    # Generate the network, filling in connections row by row, column by column.
    board_network = dict()
    for row in range(1, rows + 1):
        for column in range(1, columns + 1):
            # Define what we're looking at
            tile_name = column_names[column - 1] + str(row)

            # Generate connections
            adjacent_tiles = []
            for row_delta in [-1, 0, 1]:
                for column_delta in [-1, 0, 1]:
                    # If both are zero, we're looking at ourselves.
                    if row_delta == 0 and column_delta == 0:
                        continue

                    # Are we past the edge of the board?
                    if (row + row_delta > rows) or (row + row_delta < 1):
                        continue
                    if (column + column_delta > columns) or (column + column_delta < 1):
                        continue

                    # Otherwise map our the tile next to it.
                    adjacent_tiles.append(
                        column_names[column - 1 + column_delta] + str(row + row_delta)
                    )

            # Assign and move on
            board_network[tile_name] = adjacent_tiles

    return board_network


def generate_letters(dice_selection: List[List]) -> List[str]:
    """
    Given a selection of dice options, return a list of the letters that are set.
    TODO: Do we need counts for the number of dice, etc?
    :param dice_selection: The list of lists of the dice contents.
    :return: A list of strings of the letters that are selected.
    """
    # Shake all the dice...
    random.shuffle(dice_selection)

    # ...And a random side faces up.
    letters = [random.sample(die, 1)[0] for die in dice_selection]
    return letters


def assign_letters(network: dict, letters: List[str]) -> dict:
    """
    Assign letters to a given network and return the mapping network.
    :param network: Network dictionary.
    :param letters: List of letters to be used.
    :return: Mapping of the network nodes to the letters they represent.
    """
    # Put the keys in alphabetical order.
    keys = [str(key) for key in network.keys()]
    keys.sort()

    # Go through and match each value, working A1 -> A2 -> etc
    mapping = {}
    for n, key in enumerate(keys):
        mapping[key] = letters[n]

    return mapping


class StandardBoggle:
    """
    Generate a standard, 4x4 Boggle board, using the standard dice distributions of letters.
    This is more of a convenience data type than anything else.
    """

    def __init__(self):
        self.network = generate_network(4, 4)  # Standard!
        self.letters = generate_letters(dice)
        self.mapping = assign_letters(self.network, self.letters)

    def __repr__(self):
        return " ".join(self.letters[0:4]) + "\n" + \
               " ".join(self.letters[4:8]) + "\n" + \
               " ".join(self.letters[8:12]) + "\n" + \
               " ".join(self.letters[12:16])
