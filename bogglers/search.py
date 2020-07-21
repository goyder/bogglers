import typing
import bogglers.data as data
"""
search.py

Script to search a given Boggle network.
"""

# TODO: Figure out an alternative to this one? Do we have to load this every time?
dictionary = data.generate_dictionary()


def generate_words(network, letter_mapping, max_letters=7, min_letters=3):
    """
    Search through a Network and find all the words you can.
    :param network: Network to search.
    :param letter_mapping: Mapping between each node and the letter.
    :param max_letters: Number of letters before you punch out and stop searching. (Heuristic value.)
    :param min_letters: Minimum number of letters in a word.
    :return: List of words that were found.
    """

    # Starting points
    starting_values = list(network.keys())

    # What have we found?
    words = []

    # Search, starting from all our nodes
    for starting_value in starting_values:
        letters = ""
        path = []
        words = search_node(starting_value, network, letter_mapping, letters,
                            path, words, max_letters, min_letters)

    return words


def search_node(tile: str, network: dict, letter_mapping: typing.Dict[str, str],
                letters: str, path: typing.List[str], words: typing.List[str],
                max_letters: int, min_letters: int) -> typing.List[str]:
    """
    For a node, figure out whether we can make a word, where we should go next, whether we should stop...
    This is a recursive function.
    :param tile: Current tile to search on.
    :param network: The network dictionary to use.
    :param letter_mapping: The mapping of tiles to letters to be used.
    :param letters: The current list of letters ("word") that the path is on thusly.
    :param path: The path of tiles ("A1", "A2", etc...) that have been taken so far.
    :param words: Words that have been discovered so far.
    :param max_letters: The maximum number of letters to be searched for before giving up on a branch.
    :param min_letters: The minimum number of letters allowed within a Boggle world.
    :return: A list of the words found in the search.
    """
    # Check if we've made a new word
    letters = letters + letter_mapping[tile]
    if letters in dictionary:
        # TODO: Evaluate the performance import of this
        if (letters not in words) and (len(letters) >= min_letters):
            words.append(letters)

    # Check if we should stop - heuristic approach
    if len(path) >= max_letters:
        return words

    # TODO: Stop if there are no more words that can be made with this string of letters

    # Check what next steps are available to us
    next_steps = [connection for connection in network[tile] if connection not in path]

    # And if we have next steps, take them
    for next_step in next_steps:
        new_path = path + [tile]
        words = search_node(next_step, network, letter_mapping, letters,
                            new_path, words, max_letters, min_letters)

    return words



